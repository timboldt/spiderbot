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

#include "servo.h"

#include <math.h>

Servo::Servo(uint16_t min_micros, uint16_t max_micros, int16_t zero_degrees,
             bool reversed)
    : _min_micros(min_micros),
      _max_micros(max_micros),
      _zero_degrees(zero_degrees),
      _reversed(reversed) {
    // TODO: Verify parameters?
}

uint16_t Servo::degreesToMicros(float deg) {
    if (this->_reversed) {
        deg = -deg;
    }
    int16_t micros = int16_t(roundf(deg * 100.0f / 9.0f)) + this->_zero_degrees;
    if (micros < this->_min_micros) {
        micros = this->_min_micros;
    }
    if (micros > this->_max_micros) {
        micros = this->_max_micros;
    }

    return (uint16_t)micros;
}