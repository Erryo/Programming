package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func binarySearch(a []User, search string) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid].Username > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid].Username < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}

func sortSlice(list []User, filepath string) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Username < list[j].Username
	})
	ConvertToJson(list, filepath)
}

func DeleteUser(Username string, filepath string) bool {
	fmt.Println("DeleteUser: ", Username, " started")
	readUsers := GetJson("./Static/JsonData/Users.json")
	if !UserExists(readUsers, Username) {
		fmt.Printf("DeleteUser:%v does not exist\n", Username)
		return false
	}
	// sortSlice(readUsers, "./Static/JsonData/Users.json")
	userIndex, _ := binarySearch(readUsers, Username)
	newSlice := make([]User, 0)
	newSlice = append(newSlice, readUsers[:userIndex]...)
	newSlice = append(newSlice, readUsers[userIndex+1:]...)
	ConvertToJson(newSlice, filepath)
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

func UserExists(allUsers []User, Username string) bool {
	for _, value := range allUsers {
		if value.Username == Username {
			return true
		}
	}
	return false
}

func AppendToArray(filepath string, user User) string {
	var readUsers []User = GetJson(filepath)
	if !UserExists(readUsers, user.Username) {
		user.Password = Encrypt(user.Password, 9)
		readUsers = append(readUsers, user)
		sortSlice(readUsers, filepath)

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
