package ssc32u

import "testing"

func TestSetAngleDegrees(t *testing.T) {
	var tests = []struct {
		in   float32
		want uint
	}{
		{-90, 500},
		{0, 500},
		{45, 1000},
		{90, 1500},
		{135, 2000},
		{179, 2488},
		{180, 2500},
		{300, 2500},
		{900, 2500},
	}

	for _, tt := range tests {
		s := Servo{name: "test-servo",
			id:       42,
			position: 1500,
		}
		s.SetAngleDegrees(tt.in)
		if s.position != tt.want {
			t.Errorf("SetAngleDegrees(%f) = %d; want %d", tt.in, s.position, tt.want)
		}

	}
}
