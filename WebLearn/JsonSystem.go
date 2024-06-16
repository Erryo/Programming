package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func DeleteUser(Username string, filepath string) bool {
	fmt.Println("DeleteUser: ", Username, " started")
	var userIndex int
	readUsers := GetJson(filepath)
	if len(readUsers) == 0 {
		return false
	} else if len(readUsers) == 1 {
		if readUsers[0].Username == Username {
			userIndex = 0
		}
	} else {
		_, found := slices.BinarySearchFunc(readUsers, Username, func(a User, b string) int { return cmp.Compare(a.Username, b) })
		if !found {
			fmt.Printf("DeleteUser:%v does not exist\n", Username)
			return false
		}
		userIndex, found = slices.BinarySearchFunc(readUsers, Username, func(i User, j string) int {
			return cmp.Compare(i.Username, j)
		})

		if !found {
			return false
		}
	}
	readUsers = slices.Delete(readUsers, userIndex, userIndex)
	ConvertToJson(readUsers, filepath)
	return true
}

func CheckData(user User, filepath string) bool {
	var readUsers []User = GetJson(filepath)
	for _, value := range readUsers {
		if value.Username == user.Username {
			if value.Password == Encrypt(user.Password, 9) {
				return true
			}
		}
	}
	return false
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
