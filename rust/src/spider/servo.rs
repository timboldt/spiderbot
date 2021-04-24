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

pub struct Servo {
    pub min_micros: u16,
    pub max_micros: u16,
    pub zero_degrees: i16,
    pub reversed: bool,
}

impl Servo {
    pub fn new(min: u16, max: u16, zero: i16, reversed: bool) -> Servo {
        Servo {
            min_micros: min,
            max_micros: max,
            zero_degrees: zero,
            reversed: reversed,
        }
    }

    pub fn degrees_to_micros(&self, degrees: f32) -> u16 {
        let micros_per_deg = 100.0 / 9.0;
        let micros = if self.reversed {
            f32::round(-micros_per_deg * degrees) as i16
        } else {
            f32::round(micros_per_deg * degrees) as i16
        };
        core::cmp::max(
            core::cmp::min(
                core::cmp::max(micros + self.zero_degrees, 0) as u16,
                self.max_micros,
            ),
            self.min_micros,
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_degrees_to_micros() {
        let s = Servo {
            min_micros: 800,
            max_micros: 2200,
            zero_degrees: 500,
            reversed: false,
        };
        assert_eq!(s.degrees_to_micros(-10.0), 800);
        assert_eq!(s.degrees_to_micros(0.0), 800);
        assert_eq!(s.degrees_to_micros(10.0), 800);
        assert_eq!(s.degrees_to_micros(45.0), 1000);
        assert_eq!(s.degrees_to_micros(90.0), 1500);
        assert_eq!(s.degrees_to_micros(200.0), 2200);
    }

    #[test]
    fn test_reverse_degrees_to_micros() {
        let s = Servo {
            min_micros: 200,
            max_micros: 2200,
            zero_degrees: 1700,
            reversed: true,
        };
        assert_eq!(s.degrees_to_micros(-90.0), 2200);
        assert_eq!(s.degrees_to_micros(-10.0), 1811);
        assert_eq!(s.degrees_to_micros(0.0), 1700);
        assert_eq!(s.degrees_to_micros(10.0), 1589);
        assert_eq!(s.degrees_to_micros(45.0), 1200);
        assert_eq!(s.degrees_to_micros(90.0), 700);
        assert_eq!(s.degrees_to_micros(200.0), 200);
    }
}
