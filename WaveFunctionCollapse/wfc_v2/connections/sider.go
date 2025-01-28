package connections

import "fmt"

const (
	UP    string = "┴"
	RIGHT string = "├"
	DOWN  string = "┬"
	LEFT  string = "┤"
	EMPTY string = "▒"
)

var (
	Up    symbol
	Right symbol
	Down  symbol
	Left  symbol
	Empty symbol
)

type Tile struct {
	Entropy uint8
	Symbol  symbol
}

type symbol struct {
	Icon                  string
	up, down, left, right []string
}

// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code) :
func Side() {
	Up = symbol{Icon: UP}
	Right = symbol{Icon: RIGHT}
	Down = symbol{Icon: DOWN}
	Left = symbol{Icon: LEFT}
	Empty = symbol{Icon: EMPTY}

	sUp := [4]uint8{1, 1, 0, 1}
	sRight := [4]uint8{1, 1, 1, 0}
	sDown := [4]uint8{0, 1, 1, 1}
	sLeft := [4]uint8{1, 0, 1, 1}
	sEmpty := [4]uint8{0, 0, 0, 0}

	Up, Right = connect(Up, Right, sUp, sRight)
	Up, Left = connect(Up, Left, sUp, sLeft)
	//[[I1]]
	Up, _ = connect(Up, Up, sUp, sUp)
	Up, Down = connect(Up, Down, sUp, sDown)
	Up, Empty = connect(Up, Empty, sUp, sEmpty)

	Right, Left = connect(Right, Left, sRight, sLeft)
	Right, Down = connect(Right, Down, sRight, sDown)
	Right, _ = connect(Right, Right, sRight, sRight)
	Right, Empty = connect(Right, Empty, sRight, sEmpty)

	Down, _ = connect(Down, Down, sDown, sDown)
	Down, Left = connect(Down, Left, sDown, sLeft)
	Down, Empty = connect(Down, Empty, sDown, sEmpty)

	Left, _ = connect(Left, Left, sLeft, sLeft)
	Left, Empty = connect(Left, Empty, sLeft, sEmpty)

	Empty, _ = connect(Empty, Empty, sEmpty, sEmpty)
}

// 0 = up
// 1 = right
// 2 = down
// 3 = left
// Creates the connections for the given 2 symbols
// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code)
// anything written to the victim gets overwritten by the Target when the 2 are the same symbol [[I1]]
func connect(victim, target symbol, victimData, targetData [4]uint8) (symbol, symbol) {
	if victim.Icon != target.Icon {
		if targetData[2] == victimData[0] {
			target.down = append(target.down, victim.Icon)
		}
		if targetData[3] == victimData[1] {
			target.left = append(target.left, victim.Icon)
		}
		if targetData[0] == victimData[2] {
			target.up = append(target.up, victim.Icon)
		}
		if targetData[1] == victimData[3] {
			target.right = append(target.right, victim.Icon)
		}

	}
	if targetData[2] == victimData[0] {
		victim.up = append(victim.up, target.Icon)
	}
	if targetData[3] == victimData[1] {
		victim.right = append(victim.right, target.Icon)
	}
	if targetData[0] == victimData[2] {
		victim.down = append(victim.down, target.Icon)
	}
	if targetData[1] == victimData[3] {
		victim.left = append(victim.left, target.Icon)
	}
	return victim, target
}

// For testing purposes only
func DrawAllConnections() {
	allSymbols := []symbol{Up, Down, Left, Right, Empty}

	for _, currentSymbol := range allSymbols {

		for _, neigbourSymbol := range currentSymbol.up {
			fmt.Print(neigbourSymbol, " ")
		}
		fmt.Println()
		for range currentSymbol.up {
			fmt.Print(currentSymbol.Icon, " ")
		}
		fmt.Println("\n----D")

		for range currentSymbol.down {
			fmt.Print(currentSymbol.Icon, " ")
		}
		fmt.Println()
		for _, neigbourSymbol := range currentSymbol.down {
			fmt.Print(neigbourSymbol, " ")
		}
		fmt.Println("\n----L")

		for _, neigbourSymbol := range currentSymbol.left {
			fmt.Print(neigbourSymbol, currentSymbol.Icon)
			fmt.Print("  ")
		}
		fmt.Println("\n----R")
		for _, neigbourSymbol := range currentSymbol.right {
			fmt.Print(currentSymbol.Icon, neigbourSymbol)
			fmt.Print("  ")
		}

		fmt.Println("\n======================", currentSymbol, "===================")
	}
}
