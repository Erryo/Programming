package main

import (
	"fmt"
	"math/rand"
	"time"

	"seehuhn.de/go/ncurses"
)

type Tile struct {
	icon    string // UP DOWN LEFT RIGHT EMPTY
	race    string // "up" "down" "left" "right" "empty" "null"
	entropy uint8
	posb    []BaseTile // UP DOWN LEFT RIGHT EMPTY
}

type BaseTile struct {
	Name  string
	Icon  string
	North []string // Not allowed
	East  []string
	South []string
	West  []string
}

const (
	W     int    = 40
	H     int    = 20
	UP    string = "┴"
	RIGHT string = "├"
	DOWN  string = "┬"
	LEFT  string = "┤"
	EMPTY string = "▒"
	NULL  string = " "
)

var win ncurses.Window = *ncurses.Init()

var baseUp BaseTile = BaseTile{
	Name:  "up",
	Icon:  UP,
	North: []string{UP, EMPTY},
	East:  []string{RIGHT, EMPTY},
	South: []string{UP, LEFT, RIGHT},
	West:  []string{LEFT, EMPTY},
}

var baseRight BaseTile = BaseTile{
	Name:  "right",
	Icon:  RIGHT,
	North: []string{UP, EMPTY},
	East:  []string{RIGHT, EMPTY},
	South: []string{DOWN, EMPTY},
	West:  []string{UP, DOWN, RIGHT},
}

var baseDown BaseTile = BaseTile{
	Name:  "down",
	Icon:  DOWN,
	North: []string{RIGHT, LEFT, DOWN},
	East:  []string{EMPTY, RIGHT},
	South: []string{DOWN, EMPTY},
	West:  []string{LEFT, EMPTY},
}

var baseLeft BaseTile = BaseTile{
	Name:  "left",
	Icon:  LEFT,
	North: []string{UP, EMPTY},
	East:  []string{LEFT, UP, DOWN},
	South: []string{DOWN, EMPTY},
	West:  []string{LEFT, EMPTY},
}

var baseEmpty BaseTile = BaseTile{
	Name:  "empty",
	Icon:  EMPTY,
	North: []string{LEFT, RIGHT, DOWN},
	East:  []string{UP, LEFT, DOWN},
	South: []string{UP, RIGHT, LEFT},
	West:  []string{DOWN, UP, RIGHT},
}

func removeElems(elems []string, tile *Tile) {
	array := tile.posb[:0]
	var shouldDelete bool
	for _, v1 := range tile.posb {
		shouldDelete = false
		for _, v2 := range elems {
			if v1.Icon == v2 {
				shouldDelete = true
				tile.entropy -= 1
				break
			}
		}
		if !shouldDelete {
			array = append(array, v1)
		}
	}
	tile.posb = array
}

func genGrid() [H][W]Tile {
	var grid [H][W]Tile
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			grid[i][j] = Tile{
				icon:    NULL,
				race:    "null",
				entropy: 5,
				posb:    []BaseTile{baseUp, baseRight, baseDown, baseLeft, baseEmpty},
			}
		}
	}
	return grid
}

func drawGrid(grid [H][W]Tile) {
	height, width := win.GetMaxYX()
	height -= H
	width -= W
	win.Move(height/2, width/2)
	for i, row := range grid {
		for _, val := range row {
			if val.race != "null" {
				win.Print(val.icon)
				//			fmt.Print(val.icon)
			} else {
				// win.Print(val.icon)
				win.Print(val.entropy)
				//			fmt.Print(val.entropy)
			}

			win.Refresh()
			// time.Sleep(time.Second / 50)
		}
		win.Println("")
		// fmt.Println("")
		win.Move((height/2)+i+1, width/2)
	}
}

func collapseTile(grid *[H][W]Tile, x, y int) {
	for {

		if grid[y][x].race != "null" {
			continue
		}
		var name string
		var icon string

		icon_ind := rand.Intn(int(grid[y][x].entropy))
		toAddtile := grid[y][x].posb[icon_ind]
		icon = toAddtile.Icon
		name = toAddtile.Name
		if y > 0 {
			removeElems(toAddtile.North, &grid[y-1][x])
		}
		if y < H-1 {
			removeElems(toAddtile.South, &grid[y+1][x])
		}
		if x > 0 {
			removeElems(toAddtile.West, &grid[y][x-1])
		}
		if x < W-1 {
			removeElems(toAddtile.East, &grid[y][x+1])
		}

		var tile Tile = Tile{
			icon:    icon,
			race:    name,
			entropy: 0,
			posb:    []BaseTile{},
		}
		grid[y][x] = tile
		break
	}
}

func findLowEntropy(grid [H][W]Tile) (bool, int, int) {
	var done bool = true
	var entropy uint8 = 5
	y, x := 0, 0

	for i, row := range grid {
		for j, tile := range row {
			if tile.race == "null" {
				done = false
			}
			if tile.entropy == 0 {
				continue
			}
			if tile.entropy < entropy {
				entropy = tile.entropy
				y = i
				x = j
			}
		}
	}
	return done, y, x
}

func main() {
	start := time.Now()
	_, width := win.GetMaxYX()
	win.Move(0, (width/2)-20)
	win.Println("Welcome to Wave Function Collapse World")

	//arr := []string{UP, RIGHT, DOWN, LEFT, EMPTY}
	//for _, t := range arr {
	//	fmt.Println("----------------")
	//	for _, v := range arr {
	//		fmt.Println(" " + v)
	//		fmt.Println(v + t + v)
	//		fmt.Println(" " + v)
	//		fmt.Println("-___-___-___-")
	//	}
	//}

	var done bool
	grid := genGrid()
	x := rand.Intn(W)
	y := rand.Intn(H)

	for i := 1; i > 0; i++ {

		collapseTile(&grid, x, y)
		drawGrid(grid)
		done, y, x = findLowEntropy(grid)
		if done {
			break
		}
		win.Move(0, (width/2)-20)
		win.Println("X: ", x, "Y:", y)
		win.Println(grid[y][x].icon)
	}
	ncurses.EndWin()
	fmt.Println(time.Since(start))
}
