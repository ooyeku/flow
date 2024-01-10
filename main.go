package main

import (
	"github.com/asdine/storm"
	"goworkflow/handle"
	"goworkflow/services"
	//"goworkflow/store/mock"
	"goworkflow/store/inmemory"
	"log"
)

func main() {
	// Create a mock mockStore
	//mockStore := mock.NewMockStore()

	db, err := storm.Open("goworkflow.db", storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	// Create a new in-memory store
	inMemoryStore := inmemory.NewInMemoryTaskStore(db)

	// create a new taskMake service
	taskService := services.NewTaskService(inMemoryStore)

	// Create a taskMake handler
	taskHandler := handle.NewTaskControl(taskService)

	log.Printf("Task handler: %v", taskHandler)

	// Create a new taskMake
	taskMake, err := taskHandler.CreateTask(handle.CreateTaskRequest{
		Title:       "Task 1",
		Description: "Task 1 description",
		Owner:       "Me",
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
