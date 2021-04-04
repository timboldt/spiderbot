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

type Leg struct {
	servos [3]*Servo
	hipPt  Point3D
}

type Point3D struct {
	x, y, z int16
}

func NewLeg(hipPt Point3D, hip *Servo, coxa *Servo, tibia *Servo) *Leg {
	return &Leg{
		servos: [3]*Servo{hip, coxa, tibia},
		hipPt:  hipPt,
	}
}

func (l *Leg) ServoValues(toePt Point3D) (uint16, uint16, uint16) {
	var hipDeg int16 = 45
	var coxaDeg int16 = 90
	var tibiaDeg int16 = 90
	return l.servos[0].DegreesToMicros(hipDeg), l.servos[1].DegreesToMicros(coxaDeg), l.servos[2].DegreesToMicros(tibiaDeg)
}
