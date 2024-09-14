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
	lNo := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(lNo, line)
		line = transformString(line)
		fmt.Println(lNo, line, getNumOfLine(line))

		sum += getNumOfLine(line)
		lNo++
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return 0
	}
	return sum
}

func transformString(line string) string {
	//zero := regexp.MustCompile(`(zero)`)
	//one := regexp.MustCompile(`(one)`)
	//two := regexp.MustCompile(`(two)`)
	//three := regexp.MustCompile(`(three)`)
	//four := regexp.MustCompile(`(four)`)
	//five := regexp.MustCompile(`(five)`)
	//six := regexp.MustCompile(`(six)`)
	//seven := regexp.MustCompile(`(seven)`)
	//eight := regexp.MustCompile(`(eight)`)
	//nine := regexp.MustCompile(`(nine)`)
	//line = zero.ReplaceAllString(line, "0")
	//line = one.ReplaceAllString(line, "1")
	//line = two.ReplaceAllString(line, "2")
	//line = three.ReplaceAllString(line, "3")
	//line = four.ReplaceAllString(line, "4")
	//line = five.ReplaceAllString(line, "5")
	//line = six.ReplaceAllString(line, "6")
	//line = seven.ReplaceAllString(line, "7")
	//line = eight.ReplaceAllString(line, "8")
	//line = nine.ReplaceAllString(line, "9")
	letters := make([][]string, 10)
	letters[0] = []string{"z", "e", "r", "o"}
	letters[1] = []string{"o", "n", "e"}
	letters[2] = []string{"t", "w", "o"}
	letters[3] = []string{"t", "h", "r", "e", "e"}
	letters[4] = []string{"f", "o", "u", "r"}
	letters[5] = []string{"f", "i", "v", "e"}
	letters[6] = []string{"s", "i", "x"}
	letters[7] = []string{"s", "e", "v", "e", "n"}
	letters[8] = []string{"e", "i", "g", "h", "t"}
	letters[9] = []string{"n", "i", "n", "e"}
	lnCopy := []byte(line)
	for ind, char := range line {
		for i, letter := range letters {
			if string(char) == letter[0] {
				lnCopy[ind] = byte(i)
			}
		}
	}
	return string(lnCopy)
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
