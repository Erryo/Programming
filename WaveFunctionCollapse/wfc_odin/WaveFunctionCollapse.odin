package main

import "core:fmt"
import "core:math/rand"
import "core:os"
import "core:time"
import sdl "vendor:sdl3"
import img "vendor:sdl3/image"

main :: proc() {
	args := os.args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "side":
			createConnections()
		case "debug":
			Debug = true
		}

	}
	state: ^state = nil
	initSDL()
	initState(&state)
	wave := startWFC(state)

	saveToFile(state)

	closeSDL(&state)
	free(wave)
}

startWFC :: proc(s: ^state) -> ^wave {
	startTime := time.now()
	wave := initWave()
	models := loadTileModels()
	s.wave = wave

	cellPosition := getRandomCell(wave)
	collapseCell(&wave[cellPosition.y][cellPosition.x])
	propagate(wave, cellPosition.x, cellPosition.y, &models, s)

	coordinates: [dynamic]Vector2
	finished: bool
	defer delete(coordinates)
	for {
		finished = getLowestEntropy(wave, &coordinates)
		if finished {
			break
		}

		if len(coordinates) == 0 {fmt.panicf("why  0")}

		cellPosition = coordinates[rand.int_max(len(coordinates))]
		collapseCell(&wave[cellPosition.y][cellPosition.x])
		if propagate(wave, cellPosition.x, cellPosition.y, &models, s) {
			return wave
		}


		clear(&coordinates)

		if doInput(s) {
			fmt.println(time.since(startTime))
			return wave
		}
		drawWave(s)
	}

	fmt.println(time.since(startTime))
	return wave
}

propagate :: proc(wave: ^wave, x, y: int, models: ^[TOTAL_TILES]tileModel, s: ^state) -> bool {
	append(&s.trace, Vector2{x, y})
	traceIdx := len(s.trace) - 1
	if Debug && waitForSpace(s) {
		return true
	}
	reducedEntropy: bool
	idxCurrentTile := getIndex(wave^[y][x])
	if idxCurrentTile == -1 {
		return false
	}

	top := models[idxCurrentTile].top
	left := models[idxCurrentTile].left
	down := models[idxCurrentTile].down
	right := models[idxCurrentTile].right

	if getEntropy(wave[y][x]) != 1 {
		for b, i in wave[y][x] {
			if b {
				top = boolOr(top, models[i].top)
				left = boolOr(left, models[i].left)
				down = boolOr(down, models[i].down)
				right = boolOr(right, models[i].right)
			}

		}
	}

	if y + 1 < WAVE_HEIGTH && y + 1 >= 0 && getEntropy(wave[y + 1][x]) != 1 {
		wave[y + 1][x], reducedEntropy = boolAnd(wave^[y + 1][x], down)
		if reducedEntropy {
			if propagate(wave, x, y + 1, models, s) {return true}
		}
	}
	if y - 1 >= 0 && y - 1 < WAVE_HEIGTH && getEntropy(wave[y - 1][x]) != 1 {
		wave[y - 1][x], reducedEntropy = boolAnd(wave^[y - 1][x], top)
		if reducedEntropy {
			if propagate(wave, x, y - 1, models, s) {return true}
		}
	}
	if x + 1 < WAVE_WIDTH && x + 1 >= 0 && getEntropy(wave[y][x + 1]) != 1 {
		wave[y][x + 1], reducedEntropy = boolAnd(wave^[y][x + 1], right)
		if reducedEntropy {
			if propagate(wave, x + 1, y, models, s) {return true}
		}
	}
	if x - 1 >= 0 && x - 1 < WAVE_WIDTH && getEntropy(wave[y][x - 1]) != 1 {
		wave[y][x - 1], reducedEntropy = boolAnd(wave^[y][x - 1], left)
		if reducedEntropy {
			if propagate(wave, x - 1, y, models, s) {return true}
		}

	}
	unordered_remove(&s.trace, traceIdx)
	return false
}

initWave :: proc() -> ^wave {
	wave := new(wave)
	row: [WAVE_WIDTH]tile
	row = true
	for i in 0 ..< WAVE_HEIGTH {
		wave[i] = row
	}

	return wave
}


//don't use on uncollapsed wave
getLowestEntropy :: proc(wave: ^wave, coords: ^[dynamic]Vector2) -> bool {
	lowestEntropy: u8
	cellEntropy: u8
	allCollapsed: bool = true
	lowestEntropy = TOTAL_TILES

	for y in 0 ..< WAVE_HEIGTH {
		row: for x in 0 ..< WAVE_WIDTH {

			cellEntropy = getEntropy(wave[y][x])
			if cellEntropy == 1 {
				continue row
			}

			allCollapsed = false

			if cellEntropy > lowestEntropy {continue row}


			if cellEntropy < lowestEntropy {
				clear(coords)
				lowestEntropy = cellEntropy
			}

			append(coords, Vector2{x, y})

		}
	}

	return allCollapsed
}

getRandomCell :: proc(wave: ^wave) -> Vector2 {
	cell: Vector2
	for {
		cell.x = rand.int_max(WAVE_WIDTH)
		cell.y = rand.int_max(WAVE_HEIGTH)
		if getEntropy(wave[cell.y][cell.x]) != 1 {
			break
		}

	}
	return cell

}

collapseCell :: proc(cell: ^tile) {
	if cell == nil {
		return
	}
	idx: int

	for {
		idx = rand.int_max(TOTAL_TILES)
		if cell[idx] {
			cell^ = tile{}
			cell[idx] = true
			return
		}

	}

}
