#! /usr/bin/env python3

import math
import spider_leg as leg
import ssc32
import time

class SpiderBody:
    def __init__(self, ssc):
        self._ssc = ssc
        self.legs = [leg.SpiderLeg(), leg.SpiderLeg(), leg.SpiderLeg(), leg.SpiderLeg()]
        self._tibia_servos = [0, 15, 31, 16]
        self._femur_servos = [1, 14, 30, 17]
        self._coxa_servos = [2, 13, 29, 18]

    def _commit(self, millis):
        for i in range(4):
            a, b, g = self.legs[i].get_servo_angles()
            if i % 2 == 1:
                self._ssc[self._tibia_servos[i]].degrees = a - 90
                self._ssc[self._femur_servos[i]].degrees = b - 180
                self._ssc[self._coxa_servos[i]].degrees = 90 - g
            else:
                self._ssc[self._tibia_servos[i]].degrees = 90 - a
                self._ssc[self._femur_servos[i]].degrees = 180 - b
                self._ssc[self._coxa_servos[i]].degrees = g - 90
        self._ssc.commit(millis)
        while not self._ssc.is_done():
            time.sleep(0.1)
    
    def stand_up(self):
        for i in range(4):
            (self.legs[i].x, self.legs[i].y, self.legs[i].z) = self.legs[i].get_start_xyz()
        self._commit(1000)

    def lift_leg(self, id, mm=30):
        self.legs[id].z += mm
        self._commit(1000)

    def shift_weight_off_leg(self, id, mm=20):
        if i % 2 == 1:
            mm = -mm
        self.legs[id].x += mm
        self.legs[id].y += mm
        self.legs[(id+1)%4].x += mm * 0.7
        self.legs[(id+1)%4].y -= mm * 0.7
        self.legs[(id+2)%4].x -= mm
        self.legs[(id+2)%4].y -= mm
        self.legs[(id+3)%4].x -= mm * 0.7
        self.legs[(id+3)%4].y += mm * 0.7
        self._commit(1000)

if __name__ == "__main__":
    body = SpiderBody(ssc32.SSC32('/dev/ttyS6', 9600, count=32))
    body.stand_up()
    for i in range(4):
        if i % 2 == 0:
            body.shift_weight_off_leg(i)
            body.lift_leg(i)
            body.stand_up()