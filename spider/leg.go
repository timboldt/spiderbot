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

const coxaLength = 24
const femurLength = 38
const tibiaLength = 80

type Leg struct {
	x, y, z float32
}

funct NewLeg() *Leg {
	return &Leg{
		x: coxaLength + femurLength * 0.55,
		y: coxaLength + femurLength * 0.55,
		z: 0,
	}
}


// def get_servo_angles(self):
// assert self.x >= 0 and self.x <= 100
// assert self.y >= 0 and self.y <= 100

// # Full horizontal distance (i.e., in the XY plane) from body to target.
// xy_total = math.sqrt(self.x * self.x + self.y * self.y)

// # Horizontal distance (i.e., in the XY plane) to reach from the end of the coxa.
// xy_reach = xy_total - self._coxa_len

// # Absolute distance from coxa to tibia tip.
// xyz_reach = math.sqrt(xy_reach * xy_reach + self.z * self.z)
// assert xyz_reach > 60 and xyz_reach < 110

// # We now have a triangle with sides [femur_len, tibia_len, xyz_reach].
// # Solve using the law of cosines.
// cos_alpha = (
// 	self._femur_len * self._femur_len
// 	+ self._tibia_len * self._tibia_len
// 	- xyz_reach * xyz_reach
// 	) / (
// 	2.0 * self._femur_len * self._tibia_len
// 	)
// alpha = math.degrees(math.acos(cos_alpha))

// # We now solve for the coxa-femur angle.
// b1 = math.atan2(self.z, xy_reach)
// cos_b2 = (
// 	self._femur_len * self._femur_len
// 	+ xyz_reach * xyz_reach
// 	- self._tibia_len * self._tibia_len
// 	) / (
// 	2.0 * self._femur_len * xyz_reach
// 	)
// b2 = math.acos(cos_b2)
// beta = 180 - math.degrees(b1 + b2)

// # Coxa angle.
// gamma = math.degrees(math.atan2(self.x, self.y))

// return (alpha, beta, gamma)