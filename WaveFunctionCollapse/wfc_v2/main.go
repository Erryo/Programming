package main

import (
	"fmt"
	"time"

	"github.com/Erryo/WFC_v2/connections"
)

const (
	W int = 10
	H int = 10
)

func main() {
	programStartTime := time.Now()
	connections.Side()
	createAndSetupTable()

	fmt.Println(time.Since(programStartTime))
}

func createAndSetupTable() {
	table := make([][]connections.Tile, H)
	table_row := make([]connections.Tile, W)
	for i := range table_row {
		table_row[i].Entropy = 5
	}
	for i := range H {
		table[i] = table_row
	}
	fmt.Println(table)
}
