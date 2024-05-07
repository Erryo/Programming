package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var server = &http.Server{Addr: ":8080"}

type User struct {
	Username string
	Password string
}
type LogInError struct {
	Error string
}

var (
	LogInErr  LogInError
	LogInData User
)

func SetLoggedCookie(w http.ResponseWriter, r *http.Request, isLogged string) {
	cookie := &http.Cookie{
		Name:  "UserStatus",
		Value: isLogged,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func GetLoggedCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("UserStatus")
	if err != nil {
		fmt.Println("cookie not found")
		return false
	}
	if cookie.Value == "true" {
		fmt.Println("GetLoggedCookie->cookie true")
		return true
	}
	return false
}

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")

	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
	var is_valid bool = CheckData(LogInData, "./Static/JsonData/Users.json")
	if is_valid {
		fmt.Println("User & Pass is valid")
		SetLoggedCookie(w, r, "true")
		http.Redirect(w, r, "/mainPage", http.StatusMovedPermanently)
		// MainPageHandler(w, r)
	} else {
		LogInErr.Error = "User or password incorrect/does not exist"
		http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")

	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
	LogInErr.Error = AppendToArray("./Static/JsonData/Users.json", LogInData)
	http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
}

func HtmlHandler(w http.ResponseWriter, r *http.Request, HtmlFile string, HtmlName string, Data interface{}) {
	template, err := template.ParseFiles(HtmlFile)
	if err != nil {
		fmt.Printf("HtmlHandler:%v ->Error template.parsefiles: %v \n", HtmlName, err)
	}
	err = template.ExecuteTemplate(w, HtmlName, Data)
	if err != nil {
		fmt.Printf("HtmlHandler:%v ->error executing template: %v \n", HtmlName, err)
	}
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler: ", r.URL.Path)
	if GetLoggedCookie(w, r) {
		switch r.URL.Path {
		case "/mainPage":
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/mainPage.html", "mainPage.html", LogInData)
		case "/LogIn":
			fmt.Println("LogIn Cookie Yes")
			http.Redirect(w, r, "/mainPage", http.StatusMovedPermanently)
			// http.Handle("/css/", http.FileServer(http.Dir("./")))
		case "/Submit/LogIn":
			LogInErr.Error = ""
			SubmitLogInHandler(w, r)
		case "/TicTacToe":
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/ticTacToe.html", "ticTacToe.html", nil)
		case "/Register":
			RegisterHandler(w, r)
		case "/Schedule":
			HtmlHandler(w, r, "./Templates/schedule.html", "schedule.html", nil)
		case "/LogOut":
			SetLoggedCookie(w, r, "false")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
			// http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		default:
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)
			// http.Handle(r.URL.Path, http.FileServer(http.Dir("./")))

		}
	} else {
		switch r.URL.Path {
		case "/mainPage":
			LogInErr.Error = ""
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		case "/LogIn":
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
			// http.Handle("/css/", http.FileServer(http.Dir("./")))
		case "/LogOut":
			SetLoggedCookie(w, r, "false")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
		case "/Submit/LogIn":
			LogInErr.Error = ""
			SubmitLogInHandler(w, r)
		case "/TicTacToe":
			LogInErr.Error = ""
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		case "/Register":
			RegisterHandler(w, r)
		case "/Schedule":
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		default:
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)
			// http.Handle(r.URL.Path, http.FileServer(http.Dir("./")))

		}
	}
}

func main() {
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/", GeneralHandler)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
