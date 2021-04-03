package pca9685

// Constants/addresses used for I2C.

// The I2C address which this device listens to by default.
const Address = 0x40

// Register names and addresses.
const (
	REG_MODE1 byte = iota
	REG_MODE2
	REG_SUBADR1
	REG_SUBADR2
	REG_SUBADR3
	REG_ALLCALLADR
	REG_PWM0_ON_L
	REG_PWM0_ON_H
	REG_PWM0_OFF_L
	REG_PWM0_OFF_H
	REG_PRESCALE = 0xFE
)

// MODE1 bit values.
const (
	MODE1_ALLCAL byte = 1 << iota
	MODE1_SUB3
	MODE1_SUB2
	MODE1_SUB1
	MODE1_SLEEP
	MODE1_AI
	MODE1_EXTCLK
	MODE1_RESTART
)

// Typical PWM prescalar values (assuming a 25MHz clock).
const (
	CLOCK_MHZ       = 25
	PRESCALE_SERVO  = 121 // 50Hz
	MICROS_PER_TICK = (PRESCALE_SERVO + 1 + CLOCK_MHZ/2) / CLOCK_MHZ
)
