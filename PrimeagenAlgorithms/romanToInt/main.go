package main

import "fmt"

func romanToInt(s string) int {
	var total int
	translator := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var lo int = 0
	var hi int = 1
	for lo < len(s) {
		var hVal int
		lVal := translator[s[lo]]
		hVal = 0

		if hi < len(s) {
			hVal = translator[s[hi]]
		}

		fmt.Println("Vals:", lVal, hVal)
		fmt.Println("Total:", total)
		if lVal < hVal {
			fmt.Println("h-l", hVal-lVal)
			total += hVal - lVal
			lo++
			hi++
		} else {
			total += lVal
		}
		fmt.Println("2-Total:", total, "\n")
		lo++
		hi++
	}
	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("I"))
}
