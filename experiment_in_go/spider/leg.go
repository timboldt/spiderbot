// Copyright 2020 Google LLC
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

const coxaLength = 24
const femurLength = 38
const tibiaLength = 80

type Leg struct {
	x, y, z float64
}

func NewLeg() *Leg {
	return &Leg{
		x: coxaLength + femurLength*0.55,
		y: coxaLength + femurLength*0.55,
		z: -80,
	}
}

func (l *Leg) SetAbsolutePos(x, y, z float64) {
	l.x, l.y, l.z = x, y, z
}

func (l *Leg) SetRelativePos(x, y, z float64) {
	l.x += x
	l.y += y
	l.z += z
}

func (l Leg) GetAngles() (float64, float64, float64) {
	// Full horizontal distance (i.e., in the XY plane) from body-coxa joint to tibia tip.
	horizReach := math.Sqrt(l.x*l.x + l.y*l.y)

	// Horizontal distance (i.e., in the XY plane) from coxa-femur joint to tibia tip.
	horizReachFromCoxa := horizReach - coxaLength

	// Absolute distance from coxa-femur joint to tibia tip.
	absoluteReachFromCoxa := math.Sqrt(horizReachFromCoxa*horizReachFromCoxa + l.z*l.z)

	// We now have a triangle with sides [femurLength, tibiaLength, absoluteReachFromCoxa].
	// Solve using the law of cosines.
	// c^2 = a^2 + b^2 - 2*a*b*cos(C)
	// 2*a*b*cos(C) =  a^2 + b^2 - c^2
	// cos(C) = (a^2 + b^2 - c^2) / (2*a*b)
	cosWristNumerator := femurLength*femurLength + tibiaLength*tibiaLength - absoluteReachFromCoxa*absoluteReachFromCoxa
	cosWristDenominator := 2.0 * femurLength * tibiaLength
	wrist := math.Acos(cosWristNumerator/cosWristDenominator) * 180 / math.Pi

	// We now solve for the coxa-femur (knee) angle.
	angle1 := math.Atan2(l.z, horizReachFromCoxa)
	angle2 := math.Acos((femurLength*femurLength +
		absoluteReachFromCoxa*absoluteReachFromCoxa -
		tibiaLength*tibiaLength) /
		(2.0 * femurLength * absoluteReachFromCoxa))
	knee := 180 - (angle1+angle2)*180/math.Pi

	// Body-coxa (hip) angle.
	hip := math.Atan2(l.x, l.y) * 180 / math.Pi

	// TODO: Clean this up so that it works for both types of legs.
	return wrist - 90, knee - 180, 90 - hip
}
