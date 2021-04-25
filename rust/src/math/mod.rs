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

pub fn fast_sqrt(n : i16) -> i8 {
    // A binary search of a lookup table is another option, but this algorithm
    // seems to complete in 30-40 clock cycles, which is pretty quick.
    assert!(n >= 0);
    if n < 2 {
        return n as i8;
    }
    let small = fast_sqrt(n >> 2) << 1;
    let big = small + 1;
    if big as i16 * big as i16 > n {
        small
    } else {
        big
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sanity() {
        for n in 0..16384 {
            assert_eq!(fast_sqrt(n), (n as f32).sqrt().floor() as i8);
        }
    }
}