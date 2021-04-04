package spider

import (
	"testing"
)

func TestGetServos(t *testing.T) {
	servos := GetServos()
	want := 12
	if len(servos) != want {
		t.Fatalf("GetServos() returned %d servos, want %d", len(servos), want)
	}
}

func TestPin(t *testing.T) {
	s := Servo{
		pin: 42,
	}
	var want byte = 42
	got := s.Pin()
	if got != want {
		t.Errorf("s.Pin() returned %d, want %d", got, want)
	}
}

func TestDegreesToMicros(t *testing.T) {
	s := Servo{
		minVal:          800,
		maxVal:          2200,
		ninetyDegMicros: 1500,
		reversed:        false,
	}
	tests := []struct {
		deg  int16
		rev  bool
		want uint16
	}{
		{-10, false, 800},
		{0, false, 800},
		{10, false, 800},
		{45, false, 1000},
		{90, false, 1500},
		{200, false, 2200},
		{-10, true, 2200},
		{0, true, 2200},
		{10, true, 2200},
		{45, true, 2000},
		{90, true, 1500},
		{200, true, 800},
	}

	for _, tt := range tests {
		s.reversed = tt.rev
		got := s.DegreesToMicros(tt.deg)
		if got != tt.want {
			t.Errorf("s.DegressToMicros(%d) = %d, rev=%v, want %d", tt.deg, got, tt.rev, tt.want)
		}
	}
}
