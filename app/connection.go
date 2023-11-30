package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func connect() {
	var err error

	// Open the SQLite database file
	db, err = sql.Open("sqlite3", "GoData.db")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func closeConnection() {
	if db != nil {
		db.Close()
		fmt.Println("Connection closed")
	}
}
