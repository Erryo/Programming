package main

//go:generate npm run build
import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	server = &http.Server{Addr: ":8000"}
	db     *sql.DB
)

type User struct {
	Username string
	Password string
	Subjects []string
	Schedule []lesson
}

type lesson struct {
	Day       string
	Name      string
	Lno       string
	StartTime string
	EndTime   string
	Id        string
}

func SubmitLogInHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: r.FormValue("Username"),
		Password: Encrypt(r.FormValue("Password"), 9),
		Subjects: []string{},
	}

	fmt.Println("U: ", user.Username)
	fmt.Println("P: ", user.Password)

	dbPass := getUser(db, user.Username)
	if dbPass == user.Password {
		fmt.Println("User & Pass is valid")
		SetLoggedCookie(w, r, user.Username)
		http.Redirect(w, r, "/mainPage", http.StatusMovedPermanently)
	} else if dbPass == "" {
		http.Redirect(w, r, "/LogIn?r=2", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/LogIn?r=1", http.StatusMovedPermanently)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")

	user := User{
		Username: r.FormValue("Username"),
		Password: Encrypt(r.FormValue("Password"), 9),
		Subjects: []string{},
	}
	fmt.Println("U: ", user.Username)
	fmt.Println("P: ", user.Password)

	if len(user.Username) > 25 || len(user.Password) > 40 {
		http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
		return
	}
	if !insertUser(db, user) {

		http.Redirect(w, r, "/LogIn?r=3", http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, "/LogIn?=4", http.StatusMovedPermanently)
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

func SubmitAddSubject(w http.ResponseWriter, r *http.Request, username string) {
	fmt.Println("SubmitAddSubject called")

	if r.Method == http.MethodDelete {
		subject := r.URL.Query().Get("S")
		deleteUserToSubj(db, username, subject)
		return
	}
	r.ParseForm()
	subject := r.FormValue("Subject")

	fmt.Println("Subject", subject)
	if len(subject) == 0 || subject == "" || subject == " " {
		return
	}
	err := insertUserToSubj(db, username, subject)
	if err {
		return
	}

	htmlStr := fmt.Sprintf(`
          <li class="text-xl mt-4 ml-6" id='%s'>
        %s
	<button 
		class="inline text-lg py-0 px-2 bg-white rounded-lg hover:animate-[size_1s_ease-in-out_infinite]"
		hx-delete="/Submit/AddSubject"
		hx-vals='{"S": "%s"}'
		hx-swap="delete"
		hx-target="#%s"
		type="button">
		-
	</button>
          </li>
        `, subject, subject, subject, subject)
	fmt.Println(htmlStr)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}

func AddScheduleHandler(w http.ResponseWriter, r *http.Request, username string) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("Id")
		deleteLesson(db, username, id)
		return
	}
	r.ParseForm()
	lesson := lesson{
		Day:       r.FormValue("Day"),
		Name:      r.FormValue("Name"),
		Lno:       r.FormValue("Lno"),
		StartTime: r.FormValue("StartTime"),
		EndTime:   r.FormValue("EndTime"),
	}
	fmt.Println(lesson)
	err := insertLesson(db, username, lesson)
	if err {
		return
	}
	htmlStr := fmt.Sprintf("<tr> <td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td>  <td>%s</td>  </tr>", lesson.Day, lesson.Name, lesson.Lno, lesson.StartTime, lesson.EndTime)
	htmlStr = fmt.Sprintf("<h1>AAAAA</h1>")
	fmt.Println(htmlStr)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeneralHandler URL.Path: ", r.URL.Path)
	username, found := GetLoggedCookie(w, r)
	if found {
		switch r.URL.Path {
		case "/mainPage":
			HtmlHandler(w, r, "./Templates/mainPage.html", "mainPage.html", username)

		case "/LogIn":
			fmt.Println("LogIn Cookie Yes")
			http.Redirect(w, r, "/mainPage", http.StatusPermanentRedirect)

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
			subjects := getUserSubj(db, username)
			schedule := getUserLessons(db, username)
			user := User{
				Username: username,
				Subjects: subjects,
				Schedule: schedule,
			}
			HtmlHandler(w, r, "./Templates/schedule.html", "schedule.html", user)
		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", nil)

		case "/Submit/LogIn":
			SubmitLogInHandler(w, r)

		case "/Submit/AddSubject":
			SubmitAddSubject(w, r, username)

		case "/Submit/AddSchedule":
			AddScheduleHandler(w, r, username)
		default:
			HtmlHandler(w, r, "./Templates/notFound.html", "notFound.html", nil)

		}
	} else {
		switch r.URL.Path {
		case "/LogIn":
			retry := r.URL.Query().Get("r")
			var message string
			switch retry {
			case "1":
				message = "Incorrect Password"
			case "2":
				message = "Incorrect Username"
			case "3":
				message = "User already Exists"
			case "4":
				message = "User created succesfully"
			}
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", message)
		case "/Submit/LogIn":
			SubmitLogInHandler(w, r)

		case "/Register":
			RegisterHandler(w, r)

		case "/LogOut":
			SetLoggedCookie(w, r, "nil")
			HtmlHandler(w, r, "./Templates/logIn.html", "logIn.html", nil)

		default:
			http.Redirect(w, r, "/LogIn", http.StatusMovedPermanently)
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
	createUserToSubjectTable(db)
	createLessonTable(db)
	createUserToLessonTable(db)

	defer db.Close()

	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.HandleFunc("/", GeneralHandler)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
