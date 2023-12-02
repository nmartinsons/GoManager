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
	"strconv"
	"strings"

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
	fmt.Println("\n> Create a Task")
	task := Task{}

	fmt.Print("Title: ")
	task.Title = readInput()

	fmt.Print("Description: ")
	task.Description = readInput()

	fmt.Print("Due Date (YYYYMMDD): ")
	dueDateStr := readInput()
	task.DueDate, _ = strconv.Atoi(dueDateStr)

	fmt.Print("Priority (1-5): ")
	priorityStr := readInput()
	task.Priority, _ = strconv.Atoi(priorityStr)

	fmt.Print("Status (completed/incomplete): ")
	task.Status = readInput()

	id, err := createTaskInDB(task)
	if err != nil {
		fmt.Println("Error creating task:", err)
	} else {
		fmt.Printf("Task created with ID %d\n", id)
	}
}

func createTaskInDB(task Task) (int64, error) {
	result, err := db.Exec(`
		INSERT INTO tasks (title, description, dueDate, priority, status)
		VALUES (?, ?, ?, ?, ?)
	`, task.Title, task.Description, task.DueDate, task.Priority, task.Status)

	if err != nil {
		log.Println("Error creating task:", err)
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func viewTasks() {
	tasks, err := getAllTasks()
	if err != nil {
		fmt.Println("Error viewing tasks: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found. Create a task first!")
		return
	}
	fmt.Println("\n> List of Tasks")
	for _, task := range tasks {
		fmt.Printf("Title: %s\nDescription: %s\nDue Date: %d\nPriority: %d\nStatus: %s\n\n",
			task.Title, task.Description, task.DueDate, task.Priority, task.Status)
	}
}

func getAllTasks() ([]Task, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println("Error getting tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
		if err != nil {
			log.Println("Error scanning task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func updateTask() {
	fmt.Print("\n> Enter Task ID to Update: ")
	idStr := readInput()
	id, _ := strconv.Atoi(idStr)

	updatedTask := Task{}

	fmt.Print("Updated Title: ")
	updatedTask.Title = readInput()

	fmt.Print("Updated Description: ")
	updatedTask.Description = readInput()

	fmt.Print("Updated Due Date (YYYYMMDD): ")
	dueDateStr := readInput()
	updatedTask.DueDate, _ = strconv.Atoi(dueDateStr)

	fmt.Print("Updated Priority (1-5): ")
	priorityStr := readInput()
	updatedTask.Priority, _ = strconv.Atoi(priorityStr)

	fmt.Print("Updated Status (completed/incomplete): ")
	updatedTask.Status = readInput()

	err := updateTaskInDB(id, updatedTask)
	if err != nil {
		fmt.Println("Error updating task:", err)
	} else {
		fmt.Printf("Task ID %d updated successfully.\n", id)
	}
}

func updateTaskInDB(id int, updatedTask Task) error {
	_, err := db.Exec(`
		UPDATE tasks
		SET title=?, description=?, dueDate=?, priority=?, status=?
		WHERE id=?
	`, updatedTask.Title, updatedTask.Description, updatedTask.DueDate, updatedTask.Priority, updatedTask.Status, id)

	if err != nil {
		log.Println("Error updating task:", err)
		return err
	}

	return nil
}

func deleteTask() {
	fmt.Print("\n> Enter Task ID to Delete: ")
	idStr := readInput()
	id, _ := strconv.Atoi(idStr)

	err := deleteTaskInDB(id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
	} else {
		fmt.Printf("Task ID %d deleted successfully.\n", id)
	}
}

func deleteTaskInDB(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		log.Println("Error deleting task:", err)
		return err
	}

	return nil
}

func displayCompleteTasks() {
	completedTasks, err := getCompletedTasks()
	if err != nil {
		fmt.Println("Error fetching completed tasks:", err)
		return
	}

	fmt.Printf("\nNumber of Completed Tasks: %d\n", len(completedTasks))
	fmt.Println("\n> List of Completed Tasks")
	for _, task := range completedTasks {
		fmt.Printf("Title: %s\nDescription: %s\nDue Date: %d\nPriority: %d\nStatus: %s\n\n",
			task.Title, task.Description, task.DueDate, task.Priority, task.Status)
	}
}

func getCompletedTasks() ([]Task, error) {
	rows, err := db.Query("SELECT * FROM tasks WHERE status='completed'")
	if err != nil {
		log.Println("Error fetching completed tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
		if err != nil {
			log.Println("Error scanning completed task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func readInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func main() {
	// Call the connect function to establish a connection
	connect()

	defer closeConnection()

	var choice int
	for {
		fmt.Println("\n> Select an operation:")
		fmt.Println("1. Create Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Display Completed Tasks")
		fmt.Println("6. Exit")
		fmt.Print("> Enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Println("Error reading choice:", err)
			continue
		}

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
			fmt.Println("Exiting Task Manager.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
