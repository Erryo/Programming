package connections

import "fmt"

const (
	UP          string = "┴"
	RIGHT       string = "├"
	DOWN        string = "┬"
	LEFT        string = "┤"
	EMPTY       string = "▒"
	TOTAL_ICONS        = 5
)

var (
	Up            Symbol
	Right         Symbol
	Down          Symbol
	Left          Symbol
	Empty         Symbol
	IconSymbolMap map[string]Symbol
)

type Tile struct {
	Entropy       uint8
	Symbol        Symbol
	Possibilities []string
}

type Symbol struct {
	Icon                  string
	Up, Down, Left, Right []string
}

// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code) :
func Side() {
	Up = Symbol{Icon: UP}
	Right = Symbol{Icon: RIGHT}
	Down = Symbol{Icon: DOWN}
	Left = Symbol{Icon: LEFT}
	Empty = Symbol{Icon: EMPTY}

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

	// Cannot use this Map because it needs to be Global
	tempMap := make(map[string]Symbol, TOTAL_ICONS)

	tempMap[UP] = Up
	tempMap[LEFT] = Left
	tempMap[DOWN] = Down
	tempMap[RIGHT] = Right
	tempMap[EMPTY] = Empty

	IconSymbolMap = tempMap
}

// 0 = up
// 1 = right
// 2 = down
// 3 = left
// Creates the connections for the given 2 symbols
// DO NOT TOUCH!!!
// 20.01.25 I assume it is done(it has been over 2 months since writing the code)
// anything written to the victim gets overwritten by the Target when the 2 are the same symbol [[I1]]
func connect(victim, target Symbol, victimData, targetData [4]uint8) (Symbol, Symbol) {
	if victim.Icon != target.Icon {
		if targetData[2] == victimData[0] {
			target.Down = append(target.Down, victim.Icon)
		}
		if targetData[3] == victimData[1] {
			target.Left = append(target.Left, victim.Icon)
		}
		if targetData[0] == victimData[2] {
			target.Up = append(target.Up, victim.Icon)
		}
		if targetData[1] == victimData[3] {
			target.Right = append(target.Right, victim.Icon)
		}

	}
	if targetData[2] == victimData[0] {
		victim.Up = append(victim.Up, target.Icon)
	}
	if targetData[3] == victimData[1] {
		victim.Right = append(victim.Right, target.Icon)
	}
	if targetData[0] == victimData[2] {
		victim.Down = append(victim.Down, target.Icon)
	}
	if targetData[1] == victimData[3] {
		victim.Left = append(victim.Left, target.Icon)
	}
	return victim, target
}

// For testing purposes only
func DrawAllConnections() {
	allSymbols := []Symbol{Up, Down, Left, Right, Empty}

	for _, currentSymbol := range allSymbols {

		for _, neigbourSymbol := range currentSymbol.Up {
			fmt.Print(neigbourSymbol, " ")
		}
		fmt.Println()
		for range currentSymbol.Up {
			fmt.Print(currentSymbol.Icon, " ")
		}
		fmt.Println("\n----D")

		for range currentSymbol.Down {
			fmt.Print(currentSymbol.Icon, " ")
		}
		fmt.Println()
		for _, neigbourSymbol := range currentSymbol.Down {
			fmt.Print(neigbourSymbol, " ")
		}
		fmt.Println("\n----L")

		for _, neigbourSymbol := range currentSymbol.Left {
			fmt.Print(neigbourSymbol, currentSymbol.Icon)
			fmt.Print("  ")
		}
		fmt.Println("\n----R")
		for _, neigbourSymbol := range currentSymbol.Right {
			fmt.Print(currentSymbol.Icon, neigbourSymbol)
			fmt.Print("  ")
		}

		fmt.Println("\n======================", currentSymbol, "===================")
	}
}
