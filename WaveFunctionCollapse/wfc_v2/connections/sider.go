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
	entropy uint8
	symbol  symbol
}

type symbol struct {
	icon                  string
	up, down, left, right []string
}

// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code) :
func Side() {
	Up = symbol{icon: UP}
	Right = symbol{icon: RIGHT}
	Down = symbol{icon: DOWN}
	Left = symbol{icon: LEFT}
	Empty = symbol{icon: EMPTY}

	sUp := [4]uint8{1, 1, 0, 1}
	sRight := [4]uint8{1, 1, 1, 0}
	sDown := [4]uint8{0, 1, 1, 1}
	sLeft := [4]uint8{1, 0, 1, 1}
	sEmpty := [4]uint8{0, 0, 0, 0}

	Up, Right = connect(Up, Right, sUp, sRight)
	Up, Left = connect(Up, Left, sUp, sLeft)
	Up, Up = connect(Up, Up, sUp, sUp)
	Up, Down = connect(Up, Down, sUp, sDown)
	Up, Empty = connect(Up, Empty, sUp, sEmpty)

	Right, Left = connect(Right, Left, sRight, sLeft)
	Right, Down = connect(Right, Down, sRight, sDown)
	Right, Right = connect(Right, Right, sRight, sRight)
	Right, Empty = connect(Right, Empty, sRight, sEmpty)

	Down, Down = connect(Down, Down, sDown, sDown)
	Down, Left = connect(Down, Left, sDown, sLeft)
	Down, Empty = connect(Down, Empty, sDown, sEmpty)

	Left, Left = connect(Left, Left, sLeft, sLeft)
	Left, Empty = connect(Left, Empty, sLeft, sEmpty)

	Empty, Empty = connect(Empty, Empty, sEmpty, sEmpty)
}

// 0 = up
// 1 = right
// 2 = down
// 3 = left
// Creates the connections for the given 2 symbols
// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code)
func connect(victim, target symbol, victimData, targetData [4]uint8) (symbol, symbol) {
	if victim.icon != target.icon {
		if targetData[2] == victimData[0] {
			target.down = append(target.down, victim.icon)
		}
		if targetData[3] == victimData[1] {
			target.left = append(target.left, victim.icon)
		}
		if targetData[0] == victimData[2] {
			target.up = append(target.up, victim.icon)
		}
		if targetData[1] == victimData[3] {
			target.right = append(target.right, victim.icon)
		}

	}
	if targetData[2] == victimData[0] {
		victim.up = append(victim.up, target.icon)
	}
	if targetData[3] == victimData[1] {
		victim.right = append(victim.right, target.icon)
	}
	if targetData[0] == victimData[2] {
		victim.down = append(victim.down, target.icon)
	}
	if targetData[1] == victimData[3] {
		victim.left = append(victim.left, target.icon)
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
			fmt.Print(currentSymbol.icon, " ")
		}
		fmt.Println("\n----")

		for range currentSymbol.down {
			fmt.Print(currentSymbol.icon, " ")
		}
		fmt.Println()
		for _, neigbourSymbol := range currentSymbol.down {
			fmt.Print(neigbourSymbol, " ")
		}
		fmt.Println("\n----")

		for _, neigbourSymbol := range currentSymbol.left {
			fmt.Print(neigbourSymbol, currentSymbol.icon)
			fmt.Print("  ")
		}
		fmt.Println("\n----")
		for _, neigbourSymbol := range currentSymbol.left {
			fmt.Print(currentSymbol.icon, neigbourSymbol)
			fmt.Print("  ")
		}

		fmt.Println("\n======================", currentSymbol, "===================")
	}
}
