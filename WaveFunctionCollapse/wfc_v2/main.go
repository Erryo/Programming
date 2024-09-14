package main

import "fmt"

var numberToIcon map[uint8]string = map[uint8]string{
	0: "",
	1: "",
	2: "",
	3: "",
	4: "",
	5: "",
}

func main() {
	box := createBox(7, 5)
	drawBox(box)
	// Array of values - mat
	// Dictionary to draw
	//
}

// Randomness
// If there is a equal choice then it should be random
// Propagation !!!

func createBox(width, length uint8) [][]uint8 {
	var box [][]uint8
	for range width {
		var row []uint8
		for range length {
			row = append(row, 0)
		}
		box = append(box, row)
	}
	return box
}

func drawBox(box [][]uint8) {
	for _, row := range box {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
