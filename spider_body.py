#! /usr/bin/env python3

import math
import spider_leg as leg
import ssc32
import time

# Legs are:
#   0 - Back right
#   1 - Front right
#   2 - Front left
#   3 - Back left
class SpiderBody:
    def __init__(self, ssc):
        self._ssc = ssc
        self.legs = [leg.SpiderLeg(), leg.SpiderLeg(), leg.SpiderLeg(), leg.SpiderLeg()]
        self._tibia_servos = [0, 15, 31, 16]
        self._femur_servos = [1, 14, 30, 17]
        self._coxa_servos = [2, 13, 29, 18]
        self.BACK_RIGHT = 0
        self.FRONT_RIGHT = 1
        self.FRONT_LEFT = 2
        self.BACK_LEFT = 3
        self.commit_speed = 10

    def _commit(self, millis):
        for i in range(4):
            a, b, g = self.legs[i].get_servo_angles()
            if i == self.BACK_LEFT or i == self.FRONT_RIGHT:
                self._ssc[self._tibia_servos[i]].degrees = a - 90
                self._ssc[self._femur_servos[i]].degrees = b - 180
                self._ssc[self._coxa_servos[i]].degrees = 90 - g
            else:
                self._ssc[self._tibia_servos[i]].degrees = 90 - a
                self._ssc[self._femur_servos[i]].degrees = 180 - b
                self._ssc[self._coxa_servos[i]].degrees = g - 90
        self._ssc.commit(millis*self.commit_speed)
        while not self._ssc.is_done():
            time.sleep(0.1)
    
    def stand_up(self):
        for i in range(4):
            (self.legs[i].x, self.legs[i].y, self.legs[i].z) = self.legs[i].get_start_xyz()
        self._commit(100)

    def sit_down(self):
        for i in range(4):
            (self.legs[i].x, self.legs[i].y, self.legs[i].z) = self.legs[i].get_start_xyz()
            self.legs[i].z += 30
        self._commit(100)

    def leg_up(self, id, mm=20):
        self.legs[id].z += mm
        self._commit(100)

    def leg_down(self, id):
        (_, _, self.legs[id].z) = self.legs[id].get_start_xyz()
        self._commit(100)

    def shift_body(self, x_mm, y_mm):
        self.legs[self.FRONT_LEFT].x += x_mm
        self.legs[self.FRONT_LEFT].y -= y_mm
        self.legs[self.FRONT_RIGHT].x -= x_mm
        self.legs[self.FRONT_RIGHT].y -= y_mm
        self.legs[self.BACK_LEFT].x += x_mm
        self.legs[self.BACK_LEFT].y += y_mm
        self.legs[self.BACK_RIGHT].x -= x_mm
        self.legs[self.BACK_RIGHT].y += y_mm
        self._commit(100)

    def shift_weight_off_leg(self, id, mm=20):
        if id == self.FRONT_LEFT:
            self.shift_body(mm * 0.7, mm * -0.7)
        elif id == self.FRONT_RIGHT:
            self.shift_body(mm * -0.7, mm * -0.7)
        elif id == self.BACK_LEFT:
            self.shift_body(mm * 0.7, mm * 0.7)
        else: # BACK_RIGHT
            self.shift_body(mm * -0.7, mm * 0.7)
    
    def walk(self):
        # https://makezine.com/2016/11/22/robot-quadruped-arduino-program/
        self.stand_up()
        first_time = True

        for _ in range(3):
            self.shift_weight_off_leg(self.BACK_RIGHT, 20)
            self.leg_up(self.BACK_RIGHT)
            self.legs[self.BACK_RIGHT].y -= 20
            self.leg_down(self.BACK_RIGHT)
            self.shift_weight_off_leg(self.BACK_RIGHT, -20)

            self.shift_weight_off_leg(self.FRONT_RIGHT, 20)
            self.leg_up(self.FRONT_RIGHT)
            if first_time:
                first_time = False
                self.legs[self.FRONT_RIGHT].y += 20
            else:
                self.legs[self.FRONT_RIGHT].y += 40
            self.leg_down(self.FRONT_RIGHT)
            self.shift_weight_off_leg(self.FRONT_RIGHT, -20)

            self.shift_body(0, 20)

            self.shift_weight_off_leg(self.BACK_LEFT, 20)
            self.leg_up(self.BACK_LEFT)
            self.legs[self.BACK_LEFT].y -= 20
            self.leg_down(self.BACK_LEFT)
            self.shift_weight_off_leg(self.BACK_LEFT, -20)

            self.shift_weight_off_leg(self.FRONT_LEFT, 20)
            self.leg_up(self.FRONT_LEFT)
            self.legs[self.FRONT_LEFT].y += 40
            self.leg_down(self.FRONT_LEFT)
            self.shift_weight_off_leg(self.FRONT_LEFT, -20)

            self.shift_body(0, 20)

        # TODO: put right legs back in position