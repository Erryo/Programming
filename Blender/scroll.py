#!/usr/bin/env python3

import time
import bpy


wait = 0
rainbow = {
    0: (255, 0, 0),
    1: (255, 127, 0),
    2: (255, 255, 0),
    3: (127, 255, 0),
    4: (0, 255, 0),
    5: (0, 255, 127),
    6: (0, 255, 255),
    7: (0, 127, 255),
    8: (0, 0, 255),
    9: (127, 0, 255),
    10: (255, 0, 255),
    11: (255, 0, 127),
}


def rainbowChange(location):
    if location >= len(rainbow):
        location = location - len(rainbow)
    return rainbow.get(location)


def mid_eff(pipe,frame,lp, offset=0, diff=0):
    px_set = 0
    
    while px_set < lp.p_size - 1  - diff:
        px = 0
        if px_set == 0:
            px = int(offset / 4)
        while px < 3:
            lp.set_pixel_y(
                pipe,
                x=px_set,
                color=rainbowChange(px_set + offset),
                y=px,
            )
            px += 1
            px_set += 1
            lp.show(frame)
            frame += 1
    return frame

def effect(lp):
    global rainbow

    lp.set_pixel_0(0, x=0, color=rainbow.get(0))
    lp.set_pixel_1(0, x=0, color=rainbow.get(4))
    lp.set_pixel_2(0, x=0, color=rainbow.get(8))
    lp.show(1)


    pipe = 0
    frame = 2
    while pipe < lp.pipes:
        frame = mid_eff(pipe, frame, lp)
        frame = mid_eff(pipe, frame, lp, 4, 1)
        frame = mid_eff(pipe, frame, lp, 8)

        pipe += 1


if __name__ == "__main__":
    LP  = bpy.data.texts["middleman.py"].as_module()
    
    lp = LP.Lightpipe()
    effect(lp)
