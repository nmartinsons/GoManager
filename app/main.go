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
	"strings"
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

	fmt.Print("Enter the title of the task to update: ")
	var oldTitle string
	fmt.Scanln(&oldTitle)

	oldTitle = strings.TrimSpace(oldTitle)

    var existingTask Task
    err := db.QueryRow("SELECT title, description, dueDate, priority, status FROM tasks WHERE title LIKE ?", "%"+oldTitle+"%").Scan(&existingTask.title, &existingTask.description, &existingTask.dueDate, &existingTask.priority, &existingTask.status)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Task not found")
        } else {
            log.Fatal(err)
        }
        return
    }


	fmt.Println("\nExisting Task Details:")
	fmt.Printf("Title: %sDescription: %sDue Date: %s\nPriority: %d\nStatus: %s\n", existingTask.title, existingTask.description, existingTask.dueDate, existingTask.priority, existingTask.status)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the new title of the task: ")
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)

	fmt.Print("Enter the new description of the task: ")
	newDescription, _ := reader.ReadString('\n')
	newDescription = strings.TrimSpace(newDescription)

	fmt.Print("Enter the new due date of the task (format: DD-MM-YYYY): ")
	newDueDateStr, _ := reader.ReadString('\n')
	newDueDate, err := time.Parse("02-01-2006", newDueDateStr[:len(newDueDateStr)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the new priority of the task: ")
	newPriorityStr, _ := reader.ReadString('\n')
	newPriority, err := strconv.Atoi(newPriorityStr[:len(newPriorityStr)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the new status of the task: ")
	newStatus, _ := reader.ReadString('\n')
	newStatus = strings.TrimSpace(newStatus)

	_, err = db.Exec("UPDATE tasks SET title = ?, description = ?, dueDate = ?, priority = ?, status = ? WHERE title = ?", newTitle, newDescription, newDueDate, newPriority, newStatus, oldTitle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nTask updated successfully!")
}


func deleteTask() {
	fmt.Print("Enter the title of the task to delete: ")
	var titleToDelete string
	fmt.Scanln(&titleToDelete)

	titleToDelete = strings.TrimSpace(titleToDelete)

	var existingTaskTitle Task
	err := db.QueryRow("SELECT title, description, dueDate, priority, status FROM tasks WHERE title LIKE ?", "%"+titleToDelete+"%").Scan(&existingTaskTitle.title, &existingTaskTitle.description, &existingTaskTitle.dueDate, &existingTaskTitle.priority, &existingTaskTitle.status)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Task not found")
        } else {
            log.Fatal(err)
        }
        return

	}

	_, err = db.Exec("DELETE FROM tasks WHERE title LIKE ?", titleToDelete)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nTask deleted successfully!")
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
