package main

import (
	"fmt"
	"math/rand"
)

func generateMap(state *state) {
	wave := createWave()
	state.wave = &wave
	finished := false

	cellX, cellY := getRandomCell(wave)
	collapseCell(&wave[cellY][cellX])
	propagate(&wave, cellX, cellY)

	var coordinates [][2]int
	for !finished {

		coordinates, finished = getLowestEntropy(wave)
		fmt.Println(len(coordinates))
		if len(coordinates) == 0 {
			cellX, cellY = getRandomCell(wave)
			collapseCell(&wave[cellY][cellX])
			propagate(&wave, cellX, cellY)
			continue

		}

		coordToCollapse := coordinates[rand.Intn(len(coordinates))]
		cellToCollapse := &wave[coordToCollapse[1]][coordToCollapse[0]]
		collapseCell(cellToCollapse)
	}
	fmt.Println("done")
}

func propagate(wave *[MAP_H][MAP_W][TOTAL_TILES]bool, x, y int) {
	// Bounds Check
	// collapsed Check
	if x > MAP_W || x < 0 {
		return
	}
	if y > MAP_H || y < 0 {
		return
	}
	if getEntropy(wave[y][x]) == 1 {
		return
	}
	if y+1 < MAP_H {
		wave[y+1][x] = boolAnd((*wave)[y+1][x], (*wave)[y][x])
	}
}

func boolAnd(a, b [TOTAL_TILES]bool) (c [TOTAL_TILES]bool) {
	for i := range a {
		if a[i] == b[i] {
			c[i] = a[i]
		}
	}
	return c
}

func collapseCell(cell *[TOTAL_TILES]bool) {
	if cell == nil {
		return
	}

	for valid := true; valid; {
		idx := rand.Intn(TOTAL_TILES)
		if cell[idx] {
			// default of bool is false; now cell is an array of false
			*cell = [TOTAL_TILES]bool{}
			cell[idx] = true
			return
			valid = false
		}
	}
}

// Get a random cell which has not been collapsed
func getRandomCell(wave [MAP_H][MAP_W][TOTAL_TILES]bool) (int, int) {
	valid := true
	for valid {
		x, y := rand.Intn(MAP_W), rand.Intn(MAP_H)
		if getEntropy(wave[y][x]) == 1 {
			continue
		}
		return x, y
	}
	return -1, -1
}

func getLowestEntropy(wave [MAP_H][MAP_W][TOTAL_TILES]bool) ([][2]int, bool) {
	var lowestEntropy uint8
	var lowestEntropyY int
	var lowestEntropyX int

	var coordinates [][2]int
	var cell_Entropy uint8
	finished := true

	lowestEntropy = TOTAL_TILES

	for y, row := range wave {
	traverseRow:
		for x, cell := range row {
			cell_Entropy = getEntropy(cell)
			if cell_Entropy == 1 {
				continue traverseRow
			}
			finished = false

			if cell_Entropy == lowestEntropy {
				coordinates = append(coordinates, [2]int{x, y})
				continue traverseRow
			}

			if cell_Entropy < lowestEntropy {
				coordinates = [][2]int{}
				lowestEntropy = cell_Entropy
				lowestEntropyY = y
				lowestEntropyX = x
			}
		}
	}
	fmt.Println(finished, "IGNORE THIS MESSAGE", lowestEntropyX, lowestEntropyY)
	return coordinates, finished
}

func getEntropy(cell [TOTAL_TILES]bool) (entropy uint8) {
	for _, accepted := range cell {
		if accepted {
			entropy += 1
		}
	}
	return entropy
}

func createWave() (wave [MAP_H][MAP_W][TOTAL_TILES]bool) {
	var model [TOTAL_TILES]bool
	for i := range TOTAL_TILES {
		model[i] = true
	}

	for i, row := range wave {
		for j := range row {
			wave[i][j] = model
		}
	}
	return wave
}
