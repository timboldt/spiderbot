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

const (
	BackRight = iota
	FrontRight
	FrontLeft
	BackLeft
)

type Body struct {
	legs []*Leg
}

func NewBody() *Body {
	legs := make([]*Leg, 4)
	for i := 0; i < 4; i++ {
		legs[i] = NewLeg()
	}
	return &Body{
		legs: legs,
	}
}

// Move a leg in the body reference frame.
// X is rightward, Y is forward, Z is upward.
func (b *Body) MoveLegAbsolute(id int, x, y, z float64) {
	switch id {
	case BackRight:
		b.legs[id].SetAbsolutePos(x, -y, z)
	case FrontRight:
		b.legs[id].SetAbsolutePos(x, y, z)
	case FrontLeft:
		b.legs[id].SetAbsolutePos(-x, y, z)
	case BackLeft:
		b.legs[id].SetAbsolutePos(-x, -y, z)
	}
}

func (b Body) GetLeg(id int) *Leg {
	return b.legs[id]
}
