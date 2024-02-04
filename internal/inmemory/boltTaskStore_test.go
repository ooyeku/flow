package inmemory

import (
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/pkg/handle"
	"github.com/ooyeku/flow/pkg/services"
	"log"
	"testing"
)

func TestTaskStore(t *testing.T) {

	db, err := storm.Open("test.db", storm.BoltOptions(0600, nil))
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
	inMemoryStore := NewInMemoryTaskStore(db)

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
}
