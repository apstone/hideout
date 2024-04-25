package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initializeDB() *sql.DB {
	db, err := sql.Open("sqlite3", "hideout.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	createTableSQL := `CREATE TABLE IF NOT EXISTS greetings (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func insertGreeting(db *sql.DB, name string) {
	insertSQL := `INSERT INTO greetings (name) VALUES (?)`
	_, err := db.Exec(insertSQL, name)
	if err != nil {
		log.Fatal(err)
	}
}

func listGreetings(db *sql.DB) {
	row, err := db.Query("SELECT id, name FROM greetings")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		row.Scan(&id, &name)
		fmt.Printf("%d: Hello, %s!\n", id, name)
	}
}
