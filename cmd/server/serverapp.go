package main

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/gorilla/mux"
	"goworkflow/api"
	"goworkflow/internal/conf"
	"goworkflow/internal/inmemory"
	"goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	"log"
	"net/http"
	"time"
)

func cliSetup() (*handle.TaskControl, *storm.DB, error) {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, nil, fmt.Errorf("error opening db: %s", err)
	}
	// Intialize router, service and inmemory store
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)
	return taskRouter, db, nil
}

// loggingMiddleware logs the request method, URL, and the time it took to process the request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()
	taskRouter, db, err := cliSetup()
	if err != nil {
		log.Fatalf("error setting up cli: %s", err)
	}

	defer func(db *storm.DB) {
		_ = db.Close()
	}(db)

	// Initialize TaskHandler
	taskHandler := &api.TaskHandler{
		Control: taskRouter,
	}
	// Register handlers
	r.HandleFunc("/list", taskHandler.ListTasks).Methods("GET")
	r.HandleFunc("/task/new", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/task/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// Apply the middleware to the router
	r.Use(loggingMiddleware)

	log.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("error serving: %s", err)
	}
}
