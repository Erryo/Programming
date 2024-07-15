package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Ascii digits between  [48,57]
func SumOfVals() int {
	var sum int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += getNumOfLine(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return 0
	}
	return sum
}

func getNumOfLine(line string) int {
	var num string
	for hi := len(line) - 1; hi >= len(line)/2; {
		for lo := 0; lo < len(line)+1/2; {
			if line[lo] > 57 || line[lo] < 48 {
				lo++
				continue
			}
			if line[hi] > 57 || line[hi] < 48 {
				hi--
				continue
			}
			num = string(line[lo]) + string(line[hi])
			break
		}
		break
	}

	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return number
}
