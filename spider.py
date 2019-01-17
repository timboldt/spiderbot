#! /usr/bin/env python

import math
import sys
import time

import ssc32
import spider_leg

ssc = ssc32.SSC32('/dev/ttyS6', 9600, count=32)

for i in range(0, 32):
    ssc[i].degrees = 0
ssc[13].degrees = 45
ssc[29].degrees = -45
ssc[18].degrees = 45

# Back Right
br_tibia = ssc[0]
br_femur = ssc[1]
br_coxa = ssc[2]
br_leg = spider_leg.SpiderLeg()
br_leg.x = 70
br_leg.y = 20
br_leg.z = -40
a, b, g = br_leg.get_servo_angles()
print(a,b,g)
br_tibia.degrees = 90 - a
br_femur.degrees = 180 - b 
br_coxa.degrees = g - 90 
#backwards = True

ssc.commit(1000)
while not ssc.is_done():
    time.sleep(0.1)
sys.exit(0)

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

br_leg = spider_leg.SpiderLeg()

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
