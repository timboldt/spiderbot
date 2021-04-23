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

use micromath::F32Ext;

// Represents a 3D point in space.
// x is towards the right of the robot.
// y is towards the front of the robot.
// z is towards the top of the robot.
// Distances are in millimeters.
struct Point3D(f32, f32, f32);

pub enum Position {
    FrontRight,
    FrontLeft,
    BackRight,
    BackLeft,
}

enum Joint {
    BodyCoxa,
    CoxaFemur,
    FemurTibia,
}

const COXA_LEN: f32 = 23.5;
const FEMUR_LEN: f32 = 38.0;
const TIBIA_LEN: f32 = 81.0;

// The leg represents a three-segment leg (coxa-femur-tibia), with the coxa
// joined to the body at the hip point and the tibia touching the floor at the
// toe point. The canonical toe point is (0,0,0), representing a leg that is
// pointing out from the body corner at a 45' angle, the femur horizontal and
// the tibia vertical.
pub struct Leg {
    hip_pt: Point3D,
    toe_pt: Point3D,
}

impl Leg {
    pub fn new(leg_pos: Position) -> Leg {
        // The canonical zero position of the toe is with the coxa at "45 degrees",
        // the femur horizontal, and the tibia vertical. Therefore the hip joint is
        // displaced by (coxa+femur)/sqrt(2), using Pythagoras' theorem.
        let hip_offset = (COXA_LEN + FEMUR_LEN) / F32Ext::sqrt(2f32);
        Leg {
            hip_pt: match leg_pos {
                Position::FrontRight => Point3D(-hip_offset, -hip_offset, TIBIA_LEN),
                Position::FrontLeft => Point3D(hip_offset, -hip_offset, TIBIA_LEN),
                Position::BackRight => Point3D(-hip_offset, hip_offset, TIBIA_LEN),
                Position::BackLeft => Point3D(hip_offset, hip_offset, TIBIA_LEN),
            },
            toe_pt: Point3D(0f32, 0f32, 0f32),
        }
    }

    pub fn set_toe_point(&mut self, x: f32, y: f32, z: f32) {
        self.toe_pt = Point3D(x, y, z);
    }

    pub fn move_toe_point(&mut self, x: f32, y: f32, z: f32) {
        self.toe_pt = Point3D(x + self.toe_pt.0, y + self.toe_pt.1, z + self.toe_pt.2);
    }

    // void Leg::getJointAngles(float *bc, float *cf, float *ft) {
    //     // Hip angle is measured counter-clockwise from a line projecting out from
    //     // the right side of the spider.
    //     *bc = atan2f(this->_toe_pt.y - this->_hip_pt.y,
    //                  this->_toe_pt.x - this->_hip_pt.x);

    //     // Total horizontal distance from hip to toe.
    //     float total_horiz_reach = sqrtf((this->_toe_pt.x - this->_hip_pt.x) *
    //                                         (this->_toe_pt.x - this->_hip_pt.x) +
    //                                     (this->_toe_pt.y - this->_hip_pt.y) *
    //                                         (this->_toe_pt.y - this->_hip_pt.y));
    //     // Femur+tibia horizontal reach.
    //     float ft_horiz_reach = total_horiz_reach - coxa_len;
    //     // Femur+tibia reach in 3D space.
    //     // This gives us a triangle with sides (femur_len, tibia_len, ftReach).
    //     float ft_diag_reach = sqrtf(ft_horiz_reach * ft_horiz_reach +
    //                                 (this->_toe_pt.z - this->_hip_pt.z) *
    //                                     (this->_toe_pt.z - this->_hip_pt.z));

    //     // Solve for angles, using the law of cosines.
    //     //   c^2 = a^2 + b^2 - 2*a*b*cos(C)
    //     //   2*a*b*cos(C) =  a^2 + b^2 - c^2
    //     //   cos(C) = (a^2 + b^2 - c^2) / (2*a*b)
    //     // Or in coding terms:
    //     //   cos_num = a*a + b*b - c*c
    //     //   cos_denom = 2*a*b
    //     //   angle_c = acosf(cos_num / cos_denom)
    //     float cos_num;
    //     float cos_denom;

    //     // Coxa-Femur angle is measured counter-clockwise from horizontal, so up is
    //     // positive and down is negative. First, find the angle between the femur
    //     // and the imaginary line from the coxa-femur joint down to the  toe.
    //     cos_num = ft_diag_reach * ft_diag_reach + femur_len * femur_len -
    //               tibia_len * tibia_len;
    //     cos_denom = 2.0f * ft_diag_reach * femur_len;
    //     float femur_reach_angle = acosf(cos_num / cos_denom);

    //     // Second, find the angle between horizontal and the imaginary line from the
    //     // coxa-femur joint down to the  toe.
    //     float horiz_reach_angle =
    //         atan2f(this->_toe_pt.z - this->_hip_pt.z, ft_horiz_reach);
    //     *cf = femur_reach_angle + horiz_reach_angle;

