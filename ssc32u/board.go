// Copyright 2020 Google LLC
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

package ssc32u

import (
	"fmt"

	"github.com/tarm/serial"
)

type Board struct {
	port     *serial.Port
	tty      string
	baudrate uint
	servos   map[string]*Servo
}

func New() Board {
	return Board{
		servos: make(map[string]*Servo),
	}
}

func (b *Board) Connect(tty string, baudrate uint) bool {
	config := &serial.Config{Name: tty, Baud: int(baudrate)}
	port, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println(err)
		return false
	}
	b.port = port
	b.tty = tty
	b.baudrate = baudrate
	return true
}

func (b *Board) Close() {
	if b.port != nil {
		for i := 0; i < 32; i++ {
			b.port.Write([]byte(fmt.Sprintf("STOP%d\r", i)))
		}
		b.port.Close()
		b.port = nil
	}
}

func (b *Board) Commit(micros uint) {
	if b.port != nil {
		_, err := b.port.Write([]byte(b.commandString(micros)))
		if err == nil {
			for _, servo := range b.servos {
				servo.isModified = false
			}
		}
	}
}

func (b *Board) AddServo(id uint, name string) *Servo {
	servo := &Servo{
		name:       name,
		id:         id,
		position:   1500,
		isModified: true,
	}
	b.servos[name] = servo
	return servo
}

func (b *Board) Servo(name string) *Servo {
	return b.servos[name]
}

func (b Board) String() string {
	out := fmt.Sprintf("SSC-32U on %s (%d baud) {\n", b.tty, b.baudrate)
	for id, servo := range b.servos {
		out += fmt.Sprintf("  %02d (%s): %d\n", id, servo.name, servo.position)
	}
	out += "}"
	return out
}

func (b Board) commandString(micros uint) string {
	cmd := ""
	for _, servo := range b.servos {
		cmd += servo.commandString()
	}
	if micros > 65535 {
		micros = 65535
	}
	cmd += fmt.Sprintf("T%d\r\n", micros)
	fmt.Println(cmd)
	return cmd
}
