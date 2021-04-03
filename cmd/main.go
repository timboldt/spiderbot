// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This is an experiment designed for the Adafruit nrf52840 Express.
package main

import (
	"fmt"
	"machine"
	"strconv"
	"time"

	"github.com/timboldt/spiderbot/pkg/pca9685"
)

func main() {
	time.Sleep(10 * time.Second)
	println("hi")

	//
	// === Initialize hardware ===
	//
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ,
	})

	pwm := pca9685.New(machine.I2C0)
	if err := pwm.Configure(); err != nil {
		fmt.Printf("configure failed: %v", err)
	}

	var currPin byte
	inbuf := make([]byte, 64)
	inbufIdx := 0
	uart := machine.UART0
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			// Echo what the user types.
			uart.WriteByte(data)

			switch data {
			case '\n':
				fallthrough
			case '\r':
				if inbufIdx > 0 {
					if inbuf[0] == 'p' && inbufIdx > 1 {
						val, err := strconv.Atoi(string(inbuf[1:inbufIdx]))
						if err != nil {
							fmt.Println(err)
						} else {
							currPin = byte(val)
							fmt.Printf("Setting pin %d to %d\n", currPin, 1500)
							if err := pwm.SetPin(currPin, uint16(1500)); err != nil {
								fmt.Printf("set pin PWM failed: %v", err)
							}
						}
					} else {
						val, err := strconv.Atoi(string(inbuf[:inbufIdx]))
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Printf("Setting pin %d to %d\n", currPin, val)
							if err := pwm.SetPin(currPin, uint16(val)); err != nil {
								fmt.Printf("set pin PWM failed: %v", err)
							}
						}
					}
					inbufIdx = 0
				}
			default:
				inbuf[inbufIdx] = data
				inbufIdx++
			}
		}
	}
}
