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

type Servo struct {
	name         string
	id           uint
	position     uint
	min_position uint
	max_position uint
}

func (s *Servo) SetAngleDegrees(angle int) {
	// TODO
	s.SetPosition(999)
}

func (s *Servo) SetPosition(position uint) {
	if position > s.max_position {
		s.position = s.max_position
	} else if position < s.min_position {
		s.position = s.min_position
	} else {
		s.position = position
	}
}
