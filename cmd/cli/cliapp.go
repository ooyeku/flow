package main

import (
	"bufio"
	"fmt"
	"github.com/asdine/storm"
	"goworkflow/internal/conf"
	"goworkflow/internal/inmemory"
	"goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to your CLI app!")
	for {
		fmt.Println("Enter command: ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
		err = runCommand(cmdString)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
	}
}

// cliSetup initializes the router, service, and in-memory store. It opens the database using the dbPath obtained from conf.GetDBPath(). If there is an error opening the database, it
func cliSetup() (*handle.TaskControl, *storm.DB) {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("error opening db: %s", err)
	}

	// Intialize router, service and inmemory store
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)

	return taskRouter, db
}

// createTask is a function that prompts the user to enter the details of a task,
// creates the task using the provided TaskControl instance, and prints the ID of the created task.
//
// Parameters:
// - t: *handle.TaskControl - An instance of TaskControl that provides the CreateTask method for task creation.
//
// Example usage:
//
//	taskRouter, db := cliSetup()
//	defer db.Close()
//	createTask(taskRouter)
//
// This function utilizes the TaskControl instance to create a task, by:
// 1. Printing a prompt for the task title and reading the user input.
// 2. Printing a prompt for the task description and reading the user input.
// 3. Printing a prompt for the task owner and reading the user input.
// 4. Creating a CreateTaskRequest instance with the title, description, and owner obtained from user input.
// 5. Invoking the CreateTask method of the TaskControl instance with the CreateTaskRequest instance.
// 6. If an error occurs during task creation, printing the error message.
// 7. Printing the ID of the created task.
func createTask(t *handle.TaskControl) {
	fmt.Println("Creating task...")
	fmt.Println("Task router in cliapp.go: ", t)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	title = strings.TrimSpace(title)
	fmt.Println("Enter task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	description = strings.TrimSpace(description)
	fmt.Println("Enter task owner: ")
	owner, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	owner = strings.TrimSpace(owner)
	req := handle.CreateTaskRequest{
		Title:       title,
		Description: description,
		Owner:       owner,
	}
	res, err := t.CreateTask(req)
	if err != nil {
		fmt.Println("Error creating task: ", err)
	}
	fmt.Println("Created task with id: ", res.ID)
}

// getTask retrieves a task from the task router based on the provided task ID.
// It prompts the user to input the task ID, sends a GetTaskRequest to the task router,
// and prints the task details if it exists.
func getTask(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	id = strings.TrimSpace(id)
	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

func getTaskByTitle(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	title = strings.TrimSpace(title)
	req := handle.GetTaskByTitleRequest{
		Title: title,
	}
	task, err := t.GetTaskByTitle(&req)
	if err != nil {
		// error message is logged in GetTaskByTitle
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

func getTaskByOwner(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task owner: ")
	owner, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	owner = strings.TrimSpace(owner)
	req := handle.GetTaskByOwnerRequest{
		Owner: owner,
	}
	task, err := t.GetTaskByOwner(&req)
	if err != nil {
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

// updateTasks prompts the user to enter a task ID, then retrieves the task with the specified ID using the GetTask function from the provided TaskControl object.
// It then prompts the user to enter a new task title, description, and owner.
// After that, it constructs an UpdateTaskRequest object with the provided ID, new title, description, owner, and default values for started and completed.
// Finally, it calls the UpdateTask function from the TaskControl object to update the task with the new values.
// If any error occurs during the process, it prints an error message.
// Example usage:
// t := &handle.TaskControl{service: taskService}
// updateTasks(t)
func updateTasks(t *handle.TaskControl) {
	reader := bufio.NewReader(os.Stdin)
	// Call your function to update a task here
	fmt.Println("Enter task id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	id = strings.TrimSpace(id)

	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
	fmt.Println("Enter New task title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	title = strings.TrimSpace(title)

	fmt.Println("Enter New task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	description = strings.TrimSpace(description)

	fmt.Println("Enter New task owner: ")
	owner, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	owner = strings.TrimSpace(owner)

	update := handle.UpdateTaskRequest{
		ID:          id,
		Title:       title,
		Description: description,
		Owner:       owner,
	}
	fmt.Println("Updating task...")
	err = t.UpdateTask(&update)
	if err != nil {
		fmt.Printf("Error updating task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Updated task with id: ", task.ID)
}

// deleteTask prompts the user to enter the ID of the task to be deleted. It retrieves the task with the given ID using the TaskControl.GetTask function.
// It then displays the details of the task to the user and asks for confirmation to delete the task. If the user confirms, it deletes the task using the TaskControl.DeleteTask function
func deleteTask(t *handle.TaskControl) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task id of task to be deleted: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	id = strings.TrimSpace(id)
	// Get task first to show user what task is being deleted
	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("are you sure you want to delete this task? (y/n)")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	confirm = strings.TrimSpace(confirm)
	if confirm == "n" {
		fmt.Println("Task not deleted")
	} else if confirm == "y" {
		fmt.Println("Deleting task...")
		req := handle.DeleteTaskRequest{
			ID: id,
		}
		err = t.DeleteTask(&req)
		if err != nil {
			fmt.Printf("Error deleting task with id %s: %s\n", id, err)
			return
		}
		fmt.Println("Deleted task with id: ", task.ID)
	} else {
		fmt.Println("Invalid input")
	}
}

// listTasks lists all tasks using the given TaskControl instance.
// It fetches all tasks using the ListTasks() method and prints the task ID, title, and description for each task.
// If there is an error in fetching the tasks, it prints the error message.
func listTasks(t *handle.TaskControl) {
	fmt.Println("Listing tasks...")
	tasks, err := t.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks: ", err)
		return
	}
	// get task id and title of each task
	for _, task := range tasks {
		fmt.Printf("Task id: %s, Title: %s, Description %s\n", task.ID, task.Title, task.Description)
	}
}

func runCommand(commandStr string) error {
	taskRouter, db := cliSetup()
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)

	case "create-task":
		createTask(taskRouter)

	case "get-task":
		getTask(taskRouter)

	case "get-task-by-title":
		getTaskByTitle(taskRouter)

	case "get-task-by-owner":
		getTaskByOwner(taskRouter)

	case "update-task":
		updateTasks(taskRouter)

	case "delete-task":
		deleteTask(taskRouter)

	case "list-tasks":
		listTasks(taskRouter)
	default:
		return fmt.Errorf("unknown command: %s", commandStr)
	}

	return nil
}
