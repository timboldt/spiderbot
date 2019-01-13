#! /usr/bin/env python

import math

class SpiderLeg:
    def __init__(self, backwards):
        self._coxa_len = 24
        self._femur_len = 38
        self._tibia_len = 80
        self._backwards = backwards

        xy_start = self._coxa_len + self._femur_len * 0.55
        (self.x, self.y, self.z) = (xy_start, xy_start, -50)

    def get_servo_angles(self):
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
        alpha = math.acos(cos_alpha) * 180.0 / math.pi - 90

        # beta_tmp = (pow(self._femur_len, 2) + pow(self._tibia_len, 2) - pow(v, 2) - pow(self.z, 2)) / 2 / self._femur_len / self._tibia_len
        # if (beta_tmp > 1 or beta_tmp < -1):
        #     print("x=%f y=%f v=%f w=%f" % (self.x, self.y, v, w))
        #     print("beta=%f" % beta_tmp)
        #     if (beta_tmp > 1):
        #         beta_tmp = 1
        #     else:
        #         beta_tmp = -1
        # beta = math.acos(beta_tmp)
        beta = 0

        # Coxa angle.
        gamma = math.atan2(self.y, self.x) * 180.0 / math.pi
        if self._backwards:
            alpha = -alpha
            gamma = -gamma

        return (alpha, beta, gamma)

if __name__ == "__main__":
    leg = SpiderLeg(backwards=False)
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.z = 50
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 70
    leg.y = 20
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))

    leg.x = 80
    leg.y = 80
    a,b,g = leg.get_servo_angles()
    print("x=%f y=%f z=%f, a=%f, b=%f, g=%f" % (leg.x, leg.y, leg.z, a, b, g))
