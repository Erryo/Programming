package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type task struct {
	Value  string    `json:"value"`
	Status time.Time `json:"status"`
}

const FILE_DIR = "/home/infy/Documents/.toDoList.json"

var args []string

func createFile() {
	_, err := os.OpenFile(FILE_DIR, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("createFile:", err)
		return
	}
	convertToJson([]task{})
}

func readFile() []task {
	var tasks []task
	if _, err := os.Stat(FILE_DIR); errors.Is(err, os.ErrNotExist) {
		createFile()
	}
	readData, err := os.ReadFile(FILE_DIR)
	if err != nil {
		fmt.Println("readFile->os.ReadFile, ", err)
	}
	err = json.Unmarshal(readData, &tasks)
	if err != nil {
		fmt.Println("readFile->json.Unmarshal, ", err)
	}
	return tasks
}

func writeToFile(jsonData []byte) {
	file, err := os.OpenFile(FILE_DIR, os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("WriteToFile->OpenFile: ", err)
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("WriteToFile->WriteString: ", err)
	}
}

func convertToJson(tasks []task) {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("ConvertToJson->json.Marshal: ", err)
	}
	file, err := os.Open(FILE_DIR)
	if err != nil {
		createFile()
	} else {
		file.Close()
	}
	writeToFile(jsonData)
}

func displayHelp() {
	fmt.Println("<--------IDoToo-------->")
	fmt.Println("-h display this message")
	fmt.Println("-c clear the entire list")
	fmt.Println("-d mark as done, takes the index as an argument(zero based index)")
	fmt.Println("-s show the toDoList")
	fmt.Println("-a add new task")
	fmt.Println("<--------IDidDo-------->")
}

func showList() {
	var longest int
	tasks := readFile()
	cYear, cMonth, cDay := time.Now().Date()
	for _, value := range tasks {
		if len(value.Value) > longest {
			longest = len(value.Value)
		}
	}
	fmt.Printf("<-%vIDoToo%v-->\n", strings.Repeat("-", (longest-5)/2), strings.Repeat("-", (longest-4)/2))
	for index, value := range tasks {
		var elapsed string
		space := strings.Repeat(" ", longest-len(value.Value))

		vYear, vMonth, vDay := value.Status.Date()
		eYear := int(cYear) - int(vYear)
		eMonth := int(cMonth) - int(vMonth)
		eDay := cDay - vDay

		if eYear >= 1 {
			elapsed = strconv.Itoa(eYear)
			elapsed += "y"
		} else if eMonth >= 1 {
			elapsed = strconv.Itoa(eMonth)
			elapsed += "m"
		} else if eDay >= 7 {
			elapsed = strconv.Itoa((eDay) / 7)
			elapsed += "w"
		} else if eDay > 1 && eDay < 7 {
			elapsed = strconv.Itoa(eDay)
			elapsed += "d"
		} else if eDay <= 1 {
			elapsed = strconv.Itoa(time.Now().Hour() - value.Status.Hour())
			elapsed += "h"
		} else {
			elapsed = ""
		}

		fmt.Printf("%v| %v %v|%v\n", index, value.Value, space, elapsed)
	}
	fmt.Printf("<-%vIDidDo%v-->\n", strings.Repeat("-", (longest-4)/2), strings.Repeat("-", (longest-4)/2))
}

func addToDo() {
	var data string = args[1]
	var newTask task = task{Value: data, Status: time.Now()}
	tasks := readFile()
	tasks = append(tasks, newTask)
	convertToJson(tasks)
	showList()
}

func deleteTask(index int) {
	readData := readFile()
	if index >= len(readData) {
		return
	}
	var newData []task
	newData = append(newData, readData[:index]...)
	newData = append(newData, readData[index+1:]...)
	convertToJson(newData)
	showList()
}

func main() {
	args = os.Args[1:]
	// getArgs()
	if len(args) == 0 {
		showList()
		return
	}
	switch args[0] {
	case "-h":
		displayHelp()
	case "-a":
		if len(args) != 2 {
			displayHelp()
			return
		}
		addToDo()
	case "-c":
		convertToJson([]task{})
	case "-s":
		showList()
	case "-d":
		if len(args) != 2 {
			displayHelp()
			return
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Please introduce a valid index")
			return
		}
		deleteTask(index)
	default:
		displayHelp()
	}
}
