package main

import (
	"encoding/json"
	"fmt"
	"github.com/asdine/storm"
	"goworkflow/handle"
	"goworkflow/services"
	"goworkflow/store/inmemory"
	"log"
)

func main() {

	db, err := storm.Open("store/inmemory/goworkflow.db", storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	inMem := inmemory.NewInMemoryTaskStore(db)

	service := services.NewTaskService(inMem)

	router := handle.NewTaskControl(service)

	// get all tasks in inmemory db
	tasks, err := router.ListTasks()
	if err != nil {
		log.Fatalf("Failed to list tasks: %v", err)
	}

	// print all tasks
	taskJson, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Failed to marshal tasks: %v", err)
	}

	fmt.Println(string(taskJson))
}
