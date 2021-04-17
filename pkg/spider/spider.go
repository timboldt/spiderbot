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
	"github.com/timboldt/spiderbot/pkg/pca9685"
)

type Spider struct {
	pwm    pca9685.Device
	servos [12]Servo
	legs   [4]Leg
}

var (
	theSpider Spider
)

// Init initializes the Spider instance, which is a simple singleton.
func Init(pwm pca9685.Device) *Spider {
	theSpider.pwm = pwm
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
		// FR BC
		{
			pin:           0,
			minVal:        1500,
			maxVal:        2500,
			zeroDegMicros: 1700,
			reversed:      false,
		},
		// FR CF
		{
			pin:           1,
			minVal:        1200,
			maxVal:        2600,
			zeroDegMicros: 2111,
			reversed:      false,
		},
		// FR FT
		{
			pin:           2,
			minVal:        1400,
			maxVal:        2500,
			zeroDegMicros: 900,
			reversed:      false,
		},
		// FL BC
		{
			pin:           3,
			minVal:        700,
			maxVal:        1700,
			zeroDegMicros: -400,
			reversed:      false,
		},
		// FL CF
		{
			pin:           4,
			minVal:        500,
			maxVal:        1900,
			zeroDegMicros: 1045,
			reversed:      true,
		},
		// FL FT
		{
			pin:           5,
			minVal:        1300,
			maxVal:        2400,
			zeroDegMicros: 2800,
			reversed:      true,
		},
		// BR BC
		{
			pin:           6,
			minVal:        700,
			maxVal:        1700,
			zeroDegMicros: 1800,
			reversed:      false,
		},
		// BR CF
		{
			pin:           7,
			minVal:        700,
			maxVal:        2100,
			zeroDegMicros: 1189,
			reversed:      true,
		},
		// BR FT
		{
			pin:           8,
			minVal:        1500,
			maxVal:        2500,
			zeroDegMicros: 3100,
			reversed:      true,
		},
		// BL BC
		{
			pin:           9,
			minVal:        1400,
			maxVal:        2400,
			zeroDegMicros: 3500,
			reversed:      false,
		},
		// BL CF
		{
			pin:           10,
			minVal:        1000,
			maxVal:        2200,
			zeroDegMicros: 1600,
			reversed:      false,
		},
		// BL FT
		{
			pin:           11,
			minVal:        1100,
			maxVal:        2200,
			zeroDegMicros: 600,
			reversed:      false,
		},
	}
}

func (s *Spider) SendCommandsToServos() {
	// fmt.Println("Moving!")
	for leg := LegPosition(0); leg < LegPosition(4); leg++ {
		bc, cf, ft := s.legs[leg].JointAngles()
		// fmt.Printf("Leg: %d Servo: %d Angle: %f usec: %d\n", leg, servoId(leg, BodyCoxa), bc, s.servos[servoId(leg, BodyCoxa)].RadiansToMicros(bc))
		// fmt.Printf("Leg: %d Servo: %d Angle: %f usec: %d\n", leg, servoId(leg, CoxaFemur), cf, s.servos[servoId(leg, CoxaFemur)].RadiansToMicros(cf))
		// fmt.Printf("Leg: %d Servo: %d Angle: %f usec: %d\n", leg, servoId(leg, FemurTibia), ft, s.servos[servoId(leg, FemurTibia)].RadiansToMicros(ft))
		s.pwm.SetPin(servoId(leg, BodyCoxa), s.servos[servoId(leg, BodyCoxa)].RadiansToMicros(bc))
		s.pwm.SetPin(servoId(leg, CoxaFemur), s.servos[servoId(leg, CoxaFemur)].RadiansToMicros(cf))
		s.pwm.SetPin(servoId(leg, FemurTibia), s.servos[servoId(leg, FemurTibia)].RadiansToMicros(ft))
	}
}

func (s *Spider) SetAll(pt Point3D) {
	for leg := LegPosition(0); leg < LegPosition(4); leg++ {
		s.legs[leg].toePt = pt
	}
}
