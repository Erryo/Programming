package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const IP = "http://localhost:8080"

type User struct {
	Username string
	Password string
	Subjects []string
	Schedule [][]string
}

func testLogIn() {
	user := url.Values{}
	user.Set("Username", "Infy")
	user.Set("Password", "Marius")
	r, err := http.PostForm(IP+"/LogIn", user)
	if err != nil {
		fmt.Println("Get failed")
		return
	}
	fmt.Println(r.Cookies())
}

func main() {
	for i := 0; i < 10; i++ {
		testLogIn()
	}
}
