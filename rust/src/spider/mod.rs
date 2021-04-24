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

pub mod leg;
pub mod servo;

use leg::Leg;
use servo::Servo;

pub struct Spider {
    legs: [leg::Leg; 4],
    servos: [servo::Servo; 12],
}

impl Spider {
    pub fn new() -> Spider {
        Spider {
            legs: [
                Leg::new(leg::Position::FrontRight),
                Leg::new(leg::Position::FrontLeft),
                Leg::new(leg::Position::BackRight),
                Leg::new(leg::Position::BackLeft),
            ],
            // Robot-specific servo configuration.
            // This is specific to my robot. Override with appropriate values for your bot.
            servos: [
                // FR BC
                Servo::new(1500, 2500, 1700, false),
                // FR CF
                Servo::new(1200, 2600, 2111, true),
                // FR FT
                Servo::new(1400, 2500, 900, false),
                // FL BC
                Servo::new(700, 1700, -400, false),
                // FL CF
                Servo::new(500, 1900, 1045, false),
                // FL FT
                Servo::new(1300, 2400, 2800, true),
                // BR BC
                Servo::new(700, 1700, 1800, false),
                // BR CF
                Servo::new(700, 2100, 1189, false),
                // BR FT
                Servo::new(1500, 2500, 3100, true),
                // BL BC
                Servo::new(1400, 2400, 3500, false),
                // BL CF
                Servo::new(1000, 2200, 1600, true),
                // BL FT
                Servo::new(1100, 2200, 600, false),
            ],
        }
    }
}
