package main

import "core:encoding/json"
import "core:fmt"
import "core:os"

printWave :: proc(wave: ^wave) {
	for i in -1 ..< WAVE_WIDTH {
		if i != -1 {
			fmt.print(i, " ")
			continue
		}

		fmt.print("-| ")
	}
	fmt.println()

	entropy: u8
	for i in 0 ..< WAVE_HEIGTH {
		fmt.print(i, " ")
		for j in 0 ..< WAVE_WIDTH {
			fmt.print(getEntropy(wave[i][j]), " ")
		}
		fmt.println()
	}
	fmt.println("----------")
}

getEntropy :: proc(cell: tile) -> u8 {
	entropy: u8
	for value, _ in cell {
		if value {
			entropy += 1
		}
	}
	return entropy
}

boolAnd :: proc(a, b: tile) -> (tile, bool) {
	reducedEntropy: bool
	c: tile
	for i in 0 ..< TOTAL_TILES {
		if a[i] == b[i] || (a[i] != b[i] && a[i] == false) {
			c[i] = a[i]
			continue
		}
		reducedEntropy = true
	}
	return c, reducedEntropy

}
boolOr :: proc(a, b: tile) -> tile {
	a := a
	for i in 0 ..< TOTAL_TILES {
		if b[i] {
			a[i] = b[i]
			continue
		}
	}
	return a

}

boolTrue :: proc(a: []tile) -> tile {
	result: tile
	outer: for i in 0 ..< TOTAL_TILES {
		for v, _ in a {
			if v[i] {
				result[i] = true
				continue outer
			}

		}
	}
	return result
}
initModels :: proc() -> [TOTAL_TILES]tileModel {
	models: [TOTAL_TILES]tileModel
	for i in 0 ..< TOTAL_TILES {
		models[i].index = i
	}
	return models

}

loadTileModels :: proc() -> [TOTAL_TILES]tileModel {
	models: [TOTAL_TILES]tileModel

	data, err := os.read_entire_file_from_filename_or_err(PATH_TO_MODELS)
	if err != nil {
		fmt.panicf("%v", err)

	}
	defer delete(data)

	marshal_err := json.unmarshal(data, &models)
	if marshal_err != nil {
		fmt.panicf("%v", err)
	}
	return models
}

loadTileShapes :: proc() -> [TOTAL_TILES][9]u8 {
	shapes: [TOTAL_TILES][9]u8

	data, err := os.read_entire_file_from_filename_or_err(PATH_TO_SHAPES)
	if err != nil {
		fmt.panicf("%v", err)
	}
	defer delete(data)

	marshal_err := json.unmarshal(data, &shapes)
	if marshal_err != nil {
		fmt.panicf("%v", err)
	}

	return shapes
}

marshalModels :: proc(models: [TOTAL_TILES]tileModel) -> []byte {
	data, err := json.marshal(models)
	if err != nil {
		fmt.panicf("error marshalling models: %v", err)
	}

	return data
}

writeTileModels :: proc(models: [TOTAL_TILES]tileModel) {
	data := marshalModels(models)
	err := os.write_entire_file_or_err(PATH_TO_MODELS, data, true)
	if err != nil {
		fmt.panicf("error writing tile models:%v", err)
	}
	defer delete(data)
}

createConnections :: proc() {
	shapes := loadTileShapes()
	models := initModels()

	for i in 0 ..< TOTAL_TILES {
		for j in i ..< TOTAL_TILES {
			connect(i, j, shapes, &models)
		}
	}
	writeTileModels(models)
}

checkValid :: proc(a, b: [9]u8, offsetA, offsetB, limit, increase: int) -> bool {
	for i := 0; i <= limit; i += increase {
		if a[offsetA + i] != b[offsetB + i] {
			return false
		}
	}


	return true
}

getIndex :: proc(cell: tile) -> int {
	for i in 0 ..< TOTAL_TILES {
		if cell[i] {
			return i
		}
	}
	return -1
}


connect :: proc(idxA, idxB: int, shapes: [TOTAL_TILES][9]u8, models: ^[TOTAL_TILES]tileModel) {
	a, b: [9]u8
	a = shapes[idxA]
	b = shapes[idxB]
	modelA, modelB := &models[idxA], &models[idxB]
	// Check A's Top with B's Bottom
	if checkValid(a, b, 0, 6, 2, 1) {
		modelA.top[idxB] = true
		modelB.down[idxA] = true
	}

	// Check A's Bottom with B's Top
	if checkValid(a, b, 6, 0, 2, 1) {
		modelA.down[idxB] = true
		modelB.top[idxA] = true
	}

	// Check A's Left with B's Right  
	if checkValid(a, b, 0, 2, 6, 3) {
		modelA.left[idxB] = true
		modelB.right[idxA] = true
	}

	// Check A's reght with B's left   
	if checkValid(a, b, 2, 0, 6, 3) {
		modelA.right[idxB] = true
		modelB.left[idxA] = true
	}
}
