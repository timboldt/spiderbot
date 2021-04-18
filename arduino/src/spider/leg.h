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

#pragma once

#include <cstdint>

// Represents a 3D point in space.
// x is towards the right of the robot.
// y is towards the front of the robot.
// z is towards the top of the robot.
// Distances are in millimeters.
typedef struct {
    float x;
    float y;
    float z;
} Point3D;

// The leg represents a three-segment leg (coxa-femur-tibia), with the coxa
// joined to the body at the hip point and the tibia touching the floor at the
// toe point. The canonical toe point is (0,0,0), representing a leg that is
// pointing out from the body corner at a 45' angle, the femur horizontal and
// the tibia vertical. 
class Leg {
   public:
    enum Position {
        kFrontRight,
        kFrontLeft,
        kBackRight,
        kBackLeft,
    };

    enum Join {
        kBodyCoxa,
        kCoxaFemur,
        kFemurTibia,
    };

    Leg(Position leg_pos);

    void setToePoint(Point3D toe_pt);
    void getJointAngles(float *bc, float *cf, float *ft);

    static constexpr float coxa_len = 23.5f;
    static constexpr float femur_len = 38.0f;
    static constexpr float tibia_len = 81.0f;

   private:
    Point3D _hip_pt;
    Point3D _toe_pt;
};
