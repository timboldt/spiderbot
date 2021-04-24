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

mod leg;
mod servo;

fn main() {
    let s = servo::Servo {
        min_micros: 800,
        max_micros: 2200,
        zero_degrees: 500,
        reversed: false,
    };
    println!("{}", s.degrees_to_micros(10.0));

    let mut legs = vec![
        leg::Leg::new(leg::Position::FrontRight),
        leg::Leg::new(leg::Position::FrontLeft),
        leg::Leg::new(leg::Position::BackRight),
        leg::Leg::new(leg::Position::BackLeft),
    ];
    legs[0].set_toe_point(10.0, 20.0, 30.0);
    legs[1].move_toe_point(1.0, 1.0, 1.0);
    for leg in legs {
        let (a, b, c) = leg.get_joint_angles();
        println!("{} {} {}", a, b, c);
    }
}
