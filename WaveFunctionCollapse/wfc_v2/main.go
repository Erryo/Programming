package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"seehuhn.de/go/ncurses"

	"github.com/Erryo/WFC_v2/connections"
)

const (
	W int = 10
	H int = 10
)

var win ncurses.Window = *ncurses.Init()

func main() {
	LOG_FILE := "/tmp/WFC_v2_log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)
	programStartTime := time.Now()
	connections.Side()
	table := createAndSetupTable()
	drawGrid(&table)
	// Pick random to collapse
	y, x := rand.Intn(H), rand.Intn(W)
	z := rand.Intn(connections.TOTAL_ICONS)
	table[y][x].Entropy = 0
	table[y][x].Symbol = connections.IconSymbolMap[table[y][x].Possibilities[z]]
	table[y][x].Possibilities = []string{table[y][x].Symbol.Icon}

	// propagate
	Bio(&table, x, y)
	drawGrid(&table)
	//...?
	win.Refresh()

	// End
	win.Move(0, 0)
	win.Print(time.Since(programStartTime))
	time.Sleep(2 * time.Second)
	win.Refresh()
}

// Recursive
func Bio(table *([H][W]connections.Tile), x, y int) {
	// Check bounds.

	self := table[y][x]
	self.Entropy = uint8(len(self.Possibilities))
	var modified bool
	// return bool if need be continued ?
	if x+1 < W {
		// Possible issue : not actually changing the values in the table
		log.Print("Bio Right")
		possibilities := unifyPossibilities(self, 3)
		log.Print(possibilities)
		table[y][x+1].Possibilities, modified = simpleAND(table[y][x+1].Possibilities, possibilities)
		if modified {
			Bio(table, x+1, y)
		}
	}
	if x-1 >= 0 {
		possibilities := unifyPossibilities(self, 1)
		table[y][x-1].Possibilities, modified = simpleAND(table[y][x-1].Possibilities, possibilities)
		if modified {
			Bio(table, x-1, y)
		}
	}
	if y+1 < H {
		possibilities := unifyPossibilities(self, 0)
		table[y+1][x].Possibilities, modified = simpleAND(table[y+1][x].Possibilities, possibilities)
		if modified {
			Bio(table, x, y+1)
		}
	}
	if y-1 >= 0 {
		possibilities := unifyPossibilities(self, 2)
		table[y-1][x].Possibilities, modified = simpleAND(table[y-1][x].Possibilities, possibilities)
		if modified {
			Bio(table, x, y+1)
		}
	}
	return

	// table[y+1][x] // T
	// table[y][x-1] // L
	// table[y-1][x] // D
	// table[y][x+1] // R
}

// Create new array to unify all Possibilities
// Takes in the Tile, uses only its Possibilities , to create one list of possible symbols for the next tile
// 0 T , 1 L , 2 D , 3 R
func unifyPossibilities(a connections.Tile, direction uint8) []string {
	tempMap := make(map[string]string)
	result := []string{}
	for _, symbol := range a.Possibilities {
		if direction == 0 {
			for _, icon := range connections.IconSymbolMap[symbol].Up {
				tempMap[icon] = icon
			}
		}
		if direction == 1 {
			for _, icon := range connections.IconSymbolMap[symbol].Left {
				tempMap[icon] = icon
			}
		}
		if direction == 2 {
			for _, icon := range connections.IconSymbolMap[symbol].Down {
				tempMap[icon] = icon
			}
		}
		if direction == 3 {
			for _, icon := range connections.IconSymbolMap[symbol].Right {
				tempMap[icon] = icon
			}
		}
	}
	for _, val := range tempMap {
		result = append(result, val)
	}
	return result
}

// Purpose: Change the Possibilities of a tile to fit the ones of another tile , i.e. Push its Possibilities
// a is currentTile.Symbol.up (what it can accept in up) , b is the Possibilities of the next tile
// Req: Further Testing
// Issue when using with multiple Possibilities: It is Possible for first Possibility(P.) to accept only 3 Symbols but
// the next P. to contain 4, due to the implementation that extra Symbol will miss
func simpleAND(a, b []string) ([]string, bool) {
	// keep track if it made any changes
	var modified bool
	result := []string{}
outer:
	for _, Ai := range a {
		for _, Bi := range b {
			if Ai == Bi {
				modified = true
				result = append(result, Ai)
				continue outer
			}
		}
	}
	return result, modified
}

func createAndSetupTable() [H][W]connections.Tile {
	var table [H][W]connections.Tile
	var table_row [W]connections.Tile
	Possibilities := []string{connections.DOWN, connections.RIGHT, connections.UP, connections.LEFT, connections.EMPTY}
	tile := connections.Tile{Entropy: connections.TOTAL_ICONS, Symbol: connections.Symbol{}, Possibilities: Possibilities}
	for i := range table_row {
		table_row[i] = tile
	}
	for i := range table {
		table[i] = table_row
	}
	return table
}

func drawGrid(table *[H][W]connections.Tile) {
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
