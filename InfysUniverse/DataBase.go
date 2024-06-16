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
	 * Subjects
	 * Schedule
	 */
	query := `
    CREATE TABLE IF NOT EXISTS IUuser (
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL, 
    PRIMARY KEY(username)
    )
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, user User) bool {
	query := `
    INSERT INTO IUuser(username,password)
    VALUES ($1,$2)`
	err := db.QueryRow(query, user.Username, user.Password).Scan()
	if err != nil {
		return false
		log.Print(err)
	}
	return true
}

func getUser(db *sql.DB, username string) string {
	var password string

	query := `
    SELECT password FROM IUuser WHERE username = $1
    `
	err := db.QueryRow(query, username).Scan(&password)
	if err != nil {
		log.Print(err)
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
