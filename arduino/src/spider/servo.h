#pragma once

#include <Arduino.h>

class Servo {
   public:
    Servo(uint16_t min_micros, uint16_t max_micros, int16_t zero_degrees,
          bool reversed);
    uint16_t degreesToMicros(float deg);

   private:
    uint16_t _min_micros;
    uint16_t _max_micros;
    int16_t _zero_degrees;
    bool _reversed;
};