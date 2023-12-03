// 1. Create tasks
// 2. View tasks
// 3. Upadate tasks
// 4. Delete tasks
// 5. Displaying number of completed tasks

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	title       string
	description string
	dueDate     string
	priority    int
	status      string
}

func createTask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the title: ")
	title, _ := reader.ReadString('\n')

	fmt.Printf("Enter the description: ")
	description, _ := reader.ReadString('\n')

	fmt.Printf("Enter the due date (format: DD-MM-YYYY): ")
	dueDateStr, _ := reader.ReadString('\n')
	dueDate, err := time.Parse("02-01-2006", dueDateStr[:len(dueDateStr)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter the priority (format: 1): ")
	priorityStr, _ := reader.ReadString('\n')
	priority, err := strconv.Atoi(priorityStr[:len(priorityStr)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter the status (format: complete/incomplete): ")
	status, _ := reader.ReadString('\n')

	_, err = db.Exec("INSERT INTO tasks (title, description, dueDate, priority, status) VALUES (?, ?, ?, ?, ?)", title, description, dueDate, priority, status)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nTask created successfully!")
}


func viewTasks() {
    rows, err := db.Query("SELECT title, description, dueDate, priority, status FROM tasks")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("\nAll Tasks:")
    for rows.Next() {
        var task Task
        err := rows.Scan(&task.title, &task.description, &task.dueDate, &task.priority, &task.status)
        if err != nil {
            log.Fatal(err)
        }
		

        fmt.Printf("\nTitle: %sDescription: %sDue Date: %s\nPriority: %d\nStatus: %s\n", task.title, task.description, task.dueDate, task.priority, task.status)
    }
}


func updateTask() {
	
}



func deleteTask() {
	
}


func displayCompleteTasks() []Task {
	rows, err := db.Query("SELECT title, description, dueDate, priority, status FROM tasks WHERE status LIKE 'complete%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var completedTasks []Task

	fmt.Println("\nCompleted Tasks:")

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.title, &task.description, &task.dueDate, &task.priority, &task.status)
		if err != nil {
			log.Fatal(err)
		}
		completedTasks = append(completedTasks, task)
		fmt.Printf("\nTitle: %sDescription: %sDue Date: %s\nPriority: %d\nStatus: %s\n", task.title, task.description, task.dueDate, task.priority, task.status)
	}
	return completedTasks
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
