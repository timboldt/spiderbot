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

import "math"

type Servo struct {
	pin           byte
	minVal        uint16
	maxVal        uint16
	zeroDegMicros int16
	reversed      bool
}

func (s Servo) Pin() byte {
	return s.pin
}

func (s *Servo) RadiansToMicros(rad float64) uint16 {
	deg := rad / math.Pi * 180
	return s.DegreesToMicros(int16(math.Round(deg)))
}

func (s *Servo) DegreesToMicros(deg int16) uint16 {
	if s.reversed {
		deg = -deg
	}
	micros := uint16(deg*100/9 + s.zeroDegMicros)
	if micros > s.maxVal {
		return s.maxVal
	}
	if micros < s.minVal {
		return s.minVal
	}
	return micros
}
