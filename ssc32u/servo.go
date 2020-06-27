package ssc32u

type Servo struct {
	name         string
	id           uint
	position     uint
	min_position uint
	max_position uint
}

func (s *Servo) SetAngleDegrees(angle int) {
	// TODO
	s.SetPosition(999)
}

func (s *Servo) SetPosition(position uint) {
	if position > s.max_position {
		s.position = s.max_position
	} else if position < s.min_position {
		s.position = s.min_position
	} else {
		s.position = position
	}
}
