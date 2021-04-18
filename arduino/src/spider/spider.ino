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

#include <AUnitVerbose.h>

#include "leg.h"
#include "servo.h"
#include "spider.h"

test(Servo, DegreesToMicros) {
    Servo s(800, 2200, 500, false);
    assertEqual(s.degreesToMicros(-10), 800);
    assertEqual(s.degreesToMicros(0), 800);
    assertEqual(s.degreesToMicros(10), 800);
    assertEqual(s.degreesToMicros(45), 1000);
    assertEqual(s.degreesToMicros(90), 1500);
    assertEqual(s.degreesToMicros(200), 2200);
}

test(Servo, ReverseDegreesToMicros) {
    Servo s(200, 2200, 1700, true);
    assertEqual(s.degreesToMicros(-90), 2200);
    assertEqual(s.degreesToMicros(-10), 1811);
    assertEqual(s.degreesToMicros(0), 1700);
    assertEqual(s.degreesToMicros(10), 1589);
    assertEqual(s.degreesToMicros(45), 1200);
    assertEqual(s.degreesToMicros(90), 700);
    assertEqual(s.degreesToMicros(200), 200);
}

test(Leg, JointAnglesAtNullPoint) {
    float bc;
    float cf;
    float ft;
    {
        Leg l(Leg::kFrontRight);
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 45.0f, 0.25f);
        assertNear(cf, 0.0f, 0.25f);
        assertNear(ft, 90.0f, 0.25f);
    }
    {
        Leg l(Leg::kFrontLeft);
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 135.0f, 0.25f);
        assertNear(cf, 0.0f, 0.25f);
        assertNear(ft, 90.0f, 0.25f);
    }
    {
        Leg l(Leg::kBackRight);
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, -45.0f, 0.25f);
        assertNear(cf, 0.0f, 0.25f);
        assertNear(ft, 90.0f, 0.25f);
    }
    {
        Leg l(Leg::kBackLeft);
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, -135.0f, 0.25f);
        assertNear(cf, 0.0f, 0.25f);
        assertNear(ft, 90.0f, 0.25f);
    }
}

test(Leg, JointAnglesSideOfBody) {
    float hip_offset = (Leg::coxa_len + Leg::femur_len) / sqrtf(2.0f);

    float bc;
    float cf;
    float ft;
    {
        Leg l(Leg::kFrontRight);
        // To the side, pulled in a bit, and down.
        l.setToePoint(Point3D{x : -hip_offset / 2.0f, y : -hip_offset, -20.0f});
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 0.0f, 0.25f);
        assertNear(cf, -42.5f, 0.25f);
        assertNear(ft, 111.0f, 0.25f);
    }
    {
        Leg l(Leg::kFrontLeft);
        // To the side, pulled in a bit, and down.
        l.setToePoint(Point3D{x : hip_offset / 2.0f, y : -hip_offset, -20.0f});
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 180.0f, 0.25f);
        assertNear(cf, -42.5f, 0.25f);
        assertNear(ft, 111.0f, 0.25f);
    }
    {
        Leg l(Leg::kBackRight);
        // To the side, pulled in a bit, and down.
        l.setToePoint(Point3D{x : -hip_offset / 2.0f, y : hip_offset, -20.0f});
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 0.0f, 0.25f);
        assertNear(cf, -42.5f, 0.25f);
        assertNear(ft, 111.0f, 0.25f);
    }
    {
        Leg l(Leg::kBackLeft);
        // To the side, pulled in a bit, and down.
        l.setToePoint(Point3D{x : hip_offset / 2.0f, y : hip_offset, -20.0f});
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 180.0f, 0.25f);
        assertNear(cf, -42.5f, 0.25f);
        assertNear(ft, 111.0f, 0.25f);
    }
}

test(Leg, JointAnglesAheadOrBehindBody) {
    float hip_offset = (Leg::coxa_len + Leg::femur_len) / sqrtf(2.0f);

    float bc;
    float cf;
    float ft;
    {
        Leg l(Leg::kFrontRight);
        // To the front (or back), stretched out a bit, and above the hip.
        l.setToePoint(Point3D{
            x : -hip_offset,
            y : 2.0f / 3.0f * hip_offset,
            Leg::tibia_len + 10.0f
        });
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 90.0f, 0.25f);
        assertNear(cf, 145.0f, 0.25f);
        assertNear(ft, 26.5f, 0.25f);
    }
    {
        Leg l(Leg::kFrontLeft);
        // To the front (or back), stretched out a bit, and above the hip.
        l.setToePoint(Point3D{
            x : hip_offset,
            y : 2.0f / 3.0f * hip_offset,
            Leg::tibia_len + 10.0f
        });
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, 90.0f, 0.25f);
        assertNear(cf, 145.0f, 0.25f);
        assertNear(ft, 26.5f, 0.25f);
    }
    {
        Leg l(Leg::kBackRight);
        // To the front (or back), stretched out a bit, and above the hip.
        l.setToePoint(Point3D{
            x : -hip_offset,
            y : -2.0f / 3.0f * hip_offset,
            Leg::tibia_len + 10.0f
        });
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, -90.0f, 0.25f);
        assertNear(cf, 145.0f, 0.25f);
        assertNear(ft, 26.5f, 0.25f);
    }
    {
        Leg l(Leg::kBackLeft);
        // To the front (or back), stretched out a bit, and above the hip.
        l.setToePoint(Point3D{
            x : hip_offset,
            y : -2.0f / 3.0f * hip_offset,
            Leg::tibia_len + 10.0f
        });
        l.getJointAngles(&bc, &cf, &ft);
        assertNear(bc, -90.0f, 0.25f);
        assertNear(cf, 145.0f, 0.25f);
        assertNear(ft, 26.5f, 0.25f);
    }
}

void setup(void) {
    Serial.begin(115200);
    while (!Serial) delay(10);
}

void loop() { aunit::TestRunner::run(); }
