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

#include <Adafruit_PWMServoDriver.h>

#include <cstdint>

#include "leg.h"
#include "servo.h"

class Spider {
   public:
    Spider(Adafruit_PWMServoDriver *pwm, Servo *_servos);

    void sendUpdatesToServos();
    void setToePositionAbsolute(Leg::Position leg, Point3D pt);
    void setToePositionRelative(Leg::Position leg, Point3D vect);

   private:
    uint8_t servoID(Leg::Position leg, Leg::Joint joint);

    Adafruit_PWMServoDriver *const _pwm;  // Does not own.
    Leg _legs[4];
    Servo *const _servos;  // Does not own.
};
