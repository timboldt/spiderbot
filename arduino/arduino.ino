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

#include "src/spider/spider.h"

Adafruit_PWMServoDriver pwm;
Spider spider(&pwm, std::move(servo_config));

void setup(void) {
    Serial.begin(115200);
    while (!Serial) {
        delay(10);
    }
    pwm.begin();
    pwm.setPWMFreq(50);  // Analog servos run at 50Hz.
    spider.sendUpdatesToServos();
}

void loop(void) {
    // for (int i = 0; i < 12; i++) {
    //     pwm.writeMicroseconds(i, 1500);
    // }
}