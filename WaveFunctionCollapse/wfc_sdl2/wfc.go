package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func generateMap(state *state) {
	start := time.Now()
	wave := createWave()
	models := getTileModel()
	state.wave = &wave
	finished := false

	cellX, cellY := getRandomCell(wave)
	state.selected = [2]int{cellX, cellY}
	collapseCell(&wave[cellY][cellX])
	propagate(&wave, cellX, cellY, models, state)

	var coordinates [][2]int
	for {

		coordinates, finished = getLowestEntropy(wave)
		if finished {
			break
		}
		if len(coordinates) == 0 {
			cellX, cellY = getRandomCell(wave)
			collapseCell(&wave[cellY][cellX])
			state.selected = [2]int{cellX, cellY}
			propagate(&wave, cellX, cellY, models, state)
			continue
		}

		coordToCollapse := coordinates[rand.Intn(len(coordinates))]
		cellToCollapse := &wave[coordToCollapse[1]][coordToCollapse[0]]
		collapseCell(cellToCollapse)
		state.selected = [2]int{cellX, cellY}
		propagate(&wave, coordToCollapse[0], coordToCollapse[1], models, state)
	}
	fmt.Print(time.Since(start))
	fmt.Println("done")
}

func propagate(wave *[MAP_H][MAP_W][TOTAL_TILES]bool, x, y int, models []TileModel, state *state) {
	sdl.Delay(30)
	// Bounds Check
	// collapsed Check
	// fmt.Println("started", x, y)
	state.selected = [2]int{x, y}
	if x > MAP_W || x < 0 {
		return
	}
	if y > MAP_H || y < 0 {
		return
	}
	//	if getEntropy(wave[y][x]) == 1 {
	//		return
	//	}
	//	fmt.Println("avoided")
	var reducedEntropy bool
	idxCurrentTile := getIndex((*wave)[y][x])

	var tops [][TOTAL_TILES]bool
	var lefts [][TOTAL_TILES]bool
	var downs [][TOTAL_TILES]bool
	var rights [][TOTAL_TILES]bool

	combinedTop := models[idxCurrentTile].Top
	combinedLeft := models[idxCurrentTile].Left
	combinedDown := models[idxCurrentTile].Down
	combinedRight := models[idxCurrentTile].Right

	//	fmt.Println(idxCurrentTile)
	if getEntropy(wave[y][x]) != 1 {
		for i, possible := range wave[y][x] {
			if possible {
				tops = append(tops, models[i].Top)
				lefts = append(lefts, models[i].Left)
				downs = append(downs, models[i].Down)
				rights = append(rights, models[i].Right)
			}
		}
		//		fmt.Println(tops, "\n", lefts, "\n", downs, "\n", rights)
		combinedTop = boolTrue(tops)
		combinedLeft = boolTrue(lefts)
		combinedRight = boolTrue(rights)
		combinedDown = boolTrue(downs)
	}
	//	fmt.Println(combinedTop, "\n", combinedLeft, "\n", combinedDown, "\n", combinedRight)

	if y+1 < MAP_H && y+1 >= 0 && getEntropy(wave[y+1][x]) != 1 {
		//		fmt.Println("D", wave[y+1][x], combinedDown)
		//		fmt.Println(boolAnd((*wave)[y+1][x], combinedDown))
		wave[y+1][x], reducedEntropy = boolAnd((*wave)[y+1][x], combinedDown)
		//		fmt.Println("Down", wave[y+1][x])
		if reducedEntropy {
			propagate(wave, x, y+1, models, state)
		}
	}
	if y-1 >= 0 && y-1 < MAP_H && getEntropy(wave[y-1][x]) != 1 {
		//		fmt.Println("T", wave[y-1][x], combinedTop)
		//		fmt.Println(boolAnd((*wave)[y-1][x], combinedTop))
		wave[y-1][x], reducedEntropy = boolAnd((*wave)[y-1][x], combinedTop)
		//		fmt.Println("Top", wave[y-1][x])
		if reducedEntropy {
			propagate(wave, x, y-1, models, state)
		}
	}
	if x+1 < MAP_W && x+1 >= 0 && getEntropy(wave[y][x+1]) != 1 {
		//		fmt.Println("R", wave[y][x+1], combinedRight)
		//		fmt.Println(boolAnd((*wave)[y][x+1], combinedRight))
		wave[y][x+1], reducedEntropy = boolAnd((*wave)[y][x+1], combinedRight)
		//		fmt.Println("Right", wave[y][x+1])
		if reducedEntropy {
			propagate(wave, x+1, y, models, state)
		}
	}
	if x-1 >= 0 && x-1 < MAP_W && getEntropy(wave[y][x-1]) != 1 {
		//		fmt.Println("L", wave[y][x-1], combinedLeft)
		//		fmt.Println(boolAnd((*wave)[y][x-1], combinedLeft))
		wave[y][x-1], reducedEntropy = boolAnd((*wave)[y][x-1], combinedLeft)
		//		fmt.Println("Left", wave[y][x-1])
		if reducedEntropy && len(wave[y][x-1]) != TOTAL_TILES {
			propagate(wave, x-1, y, models, state)
		}
	}
	// fmt.Println("END propagate")
}

// used to  get all the possible tiles for a certain direction
func boolTrue(s [][TOTAL_TILES]bool) [TOTAL_TILES]bool {
	var combinedSlice [TOTAL_TILES]bool
outer:
	for i := range TOTAL_TILES {
		for _, v := range s {
			if v[i] {
				combinedSlice[i] = true
				continue outer
			}
		}
	}
	return combinedSlice
}

func boolAnd(a, b [TOTAL_TILES]bool) (c [TOTAL_TILES]bool, reducedEntropy bool) {
	reducedEntropy = false
	for i := range a {
		if a[i] == b[i] || (a[i] != b[i] && a[i] == false) {
			c[i] = a[i]
			continue
		}
		//		fmt.Println("reducedEntropy")
		reducedEntropy = true
	}
	return c, reducedEntropy
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
				coordinates = [][2]int{{x, y}}
				lowestEntropy = cell_Entropy
			}
		}
	}
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
