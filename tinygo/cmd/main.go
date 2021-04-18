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
	"math"
	"time"

	"github.com/timboldt/spiderbot/pkg/pca9685"
	"github.com/timboldt/spiderbot/pkg/spider"
)

func main() {
	time.Sleep(100 * time.Millisecond)
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

	spdr := spider.Init(pwm)
	// inbuf := make([]byte, 64)
	// inbufIdx := 0
	// uart := machine.UART0
	theta := 0.0
	for {
		spdr.SendCommandsToServos()
		time.Sleep(10 * time.Millisecond)
		theta += math.Pi / 50
		spdr.SetAll(spider.Point3D{math.Sin(theta) * 20, math.Cos(theta) * 20, math.Sin(theta/2) * 5})
		// if uart.Buffered() > 0 {
		// 	data, _ := uart.ReadByte()
		// 	// Echo what the user types.
		// 	uart.WriteByte(data)

		// 	switch data {
		// 	case '\n':
		// 		fallthrough
		// 	case '\r':
		// 		spider.Move()
		// 		// if inbufIdx > 0 {
		// 		// 	if inbuf[0] == 's' && inbufIdx > 1 {
		// 		// 		val, err := strconv.Atoi(string(inbuf[1:inbufIdx]))
		// 		// 		if err != nil {
		// 		// 			fmt.Println(err)
		// 		// 		} else {
		// 		// 			currServo = servos[val]
		// 		// 			micros := currServo.DegreesToMicros(90)
		// 		// 			fmt.Printf("Setting pin %d to %d\n", currServo.Pin(), micros)
		// 		// 			if err := pwm.SetPin(currServo.Pin(), micros); err != nil {
		// 		// 				fmt.Printf("set servo PWM failed: %v", err)
		// 		// 			}
		// 		// 		}
		// 		// 	} else {
		// 		// 		val, err := strconv.Atoi(string(inbuf[:inbufIdx]))
		// 		// 		if err != nil {
		// 		// 			fmt.Println(err)
		// 		// 		} else {
		// 		// 			micros := currServo.DegreesToMicros(int16(val))
		// 		// 			fmt.Printf("Setting pin %d to %d\n", currServo.Pin(), micros)
		// 		// 			if err := pwm.SetPin(currServo.Pin(), micros); err != nil {
		// 		// 				fmt.Printf("set servo PWM failed: %v", err)
		// 		// 			}
		// 		// 		}
		// 		// 	}
		// 		// 	inbufIdx = 0
		// 		// }
		// 	default:
		// 		inbuf[inbufIdx] = data
		// 		inbufIdx++
		// 	}
		// }
	}
}
