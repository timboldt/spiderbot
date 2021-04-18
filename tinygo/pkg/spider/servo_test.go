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

package spider

import (
	"testing"
)

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
		minVal:        800,
		maxVal:        2200,
		zeroDegMicros: 500,
		reversed:      false,
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
