package ssc32u

import "fmt"

type Board struct {
	tty      string
	baudrate uint
	servos   map[uint]*Servo
}

func New(tty string, baudrate uint) Board {
	return Board{
		tty:      tty,
		baudrate: baudrate,
		servos:   make(map[uint]*Servo),
	}
}

func (b *Board) Close() {

}

func (b *Board) AddServo(id uint, name string) *Servo {
	servo := &Servo{
		name:         name,
		id:           id,
		position:     1500,
		min_position: 500,
		max_position: 2500,
	}
	b.servos[id] = servo
	return servo
}

func (b *Board) Servo(id uint) *Servo {
	return b.servos[id]
}

func (b Board) String() string {
	out := fmt.Sprintf("SSC-32U on %s (%d baud) {\n", b.tty, b.baudrate)
	for id, servo := range b.servos {
		out += fmt.Sprintf("  %02d (%s): %d\n", id, servo.name, servo.position)
	}
	out += "}"
	return out
}
