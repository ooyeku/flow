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

	fmt.Println("Routers initialized in cliapp.go: ", taskRouter)
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
		// Call your function to get a task here
		fmt.Println("Getting task...")
	case "update-task":
		// Call your function to update a task here
		fmt.Println("Updating task...")
	case "delete-task":
		// Call your function to delete a task here
		fmt.Println("Deleting task...")
	case "list-tasks":
		// Call your function to list all tasks here
		fmt.Println("Listing tasks...")
	default:
		return fmt.Errorf("unknown command: %s", commandStr)
	}

	return nil
}
