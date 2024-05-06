package main

import (
	"fmt"
	"os"
)

func appendToFile(Name, EPass string) {
	var UserData string = Name + "\n" + EPass + "\n"
	file, err := os.OpenFile("DataFile.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("appendToFile->OpenFile: ", err)
	}
	defer file.Close()
	_, err = file.WriteString(UserData)
	if err != nil {
		fmt.Println("appendToFile->WriteString: ", err)
	}
}

func CreateUser(Name, Pass string) {
	var file_exists bool
	var DataAsLines map[string]string
	for {
		file_exists, DataAsLines = ReadFile()
		if !file_exists {
			CreateFile(Name, Pass)
		}
		break
	}
	var Encr_Pass string = Encrypt(Pass, 9)
	_, userExists := DataAsLines[Name]
	if !userExists {
		appendToFile(Name, Encr_Pass)
	} else {
		fmt.Println("User: ", Name, "Already exists")
	}
}
