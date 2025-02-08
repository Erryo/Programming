package main

import (
	"log"
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
	// Pick random to collapse
	// y, x := rand.Intn(H), rand.Intn(W)
	y, x := 3, 3
	// z := rand.Intn(connections.TOTAL_ICONS)
	z := 3
	self := table[y][x]
	self.Entropy = 0
	self.Symbol = connections.IconSymbolMap[self.Possibilities[z]]
	self.Possibilities = []string{self.Symbol.Icon}

	table[y][x] = self
	drawGrid(&table)
	win.Refresh()
	// propagate

	// NEED:A func to propagate the changes of the initial collapse
	var modified bool
	if x+1 < W-1 {
		// Possible issue : not actually changing the values in the table
		log.Println("R", x, y)
		possibilities := unifyPossibilities(self, 3)
		log.Println(possibilities)
		table[y][x+1].Possibilities, modified = simpleAND(table[y][x+1].Possibilities, possibilities)
		if modified {
			Bio(&table, x+1, y)
		}
	}
	if x-1 >= 0 {
		log.Println("L", x, y)
		possibilities := unifyPossibilities(self, 1)
		table[y][x-1].Possibilities, modified = simpleAND(table[y][x-1].Possibilities, possibilities)
		if modified {
			Bio(&table, x-1, y)
		}
	}
	if y+1 < H-1 {
		log.Println("U", x, y)
		possibilities := unifyPossibilities(self, 0)
		table[y+1][x].Possibilities, modified = simpleAND(table[y+1][x].Possibilities, possibilities)
		if modified {
			Bio(&table, x, y+1)
		}
	}
	if y-1 >= 0 {
		log.Println("D", x, y)
		possibilities := unifyPossibilities(self, 2)
		table[y-1][x].Possibilities, modified = simpleAND(table[y-1][x].Possibilities, possibilities)
		if modified {
			Bio(&table, x, y+1)
		}
	}

	//
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
	self := table[y][x]
	table[y][x].Entropy = uint8(len(self.Possibilities))
	log.Println(x, y, table[y][x].Possibilities, table[y][x].Entropy)
	// return bool if need be continued ?
	if x+1 < W-1 && table[y][x+1].Entropy > 1 {
		// Possible issue : not actually changing the values in the table
		possibilities := unifyPossibilities(self, 3)
		log.Println("r", possibilities)
		result, modified := simpleAND(table[y][x+1].Possibilities, possibilities)
		if modified {
			table[y][x+1].Entropy = uint8(len(result))
			table[y][x+1].Possibilities = result
			Bio(table, x+1, y)
		}
	}
	if x-1 >= 0 && table[y][x-1].Entropy != 0 {
		possibilities := unifyPossibilities(self, 1)
		log.Println("l", possibilities)
		result, modified := simpleAND(table[y][x-1].Possibilities, possibilities)
		if modified {
			table[y][x-1].Entropy = uint8(len(result))
			table[y][x-1].Possibilities = result
			Bio(table, x-1, y)
		}
	}
	if y+1 < H-1 && table[y+1][x].Entropy != 0 {
		possibilities := unifyPossibilities(self, 0)
		log.Println("u", possibilities)
		result, modified := simpleAND(table[y+1][x].Possibilities, possibilities)
		if modified {
			table[y+1][x].Entropy = uint8(len(result))
			table[y+1][x].Possibilities = result
			Bio(table, x, y+1)
		}
	}
	if y-1 >= 0 && table[y-1][x].Entropy != 0 {
		possibilities := unifyPossibilities(self, 2)
		log.Println("d", possibilities)
		result, modified := simpleAND(table[y-1][x].Possibilities, possibilities)
		if modified {
			table[y-1][x].Entropy = uint8(len(result))
			table[y-1][x].Possibilities = result
			Bio(table, x, y-1)
		}
	}

	// table[y+1][x] // T
	// table[y][x-1] // L
	// table[y-1][x] // D
	// table[y][x+1] // R
}

// iterates over each possibility of a tile
// to create a new list of accpted values in a specified direction
// 0 T , 1 L , 2 D , 3 R
func unifyPossibilities(a connections.Tile, direction uint8) []string {
	tempMap := make(map[string]string)
	result := []string{}
	for _, aIcon := range a.Possibilities {
		if direction == 0 {
			for _, icon := range connections.IconSymbolMap[aIcon].Up {
				tempMap[icon] = icon
			}
		}
		if direction == 1 {
			for _, icon := range connections.IconSymbolMap[aIcon].Left {
				tempMap[icon] = icon
			}
		}
		if direction == 2 {
			for _, icon := range connections.IconSymbolMap[aIcon].Down {
				tempMap[icon] = icon
			}
		}
		if direction == 3 {
			for _, icon := range connections.IconSymbolMap[aIcon].Right {
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
	var noMatches uint8
	result := []string{}
outer:
	for _, Ai := range a {
		for _, Bi := range b {
			if Ai == Bi {
				modified = true
				noMatches++
				result = append(result, Ai)
				continue outer
			}
		}
	}
	// if all elements match then there was no change
	if int(noMatches) == len(a) {
		modified = false
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
