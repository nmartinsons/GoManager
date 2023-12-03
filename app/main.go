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
	"os"

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

	// print success message if connected
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

func viewTasks() {
	
}

func updateTask() {
}



func deleteTask() {
	
}


func displayCompleteTasks() {
	
}


func main() {
	// connect function is being called to establish a connection
	connect()

	defer closeConnection()

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Create task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Update task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Display completed tasks")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createTask()
		case 2:
			viewTasks()
		case 3:
			updateTask()
		case 4:
			deleteTask()
		case 5:
			displayCompleteTasks()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
