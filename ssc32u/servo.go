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

package ssc32u

import "fmt"

const minPosition = 500
const maxPosition = 2500

type Servo struct {
	name       string
	id         uint
	position   uint
	isModified bool
}

func (s *Servo) SetAngleDegrees(angle float32) {
	if angle < 0 {
		angle = 0
	}
	s.SetPosition(uint(angle*2000.0/180.0 + 500.0))
}

func (s *Servo) SetPosition(position uint) {
	if position > maxPosition {
		s.position = maxPosition
	} else if position < minPosition {
		s.position = minPosition
	} else {
		s.position = position
	}
	s.isModified = true
}

func (s Servo) commandString() string {
	if !s.isModified {
		return ""
	}
	return fmt.Sprintf("#%d P%d ", s.id, s.position)
}
