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
    Spider();

    // sendUpdatesToServos calculates all servo angles and updates the PWM
    // device.
    // pwm must not be null.
    // servos must point to an array of 12 servos.
    void sendUpdatesToServos(Adafruit_PWMServoDriver *pwm, Servo *servos);

    void setToePositionAbsolute(Leg::Position leg, Point3D pt);
    
    void setToePositionRelative(Leg::Position leg, Point3D vect);

   private:
    uint8_t servoID(Leg::Position leg, Leg::Joint joint);

    Leg _legs[4];
};
