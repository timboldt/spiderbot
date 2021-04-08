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

type Spider struct {
	servos [12]Servo
	legs   [4]Leg
}

var (
	theSpider Spider
)

// Init initializes the Spider instance, which is a simple singleton.
func Init() *Spider {
	theSpider.initServos()
	for i := 0; i < 4; i++ {
		theSpider.legs[i].init(LegPosition(i))
	}
	return &theSpider
}

func servoId(pos LegPosition, joint Joint) uint8 {
	return uint8(pos)*3 + uint8(joint)
}

func (s *Spider) initServos() {
	s.servos = [12]Servo{
		{
			pin:             servoId(FrontRight, Hip),
			minVal:          1500,
			maxVal:          2500,
			ninetyDegMicros: 1700,
			reversed:        true,
		},
		{
			pin:             servoId(FrontRight, Coxa),
			minVal:          1200,
			maxVal:          2600,
			ninetyDegMicros: 2111,
			reversed:        false,
		},
		{
			pin:             servoId(FrontRight, Tibia),
			minVal:          1400,
			maxVal:          2500,
			ninetyDegMicros: 1900,
			reversed:        false,
		},
		{
			pin:             servoId(FrontLeft, Hip),
			minVal:          700,
			maxVal:          1700,
			ninetyDegMicros: 1611,
			reversed:        false,
		},
		{
			pin:             servoId(FrontLeft, Coxa),
			minVal:          500,
			maxVal:          1900,
			ninetyDegMicros: 1045,
			reversed:        true,
		},
		{
			pin:             servoId(FrontLeft, Tibia),
			minVal:          1300,
			maxVal:          2400,
			ninetyDegMicros: 1900,
			reversed:        true,
		},
		{
			pin:             servoId(BackRight, Hip),
			minVal:          700,
			maxVal:          1700,
			ninetyDegMicros: 1800,
			reversed:        false,
		},
		{
			pin:             servoId(BackRight, Coxa),
			minVal:          700,
			maxVal:          2100,
			ninetyDegMicros: 1189,
			reversed:        true,
		},
		{
			pin:             servoId(BackRight, Tibia),
			minVal:          1500,
			maxVal:          2500,
			ninetyDegMicros: 2155,
			reversed:        true,
		},
		{
			pin:             servoId(BackLeft, Hip),
			minVal:          1400,
			maxVal:          2400,
			ninetyDegMicros: 2500,
			reversed:        false,
		},
		{
			pin:             servoId(BackLeft, Coxa),
			minVal:          1000,
			maxVal:          2200,
			ninetyDegMicros: 1600,
			reversed:        false,
		},
		{
			pin:             servoId(BackLeft, Tibia),
			minVal:          1100,
			maxVal:          2200,
			ninetyDegMicros: 1600,
			reversed:        false,
		},
	}
}
