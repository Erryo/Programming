#!/usr/bin/env python3

import time


wait = 0.25
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


def initPipeColors(lp):
    pipesColors = []
    for pipe in range(lp.pipes):
        pipeCol = []
        for px_set_N in range(lp.p_size):
            px_set = []
            px_set.append((0, 0, 0))
            px_set.append((0, 0, 0))
            px_set.append((0, 0, 0))
            pipeCol.append(px_set)
        pipesColors.append(pipeCol)
    print(pipesColors)
    return pipesColors


def initPipes(lp, pb):
    pipesColors = initPipeColors(lp)
    print(pipesColors)
    for pipe in range(lp.pipes):
        for px_set in range(lp.p_size):
            pipesColors[pipe][px_set][0] = pb.random_color()
            pipesColors[pipe][px_set][1] = pb.random_color()
            pipesColors[pipe][px_set][2] = pb.random_color()
            lp.set_pixel_0(pipe, x=px_set, color=pipesColors[pipe][px_set][0])
            lp.set_pixel_1(pipe, x=px_set, color=pipesColors[pipe][px_set][1])
            lp.set_pixel_2(pipe, x=px_set, color=pipesColors[pipe][px_set][2])
    return pipesColors


def rainbowChange(location):
    if location >= len(rainbow):
        location = location - len(rainbow)
    return rainbow.get(location)


def changeColor(px, location):
    new_col = px[location] + 40
    px_empty = [px[0], px[1], px[2]]
    px_empty[location] = new_col
    if new_col > 255 or new_col < 0:
        px_empty[location] -= 255
    return (px_empty[0], px_empty[1], px_empty[2])


def updateColors(pipesColors, lp):
    for pipe in range(lp.pipes):
        for px_set in range(lp.p_size):
            lp.set_pixel_0(pipe, x=px_set, color=pipesColors[pipe][px_set][0])
            lp.set_pixel_1(pipe, x=px_set, color=pipesColors[pipe][px_set][1])
            lp.set_pixel_2(pipe, x=px_set, color=pipesColors[pipe][px_set][2])
    lp.show()


def mid_eff(pipe, pipesColors, lp, offset=0, diff=0):
    px_set = 0
    while px_set < lp.p_size - 1 - diff:
        px = 0
        if px_set == 0:
            px = int(offset / 4)
        while px < 3:
            print(px, px_set, lp.p_size)
            pipesColors[pipe][px_set][px] = rainbowChange(px_set + offset)

            px += 1
            px_set += 1

            time.sleep(wait)
            updateColors(pipesColors, lp)


def effect(lp):
    global rainbow

    #       0 1 2 3
    #  0  | a c b a                         |
    #  1  | b a c b                         |
    #  2  | c b a c                         |
    #      _ _ __ _ _ _ _ __ _ _ __ _ _ _ _

    while True:
        pipesColors = initPipeColors(lp)

        pipesColors[0][0][0] = rainbow.get(0)
        pipesColors[0][0][1] = rainbow.get(4)
        pipesColors[0][0][2] = rainbow.get(8)

        lp.show()

        pipe = 0
        while pipe < lp.pipes:
            mid_eff(pipe, pipesColors, lp)
            mid_eff(pipe, pipesColors, lp, 4, 1)
            mid_eff(pipe, pipesColors, lp, 8)

            pipe += 1


class LightPipe:
    pipes = 1
    p_size = 16


if __name__ == "__main__":
    lp = LightPipe()
    effect(lp)
