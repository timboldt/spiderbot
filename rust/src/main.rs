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
//#![no_std]

//use core::panic::PanicInfo;

mod leg;
mod servo;

fn main() {
    let s = servo::Servo {
        min_micros: 800,
        max_micros: 2200,
        zero_degrees: 500,
        reversed: false,
    };
    let mut l = leg::Leg::new(leg::Position::FrontRight);
    l.set_toe_point(10.0, 20.0, 30.0);
    println!("{}", s.degrees_to_micros(10.0));
    //println!("{}", l.hip_pt.0);
}

// #[panic_handler]
// fn panic(_info: &PanicInfo) -> ! {
//     loop {}
// }
