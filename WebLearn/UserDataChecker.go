package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var alphabet = [62]string{
	"A", "b", "9", "a", "r", "E", "5", "F", "G", "H", "I", "2", "J", "K", "L", "m",
	"W", "O", "P", "Q", "6", "R", "f", "T", "4", "U", "V", "N", "X", "7", "Y", "Z",
	"C", "B", "0", "c", "d", "e", "S", "g", "h", "i", "3", "j", "k", "l", "M",
	"n", "o", "p", "1", "v", "D", "s", "t", "u", "q", "w", "x", "y", "z", "8",
}

func getIndex(victim [62]string, perp string) uint8 {
	for index, value := range victim {
		if value == perp {
			return uint8(index)
		}
	}
	return 255
}

func Encrypt(SubPass string, offset uint8) string {
	var encrypted_pass string
	for _, value := range SubPass {
		var index uint8 = getIndex(alphabet, string(value))
		result := (index + offset)
		fmt.Printf("Result is: %v, of  %v \n", result, string(value))
		if result < uint8(len(alphabet)) {
			encrypted_pass += alphabet[result]
		} else {
			falloff := result - 62
			fmt.Printf("Fallof is %v \n", falloff)
			encrypted_pass += alphabet[falloff]
		}
	}
	return encrypted_pass
}

func Decrypt(Pass string, offset uint8) string {
	var decrypted_pass string
	for _, value := range Pass {
		var index uint8 = getIndex(alphabet, string(value))
		var result int8 = int8(index - offset)
		if result >= 0 {
			decrypted_pass += alphabet[result]
		} else {
			falloff := result + 62
			fmt.Printf("Decrypt->Fallof is %v \n", falloff)
			decrypted_pass += alphabet[falloff]
		}
	}
	fmt.Println("Decrypted to : ", decrypted_pass)
	return decrypted_pass
}

func CreateFile(SubName, SubPass string) {
	// Structure should be: 2 lines allocated per User
	//line:
	// 1.Username for User nr.1
	// 2.Password for User nr.1
	// 3.Username for User nr.2
	// 4.Password for User nr.2
	// etc.
	fmt.Println("Func CreateFile started")
	file, err := os.Create("DataFile.txt")
	if err != nil {
		fmt.Println("At: CreateFile, Error: ", err)
	}
	var UserData string = SubName + "\n" + Encrypt(SubPass, 9)
	_, err = file.WriteString(UserData)
	if err != nil {
		fmt.Println("WriteString, Error: ", err)
	}
	fmt.Println("CreateFile->UserData: ", UserData)
}

func ReadFile() (bool, map[string]string) {
	fmt.Println("Func ReadFile started")
	var line_number uint16 = 0
	var temp_keeper []string             // Will be used to keep the user data in between loops
	DataLines := make(map[string]string) // [[U,P],[U,P]]
	file, err := os.Open("DataFile.txt")
	if err != nil {
		fmt.Println("While opening DataFile.txt,Error: ", err)
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, make(map[string]string)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_number++
		if line_number%2 != 0 {
			// DataLines = append(DataLines, scanner.Text())
			temp_keeper = []string{}
			temp_keeper = append(temp_keeper, scanner.Text())
		} else {
			fmt.Println("Decrypt at line: ", line_number)
			fmt.Println("Keeper is: ", temp_keeper)
			temp_keeper = append(temp_keeper, Decrypt(scanner.Text(), 9))
			DataLines[temp_keeper[0]] = temp_keeper[1]

		}
	}
	fmt.Println(DataLines)
	return true, DataLines
}

func ValidateUserAndPassword(Data map[string]string, SubName, SubPass string) bool {
	User_Pass, ok := Data[SubName]
	if !ok {
		fmt.Printf("Validate->User: %v doesn't exist\n", SubName)
		return false
	}
	fmt.Println("Sub->", SubName, ":", SubPass, "Data->", User_Pass)
	return User_Pass == SubPass
}

func CheckData(SubName, SubPass string) bool {
	var file_exists bool
	var DataAsLines map[string]string
	for {
		file_exists, DataAsLines = ReadFile()
		if !file_exists {
			CreateFile(SubName, SubPass)
		}
		break
	}
	return ValidateUserAndPassword(DataAsLines, SubName, SubPass)
	//Create a file , come up with a storage format
	//Possibly Create Personal encryprion method
	//read file
	//Decrypt or Encrypt
	//Compare
	//return bool
	//
}
