package main

import (
	"cmp"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"slices"
)

var server = &http.Server{Addr: ":8080"}

type User struct {
	Username string
	Password string
	Subjects []string
	Schedule [][]string
}
type LogInError struct {
	Error string
}

var (
	LogInErr  LogInError
	LogInData User
)

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	LogInData.Username = r.FormValue("Username")
	LogInData.Password = r.FormValue("Password")

	fmt.Println("U: ", LogInData.Username)
	fmt.Println("P: ", LogInData.Password)
	var is_valid bool = CheckData(LogInData, "./Static/JsonData/Users.json")
	if is_valid {
		fmt.Println("User & Pass is valid")
		SetLoggedCookie(w, r, LogInData.Username)
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
	readUsers := GetJson("./Static/JsonData/Users.json")
	userIndex, found := slices.BinarySearchFunc(readUsers, LogInData, func(a, b User) int {
		return cmp.Compare(a.Username, b.Username)
	})
	if found {
		LogInErr.Error = "Use already exists!"
	} else {
		readUsers = slices.Insert(readUsers, userIndex, LogInData)
		readUsers[userIndex].Password = Encrypt(readUsers[userIndex].Password, 9)
		ConvertToJson(readUsers, "./Static/JsonData/Users.json")
	}
	// LogInErr.Error = AppendToArray("./Static/JsonData/Users.json", LogInData)
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

func SubmitAddSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubmitAddSubject called")
	r.ParseForm()
	subject := r.FormValue("Subject")
	fmt.Println("Subject: ", subject)

	readUsers := GetJson("./Static/JsonData/Users.json")
	userIndex, found := slices.BinarySearchFunc(readUsers, LogInData, func(i, j User) int {
		return cmp.Compare(i.Username, j.Username)
	})
	if !found {
		fmt.Println("SubmitAddSubject: ", LogInData.Username, " notFound")
		return
	}
	subjIndex, found := slices.BinarySearch(readUsers[userIndex].Subjects, subject)
	if !found {
		readUsers[userIndex].Subjects = slices.Insert(readUsers[userIndex].Subjects, subjIndex, subject)
		fmt.Println(LogInData.Username, "'s  subjects:", readUsers[userIndex].Subjects)
		ConvertToJson(readUsers, "./Static/JsonData/Users.json")
	}
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler URL.Path: ", r.URL.Path)
	if GetLoggedCookie(w, r) {
		fmt.Println("GeneralHandler GetLoggedCookie:true")
		switch r.URL.Path {
		case "/mainPage":
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/mainPage.html", "mainPage.html", LogInData)
		case "/LogIn":
			fmt.Println("LogIn Cookie Yes")
			http.Redirect(w, r, "/mainPage", http.StatusPermanentRedirect)
		// http.Handle("/css/", http.FileServer(http.Dir("./")))
		case "/Submit/LogIn":
			LogInErr.Error = ""
			SubmitLogInHandler(w, r)
		case "/Submit/AddSubject":
			SubmitAddSubject(w, r)
		case "/TicTacToe":
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/ticTacToe.html", "ticTacToe.html", nil)
		case "/Register":
			RegisterHandler(w, r)
		case "/DeleteMe":
			fmt.Println("Detele User")
			DeleteUser(LogInData.Username, "./Static/JsonData/Users.json")
			http.Redirect(w, r, "/LogOut", http.StatusFound)
		case "/Schedule":
			fmt.Println("/Schedule casse")
			readUsers := GetJson("./Static/JsonData/Users.json")

			userIndex, found := slices.BinarySearchFunc(readUsers, LogInData, func(a, b User) int {
				return cmp.Compare(a.Username, b.Username)
			})
			if found {
				HtmlHandler(w, r, "./Templates/schedule.html", "schedule.html", readUsers[userIndex])
			} else {
				fmt.Println("/Schedule: error user not found")
				HtmlHandler(w, r, "./Templates/schedule.html", "schedule.html", LogInData)
			}
		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
			// http.Redirect(w, r, "/LogIn", http.StatusFound)
		default:
			LogInErr.Error = ""
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)
			// http.Handle(r.URL.Path, http.FileServer(http.Dir("./")))

		}
	} else {
		switch r.URL.Path {
		case "/":
			http.Redirect(w, r, "/LogIn", http.StatusFound)
		case "/mainPage":
			LogInErr.Error = ""
			http.Redirect(w, r, "/LogIn", http.StatusFound)
		case "/LogIn":
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
			// http.Handle("/css/", http.FileServer(http.Dir("./")))
		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", LogInErr)
		case "/Submit/LogIn":
			LogInErr.Error = ""
			SubmitLogInHandler(w, r)
		case "/TicTacToe":
			LogInErr.Error = ""
			http.Redirect(w, r, "/LogIn", http.StatusNotModified)
		case "/Register":
			RegisterHandler(w, r)
		case "/Schedule":
			http.Redirect(w, r, "/LogIn", http.StatusNotModified)
		default:
			LogInErr.Error = ""
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)
			// http.Handle(r.URL.Path, http.FileServer(http.Dir("./")))

		}
	}
}

func main() {
	y := &secretKey
	var err error
	*y, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	fmt.Println("main secretKey:", secretKey, "   *y   ", *y)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/", GeneralHandler)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
