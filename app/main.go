// 1. Create tasks
// 2. View tasks
// 3. Upadate tasks
// 4. Delete tasks
// 5. Displaying number of completed tasks

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

type Task struct {
	Title       string
	Description string
	DueDate     int
	Priority    int
	Status      string
}

func createTask() {
	
}

func updateTask() {
}



func deleteTask() {
	
}


func displayCompleteTasks() {
	
}


func main() {
	// Call the connect function to establish a connection
	connect()

	defer closeConnection()

}
