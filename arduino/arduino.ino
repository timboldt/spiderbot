#include <Arduino.h>

void setup(void) {
    Serial.begin(115200);
    while (!Serial) {
        delay(10);
    }
}

void loop(void) {}