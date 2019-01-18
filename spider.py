#! /usr/bin/env python3

import math
import sys
import time

import ssc32
import spider_body

ssc = ssc32.SSC32('/dev/ttyS6', 9600, count=32)
body = spider_body.SpiderBody(ssc)
body.stand_up()
for i in range(4):
    body.shift_weight_off_leg(i)
    body.lift_leg(i)
    body.stand_up()
body.sit_down()