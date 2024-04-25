package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// initializeDB initializes the database and creates the required table if it does not exist.
func initializeDB() *sql.DB {
	db, err := sql.Open("sqlite3", "hideout.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	createTableSQL := `CREATE TABLE IF NOT EXISTS passwords (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        passwordName TEXT NOT NULL UNIQUE,
        passwordValue TEXT NOT NULL
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// insertPassword inserts a new password entry into the database.
func insertPassword(db *sql.DB, passwordName, passwordValue string) {
	insertSQL := `INSERT INTO passwords (passwordName, passwordValue) VALUES (?, ?)`
	_, err := db.Exec(insertSQL, passwordName, passwordValue)
	if err != nil {
		log.Fatal(err)
	}
}

// listPasswords lists all password entries from the database.
func listPasswords(db *sql.DB) {
	row, err := db.Query("SELECT id, passwordName, passwordValue FROM passwords")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	if !row.Next() {
		fmt.Println("No passwords stored.")
	}

	for row.Next() {
		var id int
		var passwordName, passwordValue string
		row.Scan(&id, &passwordName, &passwordValue)
		fmt.Printf("%d: %s -> %s\n", id, passwordName, passwordValue)
	}
}
