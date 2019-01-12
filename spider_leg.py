import math

class SpiderLeg:
    def __init__(self):
        self._coxa_len = 27.5
        self._femur_len = 70
        self._tibia_len = 80

        xy_start = self._coxa_len + self._femur_len * 0.55
        (self.x, self.y, self.z) = (xy_start, xy_start, -50)

    def get_servo_angles(self):
        if (self.x >= 0):
            w = math.sqrt(pow(self.x, 2) + pow(self.y, 2))
        else:
            w = -1 * (math.sqrt(pow(self.x, 2) + pow(self.y, 2)))

        v = w - self._coxa_len
        alpha_tmp = (pow(self._femur_len, 2) - pow(self._tibia_len, 2) + pow(v, 2) + pow(self.z, 2)) / 2 / self._femur_len / math.sqrt(
            pow(v, 2) + pow(self.z, 2))
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

        if (w >= 0):
            gamma = math.atan2(self.y, self.x)
        else:
            gamma = math.atan2(-self.y, -self.x)
        return (alpha*180/math.pi, beta*180/math.pi, gamma*180/math.pi)

if __name__ == "__main__":
    leg = SpiderLeg()
    print(leg.get_servo_angles())
    leg.z = 0
    print(leg.get_servo_angles())
