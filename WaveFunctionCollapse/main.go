package main

import (
	"math/rand"

	"seehuhn.de/go/ncurses"
)

type Tile struct {
	icon    string // UP DOWN LEFT RIGHT EMPTY
	race    string // "up" "down" "left" "right" "empty" "null"
	entropy uint8
	posb    []string // UP DOWN LEFT RIGHT EMPTY
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
				win.Print(val.icon)
				// win.Print(val.entropy)
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
		switch grid[y][x].posb[icon_ind] {
		case UP:
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
		case RIGHT:
			icon = RIGHT
			name = "right"
			if y > 0 {
				removeElems([]string{UP, EMPTY}, &grid[y-1][x])
			}
			if y < H-1 {
				removeElems([]string{EMPTY, DOWN}, &grid[y+1][x])
			}
			if x > 0 {
				removeElems([]string{RIGHT, UP, DOWN}, &grid[y][x-1])
			}
			if x < W-1 {
				removeElems([]string{EMPTY, RIGHT}, &grid[y][x+1])
			}
		case DOWN:
			icon = DOWN
			name = "down"
			if y > 0 {
				removeElems([]string{RIGHT, LEFT, DOWN}, &grid[y-1][x])
			}
			if y < H-1 {
				removeElems([]string{DOWN, EMPTY}, &grid[y+1][x])
			}
			if x > 0 {
				removeElems([]string{LEFT, EMPTY}, &grid[y][x-1])
			}
			if x < W-1 {
				removeElems([]string{EMPTY, RIGHT}, &grid[y][x+1])
			}

		case LEFT:
			icon = LEFT
			name = "left"

			if y > 0 {
				removeElems([]string{UP, EMPTY}, &grid[y-1][x])
			}
			if y < H-1 {
				removeElems([]string{DOWN, EMPTY}, &grid[y+1][x])
			}
			if x > 0 {
				removeElems([]string{LEFT, EMPTY}, &grid[y][x-1])
			}
			if x < W-1 {
				removeElems([]string{LEFT, UP, DOWN}, &grid[y][x+1])
			}
		case EMPTY:
			icon = EMPTY
			name = "empty"
			if y > 0 {
				removeElems([]string{DOWN, LEFT, RIGHT}, &grid[y-1][x])
			}
			if y < H-1 {
				removeElems([]string{UP, RIGHT, LEFT}, &grid[y+1][x])
			}
			if x > 0 {
				removeElems([]string{RIGHT, UP, DOWN}, &grid[y][x-1])
			}
			if x < W-1 {
				removeElems([]string{LEFT, DOWN, UP}, &grid[y][x+1])
			}

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
	_, width := win.GetMaxYX()
	win.Move(0, (width/2)-20)
	win.Println("Welcome to Wave Function Collapse World")

	//arr := []string{UP, RIGHT, DOWN, LEFT, EMPTY}
	//for _, v := range arr {
	//	win.Println(" " + v)
	//	win.Println(v + RIGHT + v)
	//	win.Println(" " + v)
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
		win.Println(grid[y][x])
	}

	win.Readline(0)

	ncurses.EndWin()
}
