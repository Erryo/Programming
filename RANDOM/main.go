package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	dir, err := os.ReadDir("./bin")
	if err != nil {
		log.Fatal(err)
	}
	dirs := map[string]string{}
	for i, file := range dir {
		fileIndx := strconv.Itoa(i)
		dirs[fileIndx] = file.Name()
		fmt.Printf("%v. %v\n", i, file.Name())
	}
	lenDir := len(dirs)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res, err := strconv.Atoi(scanner.Text())
		if res < 0 || res >= lenDir || err != nil {
			fmt.Println("Please insert valid number")
			continue
		}
		str := fmt.Sprintf("./bin/%v", dirs[scanner.Text()])

		cmd := exec.Command(str)
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))
		break
	}
}
