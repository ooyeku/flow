package main

import (
	"goworkflow/handle"
	"goworkflow/services"
	"goworkflow/store/mock"
	"log"
)

func main() {
	// Create a mock mockStore
	mockStore := mock.NewMockStore()

	// create a new taskMake service
	taskService := services.NewTaskService(mockStore)

	// Create a taskMake handler
	taskHandler := handle.NewTaskControl(taskService)

	// Create a new taskMake
	taskMake, err := taskHandler.CreateTask(&handle.CreateTaskRequest{
		Title:       "Task 1",
		Description: "Task 1 description",
		Owner:       "Task 1 owner",
	})

	if err != nil {
		log.Fatalf("Failed to create taskMake: %v", err)
	}

	log.Printf("Task created: %s", taskMake.ID)

	// Get the taskMake
	taskGet, err := taskHandler.GetTask(&handle.GetTaskRequest{ID: taskMake.ID})

	if err != nil {
		log.Fatalf("Failed to get taskMake: %v", err)
	}

	log.Printf("Task retrieved: %s", taskGet.Title)
	log.Printf("Description: %s", taskGet.Description)
	log.Printf("Owner: %s", taskGet.Owner)
	log.Printf("Started: %t", taskGet.Started)
	log.Printf("Completed: %t", taskGet.Completed)
	log.Printf("Created At: %s", taskGet.CreatedAt)
	log.Printf("Updated At: %s", taskGet.UpdatedAt)
}
