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
type LogInError struct {
	Error string
}

var (
	LogInErr  LogInError
	LogInData UserLogInData
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	LogInFile := "./Templates/logIn.html"
	template, err := template.ParseFiles(LogInFile)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}
	err = template.ExecuteTemplate(w, "logIn.html", LogInErr)
	if err != nil {
		log.Fatalf("Error executing template: %v ", err)
	}
}

func DefaultErrHandler(w http.ResponseWriter, r *http.Request) {
	notFoundHtml := "./Templates/notFound.html"
	template, err := template.ParseFiles(notFoundHtml)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}
	err = template.ExecuteTemplate(w, "notFound.html", nil)
	if err != nil {
		log.Fatalf("Error executing template: %v ", err)
	}
	fmt.Println("Error,,Default'' in GeneralHandler")
}

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")

	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
	var is_valid bool = CheckData(LogInData.Username, LogInData.Password)
	if is_valid {
		fmt.Println("User & Pass is valid")
		http.Redirect(w, r, "/mainPage", http.StatusMovedPermanently)
		// MainPageHandler(w, r)
	} else {
		LogInErr.Error = "User or password incorrect/does not exist"
		http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
	}
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	mainPageFile := "./Templates/mainPage.html"
	template, err := template.ParseFiles(mainPageFile)
	if err != nil {
		fmt.Println("MainPageHandler->Error template.parsefiles ")
	}
	err = template.ExecuteTemplate(w, "mainPage.html", LogInData)
	if err != nil {
		fmt.Println("MainPageHandler->error executing template ")
	}
}

func TicTacToeHandler(w http.ResponseWriter, r *http.Request) {
	ticFile := "./Templates/ticTacToe.html"
	template, err := template.ParseFiles(ticFile)
	if err != nil {
		fmt.Println("TicTacToeHandler->Error template.parsefiles ")
	}
	err = template.ExecuteTemplate(w, "ticTacToe.html", LogInData)
	if err != nil {
		fmt.Println("TicTacToeHandler->error executing template ")
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")

	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
	LogInErr.Error = CreateUser(LogInData.Username, LogInData.Password)
	http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler: ", r.URL.Path)
	switch r.URL.Path {
	case "/mainPage":
		LogInErr.Error = ""
		MainPageHandler(w, r)
	case "/LogIn":
		LogInHandler(w, r)
		// http.Handle("/css/", http.FileServer(http.Dir("./")))
	case "/Submit/LogIn":
		LogInErr.Error = ""
		SubmitLogInHandler(w, r)
	case "/TicTacToe":
		LogInErr.Error = ""
		TicTacToeHandler(w, r)
	case "/Register":
		RegisterHandler(w, r)
	default:
		LogInErr.Error = ""
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
