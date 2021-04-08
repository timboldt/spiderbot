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

type LegPosition uint8
type Joint uint8

// Leg positions.
const (
	FrontRight LegPosition = iota
	FrontLeft
	BackRight
	BackLeft
)

// Servo connection order, within a leg.
const (
	Hip Joint = iota
	Coxa
	Tibia
)

const (
	OneMillimeter int16 = 32 // Allowing approximate +/- one meter in a 16-bit integer.
	HipLength           = 23*OneMillimeter + OneMillimeter/2
	CoxaLength          = 38 * OneMillimeter
	TibiaLength         = 81 * OneMillimeter
)

type Leg struct {
	hipPt Point3D
}

// Represents a 3D point in space.
// X is towards the right of the robot.
// Y is towards the front of the robot.
// Z is towards the top of the robot.
// Distances are expressed in 1/32 of a millimeter.
type Point3D struct {
	X, Y, Z int16
}

func (l *Leg) init(pos LegPosition) {
	// The canonical zero position of the toe is with the hip and coxa horizontal and the tibia vertical.
	// Therefore the hip is displaced by (hip+coxa)/sqrt(2), using Pythagoras' theorem.
	// 16/23 is an approximation of 1/sqrt(2).
	hipOffset := (HipLength + CoxaLength) * 16 / 23
	var hipX, hipY int16
	switch pos {
	case FrontRight:
		hipX = -hipOffset
		hipY = -hipOffset
	case FrontLeft:
		hipX = hipOffset
		hipY = -hipOffset
	case BackRight:
		hipX = -hipOffset
		hipY = hipOffset
	case BackLeft:
		hipX = hipOffset
		hipY = hipOffset
	}
	l.hipPt = Point3D{
		X: hipX,
		Y: hipY,
		Z: TibiaLength,
	}
}

// func (l *Leg) ServoValues(toePt Point3D) (uint16, uint16, uint16) {
// 	// TODO: Switch to integer math.
// 	hipDeg := int16(math.Atan2(float64(toePt.X-l.hipPt.X), float64(toePt.Y-l.hipPt.Y)) + 0.5)
// 	var coxaDeg int16 = 90
// 	var tibiaDeg int16 = 90
// 	return l.servos[0].DegreesToMicros(hipDeg), l.servos[1].DegreesToMicros(coxaDeg), l.servos[2].DegreesToMicros(tibiaDeg)
// }
