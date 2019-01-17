#! /usr/bin/env python3

import math

class SpiderLeg:
    def __init__(self):
        self._coxa_len = 24
        self._femur_len = 38
        self._tibia_len = 80
        (self.x, self.y, self.z) = self.get_start_xyz()

    def get_start_xyz(self):
        xy_start = self._coxa_len + self._femur_len * 0.55
        xy_stand = -80
        return (xy_start, xy_start, xy_stand)

    def get_servo_angles(self):
        print(self.x, self.y, self.z)
        assert self.x >= 0 and self.x <= 100
        assert self.y >= 0 and self.y <= 100

        # Full horizontal distance (i.e., in the XY plane) from body to target.
        xy_total = math.sqrt(self.x * self.x + self.y * self.y)

        # Horizontal distance (i.e., in the XY plane) to reach from the end of the coxa.
        xy_reach = xy_total - self._coxa_len

        # Absolute distance from coxa to tibia tip.
        xyz_reach = math.sqrt(xy_reach * xy_reach + self.z * self.z)
        assert xyz_reach > 60 and xyz_reach < 110

        # We now have a triangle with sides [femur_len, tibia_len, xyz_reach].
        # Solve using the law of cosines.
        cos_alpha = (
            self._femur_len * self._femur_len
            + self._tibia_len * self._tibia_len
            - xyz_reach * xyz_reach
            ) / (
            2.0 * self._femur_len * self._tibia_len
            )
        alpha = math.degrees(math.acos(cos_alpha))

        # We now solve for the coxa-femur angle.
        b1 = math.atan2(self.z, xy_reach)
        cos_b2 = (
            self._femur_len * self._femur_len
            + xyz_reach * xyz_reach
            - self._tibia_len * self._tibia_len
            ) / (
            2.0 * self._femur_len * xyz_reach
            )
        b2 = math.acos(cos_b2)
        beta = 180 - math.degrees(b1 + b2)

        # Coxa angle.
        gamma = math.degrees(math.atan2(self.x, self.y))

        return (alpha, beta, gamma)

if __name__ == "__main__":
    leg = SpiderLeg()

    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.z = 50
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 80
    leg.y = 20
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 20
    leg.y = 80
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 80
    leg.y = 80
    leg.z = 0
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 10
    leg.y = 10
    leg.z = -100
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))
