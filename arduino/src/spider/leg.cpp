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

#include "leg.h"

#include <math.h>

Leg::Leg(Position leg_pos) {
    // The canonical zero position of the toe is with the coxa at "45 degrees",
    // the femur horizontal, and the tibia vertical. Therefore the hip joint is
    // displaced by (coxa+femur)/sqrt(2), using Pythagoras' theorem.
    float hip_offset = (coxa_len + femur_len) / sqrtf(2.0f);
    this->_hip_pt = {.x = hip_offset, .y = hip_offset, .z = tibia_len};
    this->_toe_pt = {0};

    switch (leg_pos) {
        case kFrontRight:
            this->_hip_pt.x = -this->_hip_pt.x;
            this->_hip_pt.y = -this->_hip_pt.y;
            break;
        case kFrontLeft:
            this->_hip_pt.y = -this->_hip_pt.y;
            break;
        case kBackRight:
            this->_hip_pt.x = -this->_hip_pt.x;
            break;
        case kBackLeft:
            // fallthrough
        default:
            break;
    }
}

void Leg::setToePoint(Point3D toe_pt) { this->_toe_pt = toe_pt; }

void Leg::getJointAngles(float *bc, float *cf, float *ft) {
    // Hip angle is measured counter-clockwise from a line projecting out from
    // the right side of the spider.
    *bc = atan2f(this->_toe_pt.y - this->_hip_pt.y,
                 this->_toe_pt.x - this->_hip_pt.x);

    // Total horizontal distance from hip to toe.
    float total_horiz_reach = sqrtf((this->_toe_pt.x - this->_hip_pt.x) *
                                        (this->_toe_pt.x - this->_hip_pt.x) +
                                    (this->_toe_pt.y - this->_hip_pt.y) *
                                        (this->_toe_pt.y - this->_hip_pt.y));
    // Femur+tibia horizontal reach.
    float ft_horiz_reach = total_horiz_reach - coxa_len;
    // Femur+tibia reach in 3D space.
    // This gives us a triangle with sides (femur_len, tibia_len, ftReach).
    float ft_diag_reach = sqrtf(ft_horiz_reach * ft_horiz_reach +
                                (this->_toe_pt.z - this->_hip_pt.z) *
                                    (this->_toe_pt.z - this->_hip_pt.z));

    // Solve for angles, using the law of cosines.
    //   c^2 = a^2 + b^2 - 2*a*b*cos(C)
    //   2*a*b*cos(C) =  a^2 + b^2 - c^2
    //   cos(C) = (a^2 + b^2 - c^2) / (2*a*b)
    // Or in coding terms:
    //   cos_num = a*a + b*b - c*c
    //   cos_denom = 2*a*b
    //   angle_c = acosf(cos_num / cos_denom)
    float cos_num;
    float cos_denom;

    // Coxa-Femur angle is measured counter-clockwise from horizontal, so up is
    // positive and down is negative. First, find the angle between the femur
    // and the imaginary line from the coxa-femur joint down to the  toe.
    cos_num = ft_diag_reach * ft_diag_reach + femur_len * femur_len -
              tibia_len * tibia_len;
    cos_denom = 2.0f * ft_diag_reach * femur_len;
    float femur_reach_angle = acosf(cos_num / cos_denom);

    // Second, find the angle between horizontal and the imaginary line from the
    // coxa-femur joint down to the  toe.
    float horiz_reach_angle =
        atan2f(this->_toe_pt.z - this->_hip_pt.z, ft_horiz_reach);
    *cf = femur_reach_angle + horiz_reach_angle;

    // Femur-Tibia angle is measured counter-clockwise from the femur, so it
    // will always be positive, and bigger numbers represent a further reach.
    cos_num = femur_len * femur_len + tibia_len * tibia_len -
              ft_diag_reach * ft_diag_reach;
    cos_denom = 2.0 * femur_len * tibia_len;
    *ft = acosf(cos_num / cos_denom);

    // Convert radians to degrees.
    *bc = *bc / M_PI * 180.0f;
    *cf = *cf / M_PI * 180.0f;
    *ft = *ft / M_PI * 180.0f;
}
