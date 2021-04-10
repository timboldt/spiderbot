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
	"math"
	"testing"
)

func approxRadToDeg(r float64) int {
	return int(math.Round(r * 180.0 / math.Pi))
}

// Verify the hip location is in the right place relative to the canonical toe position of (0,0,0).
func TestInitLeg(t *testing.T) {
	var l Leg
	l.init(FrontRight)
	if l.hipPt.X >= 0 || l.hipPt.Y >= 0 || l.hipPt.Z <= 0 {
		t.Errorf("l.init(FrontRight) returned %v, expected hip location elsewhere", l.hipPt)
	}
	l.init(FrontLeft)
	if l.hipPt.X <= 0 || l.hipPt.Y >= 0 || l.hipPt.Z <= 0 {
		t.Errorf("l.init(FrontLeft) returned %v, expected hip location elsewhere", l.hipPt)
	}
	l.init(BackRight)
	if l.hipPt.X >= 0 || l.hipPt.Y <= 0 || l.hipPt.Z <= 0 {
		t.Errorf("l.init(BackRight) returned %v, expected hip location elsewhere", l.hipPt)
	}
	l.init(BackLeft)
	if l.hipPt.X <= 0 || l.hipPt.Y <= 0 || l.hipPt.Z <= 0 {
		t.Errorf("l.init(BackLeft) returned %v, expected hip location elsewhere", l.hipPt)
	}
}

func TestJointAnglesAtNullPoint(t *testing.T) {
	var l Leg
	var got, want int
	var bc, cf, ft float64

	for lp := LegPosition(0); lp <= LegPosition(3); lp++ {
		l.init(lp)
		toePt := Point3D{X: 0, Y: 0, Z: 0}
		bc, cf, ft = l.JointAngles(toePt)
		got = approxRadToDeg(bc)
		want = 45
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (%v, _, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(cf)
		want = 0
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, %v, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(ft)
		want = 90
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, _, %v), expected %v", lp, toePt, got, want)
		}
	}
}

func TestJointAnglesSideOfBody(t *testing.T) {
	var l Leg
	var got, want int
	var bc, cf, ft float64

	for lp := LegPosition(0); lp <= LegPosition(3); lp++ {
		l.init(lp)
		// To the side, pulled in a bit, and down.
		toePt := Point3D{X: l.hipPt.X / 2.0, Y: l.hipPt.Y, Z: -20}
		bc, cf, ft = l.JointAngles(toePt)
		got = approxRadToDeg(bc)
		want = 0
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (%v, _, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(cf)
		want = -42
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, %v, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(ft)
		want = 111
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, _, %v), expected %v", lp, toePt, got, want)
		}
	}
}

func TestJointAnglesAheadOrBehindBody(t *testing.T) {
	var l Leg
	var got, want int
	var bc, cf, ft float64

	for lp := LegPosition(0); lp <= LegPosition(3); lp++ {
		l.init(lp)
		// To the front (or back), stretched out a bit, and above the hip.
		toePt := Point3D{X: l.hipPt.X, Y: -2.0 / 3.0 * l.hipPt.Y, Z: l.hipPt.Z + 10}
		bc, cf, ft = l.JointAngles(toePt)
		got = approxRadToDeg(bc)
		want = 90
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (%v, _, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(cf)
		want = 145
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, %v, _), expected %v", lp, toePt, got, want)
		}
		got = approxRadToDeg(ft)
		want = 27
		if got != want {
			t.Errorf("%v.JointAngles(%v) returned (_, _, %v), expected %v", lp, toePt, got, want)
		}
	}
}
