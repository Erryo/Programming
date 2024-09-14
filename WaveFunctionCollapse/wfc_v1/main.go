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
	Name   string
	Icon   string
	North  []string // Not allowed
	East   []string
	South  []string
	West   []string
	Pixels [8]bool
}

const (
	W          int    = 60
	H          int    = 60
	UP         string = "┴"
	RIGHT      string = "├"
	DOWN       string = "┬"
	LEFT       string = "┤"
	EMPTY      string = "▒"
	DOWN_RIGHT string = "┌"
	DOWN_LEFT  string = "┐"
	UP_RIGHT   string = "└"
	UP_LEFT    string = "┘"
	NULL       string = " "
)

var win ncurses.Window = *ncurses.Init()

var baseUp BaseTile = BaseTile{
	Name:   "UP",
	Icon:   UP,
	North:  []string{UP, EMPTY, UP_LEFT, UP_LEFT},
	East:   []string{RIGHT, EMPTY, UP_RIGHT, DOWN_RIGHT},
	South:  []string{UP, LEFT, RIGHT, UP_LEFT, UP_LEFT},
	West:   []string{LEFT, EMPTY, UP_LEFT, DOWN_LEFT},
	Pixels: [8]bool{false, true, false, true, false, false, false, true},
}

var baseRight BaseTile = BaseTile{
	Name:   "RIGHT",
	Icon:   RIGHT,
	North:  []string{UP, EMPTY, UP_RIGHT, UP_LEFT},
	East:   []string{RIGHT, EMPTY, UP_RIGHT, DOWN_RIGHT},
	South:  []string{DOWN, EMPTY, DOWN_RIGHT, DOWN_LEFT},
	West:   []string{UP, DOWN, RIGHT, DOWN, UP_RIGHT, DOWN_RIGHT},
	Pixels: [8]bool{false, true, false, true, false, true, false, false},
}

var baseDown BaseTile = BaseTile{
	Name:   "DOWN",
	Icon:   DOWN,
	North:  []string{RIGHT, LEFT, DOWN, DOWN_LEFT, DOWN_RIGHT},
	East:   []string{EMPTY, RIGHT, UP_RIGHT, DOWN_RIGHT},
	South:  []string{DOWN, EMPTY, DOWN_RIGHT, DOWN_LEFT},
	West:   []string{LEFT, EMPTY, UP_LEFT, DOWN_LEFT},
	Pixels: [8]bool{false, false, false, true, false, true, false, true},
}

var baseLeft BaseTile = BaseTile{
	Name:   "LEFT",
	Icon:   LEFT,
	North:  []string{UP, EMPTY, UP_RIGHT, UP_LEFT},
	East:   []string{LEFT, UP, DOWN, UP_LEFT, DOWN_LEFT},
	South:  []string{DOWN, EMPTY, DOWN_RIGHT, DOWN_LEFT},
	West:   []string{LEFT, EMPTY, UP_LEFT, DOWN_LEFT},
	Pixels: [8]bool{false, true, false, false, false, true, false, true},
}

var baseEmpty BaseTile = BaseTile{
	Name:   "EMPTY",
	Icon:   EMPTY,
	North:  []string{LEFT, RIGHT, DOWN, DOWN_RIGHT, DOWN_LEFT},
	East:   []string{UP, LEFT, DOWN, UP_LEFT, DOWN_LEFT},
	South:  []string{UP, RIGHT, LEFT, UP_RIGHT, UP_LEFT},
	West:   []string{DOWN, UP, RIGHT, UP_RIGHT, DOWN_RIGHT},
	Pixels: [8]bool{false, false, false, false, false, false, false, false},
}

var baseUpRight BaseTile = BaseTile{
	Name:   "UP_RIGHT",
	Icon:   UP_RIGHT,
	North:  []string{UP, EMPTY, UP_RIGHT, UP_LEFT},
	East:   []string{RIGHT, EMPTY, UP_RIGHT, DOWN_RIGHT},
	South:  []string{UP, RIGHT, LEFT, UP_RIGHT, UP_LEFT},
	West:   []string{UP, RIGHT, DOWN, UP_RIGHT, DOWN_RIGHT},
	Pixels: [8]bool{false, true, false, true, false, false, false, false},
}

