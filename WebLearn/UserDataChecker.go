package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func CreateFile(SubName, SubPass string) {
	fmt.Println("Func CreateFile started")
	file, err := os.Create("DataFile.txt")
	if err != nil {
		fmt.Println("At: CreateFile, Error: ", err)
	}
	var UserData string = SubName + "\n" + SubPass
	_, err = file.WriteString(UserData)
	if err != nil {
		fmt.Println("WriteString, Error: ", err)
	}
	fmt.Println("CreateFile->UserData: ", UserData)
}

func ReadFile() bool {
	fmt.Println("Func ReadFile started")
	content, err := os.ReadFile("DataFile.txt")
	if err != nil {
		fmt.Println("While reading file: DataFile.txt,Erro occured:", err)
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	c_str := string(content)
	fmt.Println("cont: ", c_str)
	return true
}

func CheckData(SubName, SubPass string) bool {
	var file_exists bool = ReadFile()
	if !file_exists {
		CreateFile(SubName, SubPass)
	}
	return false
	//Create a file , come up with a storage format
	//Possibly Create Personal encryprion method
	//read file
	//Decrypt or Encrypt
	//Compare
	//return bool
	//
}
