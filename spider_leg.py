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

        # Full distance to reach.
        w = math.sqrt(pow(self.x, 2) + pow(self.y, 2))

        # Distance to reach from the end of the coxa.
        v = w - self._coxa_len

        alpha_tmp = (pow(self._femur_len, 2) - pow(self._tibia_len, 2) + pow(v, 2) + pow(self.z, 2)) / 2 / self._femur_len / math.sqrt(pow(v, 2) + pow(self.z, 2))
        if (alpha_tmp > 1 or alpha_tmp < -1):
            print("x=%f y=%f v=%f w=%f" % (self.x, self.y, v, w))
            print("alpha=%f" % alpha_tmp)
            if (alpha_tmp > 1):
                alpha_tmp = 1
            else:
                alpha_tmp = -1
        alpha = math.atan2(self.z, v) + math.acos(alpha_tmp)

        beta_tmp = (pow(self._femur_len, 2) + pow(self._tibia_len, 2) - pow(v, 2) - pow(self.z, 2)) / 2 / self._femur_len / self._tibia_len
        if (beta_tmp > 1 or beta_tmp < -1):
            print("x=%f y=%f v=%f w=%f" % (self.x, self.y, v, w))
            print("beta=%f" % beta_tmp)
            if (beta_tmp > 1):
                beta_tmp = 1
            else:
                beta_tmp = -1
        beta = math.acos(beta_tmp)

        # Coxa angle.
        gamma = math.atan2(self.y, self.x)
        if self._backwards:
            gamma = -gamma

        return (alpha*180/math.pi-50, beta*180/math.pi-50, gamma*180/math.pi)

if __name__ == "__main__":
    leg = SpiderLeg(backwards=False)
    print(leg.get_servo_angles())
    leg.z = 50
    print(leg.get_servo_angles())
    leg.x = 70
    leg.y = 20
    print(leg.get_servo_angles())
