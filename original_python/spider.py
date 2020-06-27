#! /usr/bin/env python3

import math
import sys
import time

import ssc32
import spider_body

ssc = ssc32.SSC32('/dev/rfcomm0', 9600, count=32)
body = spider_body.SpiderBody(ssc)
# for i in range(4):
#     body.stand_up()
#     time.sleep(1)
#     body.sit_down()
#     time.sleep(1)
for i in range(4):
    body.shift_weight_off_leg(i)
    body.leg_up(i, 40)
    body.leg_down(i)
    body.stand_up()
#body.walk()
#time.sleep(3)
body.sit_down()