var baseUpLeft BaseTile = BaseTile{
	Name:   "UP_LEFT",
	Icon:   UP_LEFT,
	North:  []string{UP, EMPTY, UP_RIGHT, UP_LEFT},
	East:   []string{UP, DOWN, LEFT, UP_LEFT, DOWN_LEFT},
	South:  []string{UP, RIGHT, LEFT, UP_RIGHT, UP_LEFT},
	West:   []string{LEFT, EMPTY, UP_LEFT, DOWN_LEFT},
	Pixels: [8]bool{false, true, false, false, false, false, false, true},
}

var baseDownRight BaseTile = BaseTile{
	Name:   "DOWN_RIGHT",
	Icon:   DOWN_RIGHT,
	North:  []string{RIGHT, DOWN, LEFT, DOWN_RIGHT, DOWN_LEFT},
	East:   []string{RIGHT, EMPTY, UP_RIGHT, DOWN_RIGHT},
	South:  []string{DOWN, EMPTY, DOWN_RIGHT, DOWN_LEFT},
	West:   []string{UP, RIGHT, DOWN, UP_RIGHT, DOWN_RIGHT},
	Pixels: [8]bool{false, false, false, true, false, true, false, false},
}

var baseDownLeft BaseTile = BaseTile{
	Name:   "DOWN_LEFT",
	Icon:   DOWN_LEFT,
	North:  []string{RIGHT, DOWN, LEFT, DOWN_RIGHT, DOWN_LEFT},
	East:   []string{UP, DOWN, LEFT, UP_LEFT, DOWN_LEFT},
	South:  []string{DOWN, EMPTY, DOWN_RIGHT, DOWN_LEFT},
	West:   []string{LEFT, EMPTY, UP_LEFT, DOWN_LEFT},
	Pixels: [8]bool{false, false, false, false, false, true, false, true},
}

func removeElems(elems []string, tile *Tile) bool {
	array := tile.posb[:0]
	var deletedAnything bool
	var shouldDelete bool
	for _, v1 := range tile.posb {
		shouldDelete = false
		for _, v2 := range elems {
			if v1.Icon == v2 {
				shouldDelete = true
				deletedAnything = true
				tile.entropy -= 1
				break
			}
		}
		if !shouldDelete {
			array = append(array, v1)
		}
	}
	tile.posb = array
	win.Move(0, 0)
	win.Println(deletedAnything)
	return deletedAnything
}

func genGrid() [H][W]Tile {
	var grid [H][W]Tile
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			grid[i][j] = Tile{
				icon:    NULL,
				race:    "null",
				entropy: 9,
				posb:    []BaseTile{baseUp, baseRight, baseDown, baseLeft, baseEmpty, baseUpRight, baseUpLeft, baseDownLeft, baseDownRight},
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
			} else {
				win.Print(val.icon)
			}

			win.Refresh()
		}
		win.Println("")
		win.Move((height/2)+i+1, width/2)
	}
}

func collapseTile(grid *[H][W]Tile, x, y int) {
	if x > W || x < 0 {
		return
	}

	if y > H || y < 0 {
		return
	}
	if grid[y][x].race != "null" {
		return
	}

	icon_ind := rand.Intn(int(grid[y][x].entropy))
	toAddtile := grid[y][x].posb[icon_ind]
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
		icon:    toAddtile.Icon,
		race:    toAddtile.Name,
		entropy: 0,
		posb:    []BaseTile{},
	}
	grid[y][x] = tile
}

func findLowEntropy(grid [H][W]Tile) (bool, int, int) {
	var done bool = true
	var entropy uint8 = 9
	y, x := 0, 0

	for i, row := range grid {
		for j, tile := range row {
			if tile.race == "null" {
				done = false
			}
			if tile.entropy == 0 && tile.race == "null" {
				win.Readline(0)
				return true, y, x
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

		win.Println(time.Since(start))
	}
	win.Readline(0)
	ncurses.EndWin()
	fmt.Println(time.Since(start))
}