    //     // Femur-Tibia angle is measured counter-clockwise from the femur, so it
    //     // will always be positive, and bigger numbers represent a further reach.
    //     cos_num = femur_len * femur_len + tibia_len * tibia_len -
    //               ft_diag_reach * ft_diag_reach;
    //     cos_denom = 2.0 * femur_len * tibia_len;
    //     *ft = acosf(cos_num / cos_denom);

    //     // Convert radians to degrees.
    //     *bc = *bc / M_PI * 180.0f;
    //     *cf = *cf / M_PI * 180.0f;
    //     *ft = *ft / M_PI * 180.0f;
    // }
}

// test(Leg, JointAnglesAtNullPoint) {
//     float bc;
//     float cf;
//     float ft;
//     {
//         Leg l(Leg::kFrontRight);
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 45.0f, 0.25f);
//         assertNear(cf, 0.0f, 0.25f);
//         assertNear(ft, 90.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kFrontLeft);
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 135.0f, 0.25f);
//         assertNear(cf, 0.0f, 0.25f);
//         assertNear(ft, 90.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackRight);
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, -45.0f, 0.25f);
//         assertNear(cf, 0.0f, 0.25f);
//         assertNear(ft, 90.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackLeft);
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, -135.0f, 0.25f);
//         assertNear(cf, 0.0f, 0.25f);
//         assertNear(ft, 90.0f, 0.25f);
//     }
// }

// test(Leg, JointAnglesSideOfBody) {
//     float hip_offset = (Leg::coxa_len + Leg::femur_len) / sqrtf(2.0f);

//     float bc;
//     float cf;
//     float ft;
//     {
//         Leg l(Leg::kFrontRight);
//         // To the side, pulled in a bit, and down.
//         l.setToePoint(Point3D{x : -hip_offset / 2.0f, y : -hip_offset, -20.0f});
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 0.0f, 0.25f);
//         assertNear(cf, -42.5f, 0.25f);
//         assertNear(ft, 111.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kFrontLeft);
//         // To the side, pulled in a bit, and down.
//         l.setToePoint(Point3D{x : hip_offset / 2.0f, y : -hip_offset, -20.0f});
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 180.0f, 0.25f);
//         assertNear(cf, -42.5f, 0.25f);
//         assertNear(ft, 111.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackRight);
//         // To the side, pulled in a bit, and down.
//         l.setToePoint(Point3D{x : -hip_offset / 2.0f, y : hip_offset, -20.0f});
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 0.0f, 0.25f);
//         assertNear(cf, -42.5f, 0.25f);
//         assertNear(ft, 111.0f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackLeft);
//         // To the side, pulled in a bit, and down.
//         l.setToePoint(Point3D{x : hip_offset / 2.0f, y : hip_offset, -20.0f});
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 180.0f, 0.25f);
//         assertNear(cf, -42.5f, 0.25f);
//         assertNear(ft, 111.0f, 0.25f);
//     }
// }

// test(Leg, JointAnglesAheadOrBehindBody) {
//     float hip_offset = (Leg::coxa_len + Leg::femur_len) / sqrtf(2.0f);

//     float bc;
//     float cf;
//     float ft;
//     {
//         Leg l(Leg::kFrontRight);
//         // To the front (or back), stretched out a bit, and above the hip.
//         l.setToePoint(Point3D{
//             x : -hip_offset,
//             y : 2.0f / 3.0f * hip_offset,
//             Leg::tibia_len + 10.0f
//         });
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 90.0f, 0.25f);
//         assertNear(cf, 145.0f, 0.25f);
//         assertNear(ft, 26.5f, 0.25f);
//     }
//     {
//         Leg l(Leg::kFrontLeft);
//         // To the front (or back), stretched out a bit, and above the hip.
//         l.setToePoint(Point3D{
//             x : hip_offset,
//             y : 2.0f / 3.0f * hip_offset,
//             Leg::tibia_len + 10.0f
//         });
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, 90.0f, 0.25f);
//         assertNear(cf, 145.0f, 0.25f);
//         assertNear(ft, 26.5f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackRight);
//         // To the front (or back), stretched out a bit, and above the hip.
//         l.setToePoint(Point3D{
//             x : -hip_offset,
//             y : -2.0f / 3.0f * hip_offset,
//             Leg::tibia_len + 10.0f
//         });
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, -90.0f, 0.25f);
//         assertNear(cf, 145.0f, 0.25f);
//         assertNear(ft, 26.5f, 0.25f);
//     }
//     {
//         Leg l(Leg::kBackLeft);
//         // To the front (or back), stretched out a bit, and above the hip.
//         l.setToePoint(Point3D{
//             x : hip_offset,
//             y : -2.0f / 3.0f * hip_offset,
//             Leg::tibia_len + 10.0f
//         });
//         l.getJointAngles(&bc, &cf, &ft);
//         assertNear(bc, -90.0f, 0.25f);
//         assertNear(cf, 145.0f, 0.25f);
//         assertNear(ft, 26.5f, 0.25f);
//     }
// }
