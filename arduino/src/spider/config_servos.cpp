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

#include "servo.h"

// Robot-specific servo configuration.
// This is specific to my robot. Override with appropriate values for your bot.
Servo servo_config[12] = {
    // FR BC
    Servo(1500, 2500, 1700, false),
    // FR CF
    Servo(1200, 2600, 2111, true),
    // FR FT
    Servo(1400, 2500, 900, false),
    // FL BC
    Servo(700, 1700, -400, false),
    // FL CF
    Servo(500, 1900, 1045, false),
    // FL FT
    Servo(1300, 2400, 2800, true),
    // BR BC
    Servo(700, 1700, 1800, false),
    // BR CF
    Servo(700, 2100, 1189, false),
    // BR FT
    Servo(1500, 2500, 3100, true),
    // BL BC
    Servo(1400, 2400, 3500, false),
    // BL CF
    Servo(1000, 2200, 1600, true),
    // BL FT
    Servo(1100, 2200, 600, false),
};