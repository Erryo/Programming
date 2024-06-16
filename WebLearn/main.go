package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	server = &http.Server{Addr: ":8080"}
	db     *sql.DB
)

type User struct {
	Username string
	Password string
	Subjects []string
	Schedule [][]string
}

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}

	fmt.Println("U: ", user.Username)
	fmt.Println("P: ", user.Password)

	dbPass := getUser(db, user.Username)
	if dbPass == Encrypt(user.Password, 9) {
		fmt.Println("User & Pass is valid")
		SetLoggedCookie(w, r, user.Username)
		http.Redirect(w, r, "/mainPage", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")

	user := User{
		Username: r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}

	fmt.Println("U: ", user.Username)
	fmt.Println("P: ", user.Password)
	user.Password = Encrypt(user.Username, 9)
	if len(user.Username) > 255 || len(user.Password) > 255 {
		http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		return
	}
	insertUser(db, user)
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
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler URL.Path: ", r.URL.Path)
	username, found := GetLoggedCookie(w, r)
	if found {
		switch r.URL.Path {
		case "/mainPage":
			HtmlHandler(w, r, "./Templates/mainPage.html", "mainPage.html", nil)

		case "/LogIn":
			fmt.Println("LogIn Cookie Yes")
			http.Redirect(w, r, "/mainPage", http.StatusPermanentRedirect)

		case "/Submit/LogIn":
			SubmitLogInHandler(w, r)

		case "/Submit/AddSubject":
			SubmitAddSubject(w, r)

		case "/TicTacToe":
			HtmlHandler(w, r, "./Templates/ticTacToe.html", "ticTacToe.html", nil)

		case "/Register":
			RegisterHandler(w, r)

		case "/DeleteMe":
			fmt.Println("Detele User")
			deleteUser(db, username)
			http.Redirect(w, r, "/LogOut", http.StatusFound)

		case "/Schedule":
			fmt.Println("/Schedule casse")
		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", nil)
		default:
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)

		}
	} else {
		switch r.URL.Path {
		case "/":
			http.Redirect(w, r, "/LogIn", http.StatusFound)
		case "/mainPage":
			http.Redirect(w, r, "/LogIn", http.StatusFound)
		case "/LogIn":
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", nil)
		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", nil)
		case "/Submit/LogIn":
			SubmitLogInHandler(w, r)
		case "/TicTacToe":
			http.Redirect(w, r, "/LogIn", http.StatusNotModified)
		case "/Register":
			RegisterHandler(w, r)
		case "/Schedule":
			http.Redirect(w, r, "/LogIn", http.StatusNotModified)
		default:
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)
		}
	}
}

func main() {
	y := &secretKey
	var err error
	*y, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}

	db = connectDB()
	createUserTable(db)
	defer db.Close()

	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/", GeneralHandler)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
