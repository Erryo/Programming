package wfcv2

func main() {
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
