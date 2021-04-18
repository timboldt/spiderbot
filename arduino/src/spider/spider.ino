#include <AUnitVerbose.h>

#include "servo.h"

test(Servo, DegreesToMicros) {
    Servo s(800, 2200, 500, false);
    assertEqual(s.degreesToMicros(-10), 800);
    assertEqual(s.degreesToMicros(0), 800);
    assertEqual(s.degreesToMicros(10), 800);
    assertEqual(s.degreesToMicros(45), 1000);
    assertEqual(s.degreesToMicros(90), 1500);
    assertEqual(s.degreesToMicros(200), 2200);
}

test(Servo, ReverseDegreesToMicros) {
    Servo s(200, 2200, 1700, true);
    assertEqual(s.degreesToMicros(-90), 2200);
    assertEqual(s.degreesToMicros(-10), 1811);
    assertEqual(s.degreesToMicros(0), 1700);
    assertEqual(s.degreesToMicros(10), 1589);
    assertEqual(s.degreesToMicros(45), 1200);
    assertEqual(s.degreesToMicros(90), 700);
    assertEqual(s.degreesToMicros(200), 200);
}

void setup(void) {
    Serial.begin(115200);
    while (!Serial) delay(10);
}

void loop() { aunit::TestRunner::run(); }
