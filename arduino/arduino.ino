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

#include <Adafruit_PWMServoDriver.h>
#include <Arduino.h>
#include <math.h>

#include "src/spider/spider.h"

Adafruit_PWMServoDriver pwm;
Spider spider;

void setup(void) {
    Serial.begin(115200);
    while (!Serial) {
        delay(10);
    }
    pwm.begin();
    pwm.setPWMFreq(50);  // Analog servos run at 50Hz.
}

void loop(void) {
    float theta = 0.0;
    for (int i = 0; i < 100; i++) {
        delay(25);
        spider.sendUpdatesToServos(&pwm, servo_config);
        theta += PI / 50.0;
        spider.setToePositionAbsolute(Leg::kFrontRight,
                                      {sinf(theta) * 20.0f, 0.0, 0.0});
        spider.setToePositionAbsolute(Leg::kFrontLeft,
                                      {sinf(theta) * 20.0f, 0.0, 0.0});
        spider.setToePositionAbsolute(Leg::kBackRight,
                                      {sinf(theta) * 20.0f, 0.0, 0.0});
        spider.setToePositionAbsolute(Leg::kBackLeft,
                                      {sinf(theta) * 20.0f, 0.0, 0.0});
    }
    for (int i = 0; i < 100; i++) {
        delay(25);
        spider.sendUpdatesToServos(&pwm, servo_config);
        theta += PI / 50.0;
        spider.setToePositionAbsolute(Leg::kFrontRight,
                                      {0.0, sinf(theta) * 20.0f, 0.0});
        spider.setToePositionAbsolute(Leg::kFrontLeft,
                                      {0.0, sinf(theta) * 20.0f, 0.0});
        spider.setToePositionAbsolute(Leg::kBackRight,
                                      {0.0, sinf(theta) * 20.0f, 0.0});
        spider.setToePositionAbsolute(Leg::kBackLeft,
                                      {0.0, sinf(theta) * 20.0f, 0.0});
    }
    for (int i = 0; i < 100; i++) {
        delay(25);
        spider.sendUpdatesToServos(&pwm, servo_config);
        theta += PI / 50.0;
        spider.setToePositionAbsolute(Leg::kFrontRight,
                                      {0.0, 0.0, sinf(theta) * 20.0f});
        spider.setToePositionAbsolute(Leg::kFrontLeft,
                                      {0.0, 0.0, sinf(theta) * 20.0f});
        spider.setToePositionAbsolute(Leg::kBackRight,
                                      {0.0, 0.0, sinf(theta) * 20.0f});
        spider.setToePositionAbsolute(Leg::kBackLeft,
                                      {0.0, 0.0, sinf(theta) * 20.0f});
    }
}