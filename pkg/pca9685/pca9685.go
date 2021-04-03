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

// Package pca9685 is a driver for the PCA9685 I2C servo/pwm board.
// e.g. https://www.adafruit.com/product/815
package pca9685

import (
	"fmt"
	"time"

	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to the device.
type Device struct {
	bus     drivers.I2C
	address uint16
}

// New creates a new PCA9685 connection. The I2C bus must already be
// configured.
//
// This function only creates the Device object. It does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{
		bus:     bus,
		address: Address,
	}
}

// Configure sets up the device for communication.
func (d *Device) Configure() error {
	if err := d.reset(); err != nil {
		return err
	}

	oldMode, err := d.readRegByte(REG_MODE1)
	if err != nil {
		return err
	}

	sleepMode := (oldMode &^ MODE1_RESTART) | MODE1_SLEEP
	if err := d.writeRegByte(REG_MODE1, sleepMode); err != nil {
		return err
	}
	if err := d.writeRegByte(REG_PRESCALE, PRESCALE_SERVO); err != nil {
		return err
	}
	if err := d.writeRegByte(REG_MODE1, oldMode); err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	if err := d.writeRegByte(REG_MODE1, oldMode|MODE1_RESTART|MODE1_AI); err != nil {
		return err
	}

	return nil
}

// SetPin sets the pulse width for a given pin to an approximate number of microseconds.
// Valid pin values are 0..15.
// Valid values are 1000 to 2000, or zero.
// A value of 0 turns off the servo.
func (d *Device) SetPin(pin byte, micros uint16) error {
	if pin > 15 {
		return fmt.Errorf("invalid pin: %d", pin)
	}
	if micros != 0 && (micros < 500 || micros > 3000) {
		return fmt.Errorf("invalid servo timing: %d us", micros)
	}
	val := micros / MICROS_PER_TICK
	if val == 0 {
		// Special value for fully off is (0, 4096).
		val = 4096
	}
	reg := REG_PWM0_ON_L + pin*4
	data := []byte{0, 0, byte(val) & 0xFF, byte(val >> 8)}
	return d.bus.WriteRegister(uint8(d.address), reg, data[:])
}

func (d *Device) readRegByte(reg byte) (byte, error) {
	var val [1]byte
	if err := d.bus.ReadRegister(uint8(d.address), reg, val[:]); err != nil {
		return 0, err
	}
	if len(val) != 1 {
		return 0, fmt.Errorf("byte register is not length 1; %d", len(val))
	}
	return val[0], nil
}

func (d *Device) writeRegByte(reg, val byte) error {
	data := []byte{val}
	return d.bus.WriteRegister(uint8(d.address), reg, data[:])
}

// reset puts the device back in the default power-on state.
func (d *Device) reset() error {
	err := d.writeRegByte(REG_MODE1, MODE1_RESTART)
	// TODO: Watch for reset to complete, rather than blindly sleeping.
	time.Sleep(10 * time.Millisecond)
	return err
}
