package main

import (
	"fmt"
)

func getSide(v, v1 *BaseTile, i, j, m, n int) bool {
	var side bool
	for k := i; k < j; {
		for l := n - 1; l >= m; {
			k_ := k
			l_ := l
			if k >= len(v.Pixels) {
				k_ = 0
			}
			if l >= len(v1.Pixels) {
				l_ = 0
				// fmt.Println("l_ 000")
			}
			// fmt.Println(v.Pixels[k_], k, v1.Pixels[l_], l)
			// time.Sleep(time.Second / 4)
			if v.Pixels[k_] != v1.Pixels[l_] {
				side = false
				break
			}
			side = true
			k++
			l--
		}
		break
		fmt.Println("")
	}
	return side
}

func testTile() {
	arr := []*BaseTile{&baseUp, &baseRight, &baseDown, &baseLeft, &baseEmpty, &baseUpRight, &baseUpLeft, &baseDownRight, &baseDownLeft}
	// fmt.Println(baseDown.North, baseDown.East, baseDown.South, baseDown.West)
	baseDown.North = []string{}
	baseDown.East = []string{}
	baseDown.South = []string{}
	baseDown.West = []string{}
	for _, v := range arr {
		for _, v1 := range arr {
			// fmt.Println(v.Name, v1.Name)
			// fmt.Println(v.North, v.East, v.South, v.West)
			var upSide bool = getSide(v, v1, 0, 3, 4, 7)    // Top with Bottom
			var rightSide bool = getSide(v, v1, 2, 5, 6, 9) // Right Left
			var downSide bool = getSide(v, v1, 4, 7, 0, 3)  // Bot Top
			var leftSide bool = getSide(v, v1, 6, 9, 2, 5)  //
			if !upSide {
				v.North = append(v.North, v1.Icon)
			}
			if !rightSide {
				v.East = append(v.East, v1.Icon)
			}
			if !downSide {
				v.South = append(v.South, v1.Icon)
			}
			if !leftSide {
				v.West = append(v.West, v1.Icon)
			}
			//	fmt.Println(v.North, v.East, v.South, v.West)
			//	fmt.Println("")
		}
	}
	// fmt.Println(baseEmpty.North, baseDown.East, baseDown.South, baseDown.West)
}
