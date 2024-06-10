package main

import (
	"fmt"
)

const (
	W = 20
	H = 20
)

var DIR [][2]int

func genGrid() [H][W]bool {
	var grid [H][W]bool
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// if rand.Intn(100) < 25 {
			// 	grid[i][j] = true
			// 	continue
			// }
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
	fmt.Println("                                      -------------------")
	for _, col := range grid {
		fmt.Print("                                      ")
		for _, val := range col {
			if val {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
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
	grid := genGrid()
	grid[10][10] = true
	grid[9][10] = true
	grid[10][9] = true
	grid[11][10] = true
	grid[11][11] = true
	for i := 0; i < 19; i++ {
		drawGrid(grid)
		evolve(&grid)
		drawGrid(grid)
	}
}
