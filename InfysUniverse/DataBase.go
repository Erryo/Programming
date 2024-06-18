package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr string = "postgres://postgres:marius@localhost:5432/InfyUniverse?sslmode=disable"

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db connected succesfully")
	return db
}

func createUserTable(db *sql.DB) {
	/*
	 * Username -- identifier
	 * Password
	 */
	query := `
    CREATE TABLE IF NOT EXISTS IUuser (
    username varchar(25) NOT NULL,
    password varchar(40) NOT NULL, 
    PRIMARY KEY(username)
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("User table ", err)
	}
}

func createLessonTable(db *sql.DB) {
	/*
	* ID -- identifier
	* Day
	* Name
	* Start
	* End
	* Lesson No
	 */
	query := `CREATE TABLE IF NOT EXISTS lessons(
        id SMALLSERIAL NOT NULL PRIMARY KEY,
        day VARCHAR(15) NOT NULL,
        startTime TIME,
        endTime TIME,
        lessonNo  SMALLINT,
        name varchar(25)
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("LessTable ", err)
	}
}

func createUserToSubjectTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS UserToSubject(
    username varchar(25) NOT NULL,
    name varchar(25) NOT NULL,
    PRIMARY KEY(username,name),
    FOREIGN KEY(username) REFERENCES IUuser(username)
    )`
	_, err := db.Query(query)
	if err != nil {
		log.Fatal("Create UtoS ", err)
	}
}

func createUserToLessonTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS UserToLesson(
    username varchar(25) NOT NULL,
    lid SMALLINT NOT NULL,
    PRIMARY KEY(username,lid),
    FOREIGN KEY(username) REFERENCES IUuser(username),
    FOREIGN KEY(lid) REFERENCES lessons(id)
    )`
	_, err := db.Query(query)
	if err != nil {
		log.Fatal("UtoL ", err)
	}
}

func insertUser(db *sql.DB, user User) bool {
	query := `
    INSERT INTO IUuser(username,password)
    VALUES ($1,$2)`
	err := db.QueryRow(query, user.Username, user.Password).Scan()
	if err != nil {
		log.Print("Ins user: ", err)
		return false
	}
	return true
}

func getUser(db *sql.DB, username string) string {
	if username == "" {
		return ""
	}

	var password string
	query := `
    SELECT password FROM IUuser WHERE username = $1 
    `
	err := db.QueryRow(query, username).Scan(&password)
	if err != nil {
		log.Print("getUser ", err)
	}
	return password
}

func deleteUser(db *sql.DB, username string) {
	query := `
    DELETE FROM IUuser WHERE username = $1
    `
	_, err := db.Query(query, username)
	if err != nil {
		log.Print(err)
	}
}

func insertUserToSubj(db *sql.DB, username, subject string) {
	query := `
    INSERT INTO UserToSubject (username,name)
    VALUES ($1,$2)
    `
	_, err := db.Query(query, username, subject)
	if err != nil {
		log.Println(err)
	}
}

func getUserSubj(db *sql.DB, username string) []string {
	var subjects []string

	query := `
    SELECT name FROM UserToSubject WHERE username = $1
    `
	rows, err := db.Query(query, username)
	if err != nil {
		log.Print(err)
	}

	defer rows.Close()

	for rows.Next() {
		var subject string
		err := rows.Scan(&subject)
		if err != nil {
			log.Fatal(err)
		}
		subjects = append(subjects, subject)
	}
	return subjects
}
