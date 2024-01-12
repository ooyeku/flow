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
		fmt.Print("Enter command: ")
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

func runCommand(commandStr string) error {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("error opening db: %s", err)
	}
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)

	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	// add case clauses for the other commands in your application, such as:
	case "create-task":
		// Call your function to create a task here
		fmt.Println("Creating task...")
		fmt.Println("Task router in cliapp.go: ", taskRouter)

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter task title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Println("Enter task description: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		fmt.Println("Enter task owner: ")
		owner, _ := reader.ReadString('\n')
		owner = strings.TrimSpace(owner)

		req := handle.CreateTaskRequest{
			Title:       title,
			Description: description,
			Owner:       owner,
		}

		res, err := taskRouter.CreateTask(req)
		if err != nil {
			fmt.Println("Error creating task: ", err)
			return err
		}
		fmt.Println("Created task with id: ", res.ID)
	case "get-task":
		reader := bufio.NewReader(os.Stdin)
		// Call your function to get a task here
		fmt.Println("Enter task id: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		req := handle.GetTaskRequest{
			ID: id,
		}
		task, err := taskRouter.GetTask(&req)
		if err != nil {
			fmt.Printf("Error getting task with id %s: %s", id, err)
		}
		// unmarschal task to json
		fmt.Println("Task: ", task)

	case "update-task":
		reader := bufio.NewReader(os.Stdin)
		// Call your function to update a task here
		fmt.Println("Enter task id: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		req := handle.GetTaskRequest{
			ID: id,
		}
		task, err := taskRouter.GetTask(&req)
		if err != nil {
			fmt.Printf("Error getting task with id %s: %s", id, err)
		}
		fmt.Println("Got task: ", task.Title)
		fmt.Println("Description: ", task.Description)
		fmt.Println("Enter New task title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Println("Enter New task description: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		fmt.Println("Enter New task owner: ")
		owner, _ := reader.ReadString('\n')
		owner = strings.TrimSpace(owner)

		update := handle.UpdateTaskRequest{
			ID:          id,
			Title:       title,
			Description: description,
			Owner:       owner,
		}
		fmt.Println("Updating task...")
		err = taskRouter.UpdateTask(&update)
		if err != nil {
			fmt.Printf("Error updating task with id %s: %s", id, err)
		}
		fmt.Println("Updated task with id: ", task.ID)
	case "delete-task":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter task id of task to be deleted: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		// Get task first to show user what task is being deleted
		req := handle.GetTaskRequest{
			ID: id,
		}
		task, err := taskRouter.GetTask(&req)
		if err != nil {
			fmt.Printf("Error getting task with id %s: %s", id, err)
		}
		fmt.Println("Got task: ", task.Title)
		fmt.Println("are you sure you want to delete this task? (y/n)")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)
		if confirm == "n" {
			fmt.Println("Task not deleted")
			return nil
		} else if confirm == "y" {
			fmt.Println("Deleting task...")
			req := handle.DeleteTaskRequest{
				ID: id,
			}
			err = taskRouter.DeleteTask(&req)
			if err != nil {
				fmt.Printf("Error deleting task with id %s: %s", id, err)
			}
			fmt.Println("Deleted task with id: ", task.ID)
		} else {
			fmt.Println("Invalid input")
			return nil
		}
	case "list-tasks":
		// Call your function to list all tasks here
		fmt.Println("Listing tasks...")
		tasks, err := taskRouter.ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks: ", err)
		}
		// get task id and title of each task
		for _, task := range tasks {
			fmt.Printf("Task id: %s, Title: %s, Description %s\n", task.ID, task.Title, task.Description)
		}
	default:
		return fmt.Errorf("unknown command: %s", commandStr)
	}

	return nil
}
