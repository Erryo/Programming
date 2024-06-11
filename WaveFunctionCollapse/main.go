package main

import (
	"math/rand"
	"time"

	"seehuhn.de/go/ncurses"
)

type Tile struct {
	icon    string // UP DOWN LEFT RIGHT EMPTY
	race    string // "up" "down" "left" "right" "empty" "null"
	entropy uint8
	posb    []string // UP DOWN LEFT RIGHT EMPTY
}

const (
	W     int    = 10
	H     int    = 10
	UP    string = "┴"
	RIGHT string = "├"
	DOWN  string = "┬"
	LEFT  string = "┤"
	EMPTY string = "▒"
	NULL  string = "□"
)

var win ncurses.Window = *ncurses.Init()

func removeElems(elems []string, tile *Tile) {
	array := tile.posb[:0]
	var shouldDelete bool
	for _, v1 := range tile.posb {
		shouldDelete = false
		for _, v2 := range elems {
			if v1 == v2 {
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
				posb:    []string{UP, RIGHT, DOWN, LEFT, EMPTY},
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
				win.Print(val.entropy)
				//			fmt.Print(val.entropy)
			}

			win.Refresh()
			time.Sleep(time.Second / 50)
		}
		win.Println("")
		// fmt.Println("")
		win.Move((height/2)+i+1, width/2)
	}
}

func firstGen(grid *[H][W]Tile) {
	for {
		x := rand.Intn(W)
		y := rand.Intn(H)
		if grid[y][x].race != "null" {
			continue
		}
		var name string
		var icon string

		icon_ind := 0
		switch icon_ind {
		case 0:
			icon = UP
			name = "up"
			if y > 0 {
				removeElems([]string{UP, EMPTY}, &grid[y-1][x])
			}
			if y < H-1 {
				removeElems([]string{UP, RIGHT, LEFT}, &grid[y+1][x])
			}
			if x > 0 {
				removeElems([]string{LEFT, EMPTY}, &grid[y][x-1])
			}
			if x < W-1 {
				removeElems([]string{EMPTY, RIGHT}, &grid[y][x+1])
			}
		case 1:
			icon = RIGHT
			name = "right"
		case 2:
			icon = DOWN
			name = "down"

		case 3:
			icon = LEFT
			name = "left"

		case 4:
			icon = EMPTY
			name = "empty"

		}
		var tile Tile = Tile{
			icon:    icon,
			race:    name,
			entropy: 0,
			posb:    []string{},
		}
		grid[y][x] = tile

		break
	}
}

func main() {
	defer ncurses.EndWin()
	_, width := win.GetMaxYX()
	win.Move(0, (width/2)-20)
	win.Print("Welcome to Wave Function Collapse World")

	grid := genGrid()
	firstGen(&grid)
	drawGrid(grid)
	win.Readline(0)
}
