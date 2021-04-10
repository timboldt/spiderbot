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
	BodyCoxa Joint = iota
	CoxaFemur
	FemurTibia
)

const (
	CoxaLength  = 23.5
	FemurLength = 38.0
	TibiaLength = 81.0
)

// Represents a 3D point in space.
// X is towards the right of the robot.
// Y is towards the front of the robot.
// Z is towards the top of the robot.
// Distances are expressed in 1/32 of a millimeter.
type Point3D struct {
	X, Y, Z float64
}

type Leg struct {
	hipPt Point3D
}

func (l *Leg) init(pos LegPosition) {
	// The canonical zero position of the toe is with the coxa at "45 degrees", the femur horizontal, and the tibia vertical.
	// Therefore the hip joint is displaced by (coxa+femur)/sqrt(2), using Pythagoras' theorem.
	hipOffset := (CoxaLength + FemurLength) * math.Sqrt(2)
	var hipX, hipY float64
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

func (l *Leg) JointAngles(toePt Point3D) (float64, float64, float64) {
	// Hip angle is measured counter-clockwise from a line projecting out from the side of the spider, so FrontLeft/BackRight angles are negative.
	bodyCoxaAngle := math.Atan2(toePt.X-l.hipPt.X, toePt.Y-l.hipPt.Y)

	// Total horizontal distance from hip to toe.
	horizReach := math.Sqrt((toePt.X-l.hipPt.X)*(toePt.X-l.hipPt.X) + (toePt.Y-l.hipPt.Y)*(toePt.Y-l.hipPt.Y))
	// Femur+tibia horizontal reach.
	ftHorizReach := horizReach - CoxaLength
	// Femur+tibia reach in 3D space.
	// This gives us a triangle with sides (FemurLength, TibiaLength, ftReach).
	ftReach := math.Sqrt(ftHorizReach*ftHorizReach + (toePt.Z-l.hipPt.Z)*(toePt.Z-l.hipPt.Z))

	// Femur-Tibia angle is measured counter-clockwise from the femur, so it will always be positive, and bigger numbers represent a further reach.
	// Solve using the law of cosines.
	// c^2 = a^2 + b^2 - 2*a*b*cos(C)
	// 2*a*b*cos(C) =  a^2 + b^2 - c^2
	// cos(C) = (a^2 + b^2 - c^2) / (2*a*b)
	ftNum := FemurLength*FemurLength + TibiaLength*TibiaLength - ftReach*ftReach
	ftDenom := 2.0 * FemurLength * TibiaLength
	femurTibiaAngle := math.Acos(ftNum / ftDenom)

	// Coxa-Femur angle is measured counter-clockwise from horizontal, so up is positive and down is negative.
	// angle1 := math.Atan2(l.z, horizReachFromCoxa)
	// angle2 := math.Acos((femurLength*femurLength +
	// 	absoluteReachFromCoxa*absoluteReachFromCoxa -
	// 	tibiaLength*tibiaLength) /
	// 	(2.0 * femurLength * absoluteReachFromCoxa))
	// knee := 180 - (angle1+angle2)*180/math.Pi
	coxaFemurAngle := math.Pi / 2.0
	return bodyCoxaAngle, coxaFemurAngle, femurTibiaAngle
}
