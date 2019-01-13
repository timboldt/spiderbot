#! /usr/bin/env python

import math
import time

import ssc32
import spider_leg

ssc = ssc32.SSC32('/dev/ttyS3', 115200, count=32)

# Back Right
br_tibia = ssc[0]
br_femur = ssc[1]
br_coxa = ssc[2]
backwards = True

# # Front Right
# br_tibia = ssc[15]
# br_femur = ssc[14]
# br_coxa = ssc[13]
# backwards=False

# # Front Left
# br_tibia = ssc[31]
# br_femur = ssc[30]
# br_coxa = ssc[29]
# backwards=True

# Back Left
# br_tibia = ssc[16]
# br_femur = ssc[17]
# br_coxa = ssc[18]
# backwards=False

br_leg = spider_leg.SpiderLeg(backwards=backwards)

br_leg.x = 70
br_leg.y = 20
#br_leg.z = 27
a, b, g = br_leg.get_servo_angles()
print(a, b, g)
(br_tibia.degrees, br_femur.degrees, br_coxa.degrees) = (a, b, g)
ssc.commit(time=1000)
while not ssc.is_done():
    time.sleep(0.1)

br_leg.x = 20
br_leg.y = 70
#br_leg.z = -50
a, b, g = br_leg.get_servo_angles()
print(a, b, g)
(br_tibia.degrees, br_femur.degrees, br_coxa.degrees) = (a, b, g)
ssc.commit(time=1000)
while not ssc.is_done():
    time.sleep(0.1)

br_leg.x = 70
br_leg.y = 70
#br_leg.z = -50
a, b, g = br_leg.get_servo_angles()
print(a, b, g)
(br_tibia.degrees, br_femur.degrees, br_coxa.degrees) = (a, b, g)
ssc.commit(time=1000)
while not ssc.is_done():
    time.sleep(0.1)

br_leg.x = 44
br_leg.y = 44
#br_leg.z = -50
a, b, g = br_leg.get_servo_angles()
print(a, b, g)
(br_tibia.degrees, br_femur.degrees, br_coxa.degrees) = (a, b, g)
ssc.commit(time=1000)
while not ssc.is_done():
    time.sleep(0.1)
