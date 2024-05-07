package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func CheckData(user User, filepath string) bool {
	var readUsers []User = GetJson(filepath)
	for _, value := range readUsers {
		if value.Username == user.Username {
			if Decrypt(value.Password, 9) == user.Password {
				return true
			}
		}
	}
	return false
}

func UserExists(allUsers []User, user User) bool {
	for _, value := range allUsers {
		if value.Username == user.Username {
			return true
		}
	}
	return false
}

func AppendToArray(filepath string, user User) string {
	var readUsers []User = GetJson(filepath)
	if !UserExists(readUsers, user) {
		user.Password = Encrypt(user.Password, 9)
		readUsers = append(readUsers, user)
		ConvertToJson(readUsers, filepath)
		return ""
	} else {
		fmt.Printf("User:%v already exists \n", user.Username)
		return "User already exists!"
	}
}

func GetJson(filepath string) []User {
	var users []User
	readData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("GetJson-> os.ReadFile, ", err)
	}
	err = json.Unmarshal(readData, &users)
	if err != nil {
		fmt.Println("GetJson->json.Unmarshal, ", err)
	}
	return users
}

func CreateJsonFile(filepath string) {
	fmt.Println("Func CreateJsonFile started")
	_, err := os.Create(filepath)
	if err != nil {
		fmt.Println("CreateJsonFile->Error: ", err)
	}
}

func WriteToFile(jsonData []byte, filepath string) {
	file, err := os.OpenFile(filepath, os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("WriteToFile->OpenFile: ", err)
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("WriteToFile->WriteString: ", err)
	}
}

func ConvertToJson(users []User, filepath string) {
	fmt.Println("ConvertToJson started")
	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("ConvertToJson->json.Marshal: ", err)
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("While opening ", filepath, " ,Error: ", err)
		CreateJsonFile(filepath)
	} else {
		file.Close()
	}
	WriteToFile(jsonData, filepath)
}
