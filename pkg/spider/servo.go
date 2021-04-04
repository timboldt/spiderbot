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

type Servo struct {
	pin             byte
	minVal          uint16
	maxVal          uint16
	ninetyDegMicros int16
	reversed        bool
}

func GetServos() [12]*Servo {
	return [12]*Servo{
		//
		// Front right leg.
		//
		// Hip.
		{
			pin:             0,
			minVal:          1500,
			maxVal:          2500,
			ninetyDegMicros: 1700,
			reversed:        true,
		},
		// Coxa.
		{
			pin:             1,
			minVal:          1200,
			maxVal:          2600,
			ninetyDegMicros: 2111,
			reversed:        false,
		},
		// Tibia.
		{
			pin:             2,
			minVal:          1400,
			maxVal:          2500,
			ninetyDegMicros: 1900,
			reversed:        false,
		},
		//
		// Front left leg.
		//
		// Hip.
		{
			pin:             3,
			minVal:          700,
			maxVal:          1700,
			ninetyDegMicros: 1611,
			reversed:        false,
		},
		// Coxa.
		{
			pin:             4,
			minVal:          500,
			maxVal:          1900,
			ninetyDegMicros: 1045,
			reversed:        true,
		},
		// Tibia.
		{
			pin:             5,
			minVal:          1300,
			maxVal:          2400,
			ninetyDegMicros: 1900,
			reversed:        true,
		},
		//
		// Back right leg.
		//
		// Hip.
		{
			pin:             6,
			minVal:          700,
			maxVal:          1700,
			ninetyDegMicros: 1800,
			reversed:        false,
		},
		// Coxa.
		{
			pin:             7,
			minVal:          700,
			maxVal:          2100,
			ninetyDegMicros: 1189,
			reversed:        true,
		},
		// Tibia.
		{
			pin:             8,
			minVal:          1500,
			maxVal:          2500,
			ninetyDegMicros: 2155,
			reversed:        true,
		},
		//
		// Back left leg.
		//
		// Hip.
		{
			pin:             9,
			minVal:          1400,
			maxVal:          2400,
			ninetyDegMicros: 2500,
			reversed:        false,
		},
		// Coxa.
		{
			pin:             10,
			minVal:          1000,
			maxVal:          2200,
			ninetyDegMicros: 1600,
			reversed:        false,
		},
		// Tibia.
		{
			pin:             11,
			minVal:          1100,
			maxVal:          2200,
			ninetyDegMicros: 1600,
			reversed:        false,
		},
	}
}

func (s Servo) Pin() byte {
	return s.pin
}

func (s *Servo) DegreesToMicros(deg int16) uint16 {
	if s.reversed {
		deg = 90 - deg
	} else {
		deg = deg - 90
	}
	micros := uint16(deg*100/9 + s.ninetyDegMicros)
	if micros > s.maxVal {
		return s.maxVal
	}
	if micros < s.minVal {
		return s.minVal
	}
	return micros
}
