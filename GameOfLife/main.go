package main

import (
	"math/rand"
	"time"

	"seehuhn.de/go/ncurses"
)

const (
	W = 20
	H = 20
)

var (
	DIR [][2]int
	win *ncurses.Window
)

func genGrid() [H][W]bool {
	var grid [H][W]bool
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if rand.Intn(100) < 25 {
				grid[i][j] = true
				continue
			}
			grid[i][j] = false
		}
	}
	// (-1,-1)(-1,0)(-1,+1)
	// (0,-1)(0,0)(0,+1)
	// (+1,-1)(+1,0)(+1,+1)
	DIR = append(DIR, [2]int{-1, -1})
	DIR = append(DIR, [2]int{-1, 0})
	DIR = append(DIR, [2]int{-1, 1})
	DIR = append(DIR, [2]int{0, -1})
	DIR = append(DIR, [2]int{0, 1})
	DIR = append(DIR, [2]int{1, -1})
	DIR = append(DIR, [2]int{1, 0})
	DIR = append(DIR, [2]int{1, +1})
	return grid
}

func drawGrid(grid [H][W]bool) {
	// fmt.Println("                                      -------------------")
	win.Move(0, 0)
	win.Println("Conway's Game of Life")
	for _, col := range grid {
		// fmt.Print("                                      ")
		win.Print("                                      ")
		for _, val := range col {
			if val {
				// fmt.Print("■")
				win.Print("■")
			} else {
				// fmt.Print("□")
				win.Print("□")
			}
			// fmt.Print(" ")
			win.Print(" ")
		}
		// fmt.Println("")
		win.Println("")
	}
	win.Refresh()
}

func evolve(grid *[H][W]bool) {
	var oldG [H][W]bool = *grid
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var alive uint8 = 0
			// (-1,-1)(-1,0)(-1,+1)
			// (+0,-1)------(0,+1)
			// (+1,-1)(+1,0)(+1,+1)
			for ind, v := range DIR {
				if i == 0 && (ind == 0 || ind == 1 || ind == 2) {
					continue
				}
				if j == 0 && (ind == 0 || ind == 3 || ind == 5) {
					continue
				}
				if i == H-1 && (ind == 5 || ind == 6 || ind == 7) {
					continue
				}
				if j == W-1 && (ind == 2 || ind == 4 || ind == 7) {
					continue
				}
				if oldG[i+v[0]][j+v[1]] {
					alive++
				}
			}
			// Top->Right->Down->Left

			if grid[i][j] {
				if alive > 3 || alive < 2 {
					grid[i][j] = false
				}
			} else {
				if alive == 3 {
					grid[i][j] = true
				}
			}

		}
	}
}

func main() {
	win = ncurses.Init()
	grid := genGrid()
	for i := 0; i < 100; i++ {

		time.Sleep(time.Second / 4)
		drawGrid(grid)
		evolve(&grid)
		drawGrid(grid)
	}
	ncurses.EndWin()
}
