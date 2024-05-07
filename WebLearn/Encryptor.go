package main

import (
	"fmt"
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
		if index == 255 {
			fmt.Println("getIndex error")
		}
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
