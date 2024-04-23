package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var server = &http.Server{Addr: ":8080"}

type UserLogInData struct {
	Username string
	Password string
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	LogInFile := "./Templates/logIn.html"
	template, err := template.ParseFiles(LogInFile)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}
	err = template.ExecuteTemplate(w, "logIn.html", nil)
	if err != nil {
		log.Fatalf("Error executing template: %v ", err)
	}
}

func DefaultErrHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Error,,Default'' in GeneralHandler")
}

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	var LogInData UserLogInData
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")
	if LogInData.Username == "" {
		fmt.Println("Username is INFY")
	} else {
		fmt.Println("NOOOO")
	}
	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler: ", r.URL.Path)
	switch r.URL.Path {
	case "/end/":
		if err := server.Close(); err != nil {
			log.Fatal(err)
		}
	case "/LogIn":
		LogInHandler(w, r)
		// http.Handle("/css/", http.FileServer(http.Dir("./")))
	case "/Submit/LogIn":
		SubmitLogInHandler(w, r)
	default:
		DefaultErrHandler(w, r)
		// http.Handle(r.URL.Path, http.FileServer(http.Dir("./")))

	}
}

func main() {
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/", GeneralHandler)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
