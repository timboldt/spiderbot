#include "spider.h"

Spider::Spider()
    : _legs{Leg(Leg::kFrontRight), Leg(Leg::kFrontLeft), Leg(Leg::kBackRight),
             Leg(Leg::kBackLeft)} {}

void Spider::sendUpdatesToServos(Adafruit_PWMServoDriver *pwm, Servo *servos) {
    for (size_t leg = 0; leg < 4; leg++) {
        float bc;
        float cf;
        float ft;
        this->_legs[leg].getJointAngles(&bc, &cf, &ft);
        uint8_t servo;
        servo = servoID(Leg::Position(leg), Leg::kBodyCoxa);
        pwm->writeMicroseconds(servo, servos[servo].degreesToMicros(bc));
        servo = servoID(Leg::Position(leg), Leg::kCoxaFemur);
        pwm->writeMicroseconds(servo, servos[servo].degreesToMicros(cf));
        servo = servoID(Leg::Position(leg), Leg::kFemurTibia);
        pwm->writeMicroseconds(servo, servos[servo].degreesToMicros(ft));
    }
}

void Spider::setToePositionAbsolute(Leg::Position leg, Point3D pt) {
    this->_legs[leg].setToePoint(pt);
}

void Spider::setToePositionRelative(Leg::Position leg, Point3D vect) {
    this->_legs[leg].moveToePoint(vect);
}

uint8_t Spider::servoID(Leg::Position leg, Leg::Joint joint) {
    return uint8_t(leg) * 3 + uint8_t(joint);
}