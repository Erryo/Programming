package main

import (
	"math/rand"
	"time"

	"seehuhn.de/go/ncurses"

	"github.com/Erryo/WFC_v2/connections"
)

const (
	W           int = 10
	H           int = 10
	TOTAL_ICONS     = 5
)

var win ncurses.Window = *ncurses.Init()

func main() {
	programStartTime := time.Now()
	connections.Side()
	table := createAndSetupTable()
	drawGrid(table)
	// Pick random to collapse
	y, x := rand.Intn(H), rand.Intn(W)
	z := rand.Intn(TOTAL_ICONS)
	table[y][x].Entropy = 0
	table[y][x].Symbol = table[y][x].Possibilities[z]
	table[y][x].Possibilities = []connections.Symbol{}
	drawGrid(table)
	//propagate
	//...?

	// End
	win.Move(0, 0)
	win.Print(time.Since(programStartTime))
	time.Sleep(2 * time.Second)
	win.Refresh()
}

func createAndSetupTable() [H][W]connections.Tile {
	var table [H][W]connections.Tile
	var table_row [W]connections.Tile
	Possibilities := []connections.Symbol{connections.Down, connections.Right, connections.Up, connections.Left, connections.Empty}
	tile := connections.Tile{Entropy: TOTAL_ICONS, Symbol: connections.Symbol{}, Possibilities: Possibilities}
	for i := range table_row {
		table_row[i] = tile
	}
	for i := range table {
		table[i] = table_row
	}
	//table := make([][]connections.Tile, H)
	//table_row := make([]connections.Tile, W)
	//for i := range table_row {
	//	table_row[i].Entropy = 5
	//}
	//for i := range H {
	//	table[i] = table_row
	//}
	return table
}

func drawGrid(table [H][W]connections.Tile) {
	height, width := win.GetMaxYX()
	// Some sort of 'Off by 1 error,, if  one is not subtracted from both it does not display properly
	height -= 1
	width -= 1
	// TO DO: fix the overflow off the screen when the Table is big
	win.Move((height-H)/2, (width-W)/2)

	for i, row := range table {
		for _, tile := range row {
			if tile.Entropy != 0 {
				win.Print(tile.Entropy, " ")
			} else {
				win.Print(tile.Symbol.Icon)
			}

			win.Refresh()
		}
		win.Println("")
		win.Move(((height-H)/2)+i+1, (width-W)/2)
	}
	win.Refresh()
}
