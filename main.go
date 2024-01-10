package main

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	"goworkflow/handle"
	"goworkflow/services"
	"goworkflow/store"
	"goworkflow/store/inmemory"
	"log"
	"os"
)

// Config  represents the configuration for the application.
type Config struct {
	DatabaseType string `json:"database_type"`
}

// decodeConfig decodes the JSON data in the specified file and updates the Config object with the result.
func (c *Config) decodeConfig(file *os.File) error {
	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}

// createInMemoryDB creates a new in-memory database using the storm package.
// It returns a pointer to the created storm.DB object and an error if any occurred.
// The database is created at the specified path: "store/inmemory/goworkflow.db".
// The database is opened with bolt options 0600 (read-write mode) and no specific options.
func createInMemoryDB() (*storm.DB, error) {
	db, err := storm.Open("store/inmemory/goworkflow.db", storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// selectDatabase selects the appropriate task store based on the value of the DatabaseType field in the Config struct.
func (c *Config) selectDatabase() (store.TaskStore, error) {
	switch c.DatabaseType {
	case "bolt":
		db, err := createInMemoryDB()
		if err != nil {
			return nil, err
		}
		return inmemory.NewInMemoryTaskStore(db), nil
		// Other cases...
	}
	return nil, errors.New("invalid database type specified in config")
}

// initializeConfiguration initializes the configuration by decoding the configuration file at the given file path.
// It returns a pointer to the Config struct and an error, if any.
func initializeConfiguration(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = config.decodeConfig(file)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// marshalTasks is a function that takes a pointer to a handle.ListTasksResponse struct and returns the JSON representation of the struct as a string.
//
// The function first uses the json.Marshal() function to convert the tasks struct to its JSON representation. If there is an error during marshaling, it logs a fatal error.
//
// Finally, it returns the JSON string representation of the tasks struct.
//
// Example usage:
//
//	tasks, err := router.ListTasks()
//	if err != nil {
//		log.Fatalf("Failed to list tasks: %v", err)
//	}
//	taskJson := marshalTasks(tasks)
//	log.Printf("Tasks: %s", taskJson)
func marshalTasks(tasks *handle.ListTasksResponse) string {
	taskJson, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Failed to marshal tasks: %v", err)
	}
	return string(taskJson)
}

// initializeRouter initializes the router and returns a pointer to the TaskControl object.
// It first initializes the configuration by calling the initializeConfiguration function with "config.json" as the file path.
// If an error occurs during initialization, it logs the error and exits.
// Then, it selects the database based on the database type specified in the configuration.
// If the database type is "bolt", it creates an in-memory database by calling the createInMemoryDB function.
// If an error occurs during database selection, it logs the error and exits.
// Next, it creates a new TaskService object with the selected taskStore.
// After that, it creates a new TaskControl object with the taskService and assigns it to the router variable.
// Finally, it logs the success message and returns the router.
// Example usage:
//
//	router := initializeRouter()
//	tasks, err := router.ListTasks()
//	if err != nil {
//	    log.Fatalf("Failed to list tasks: %v", err)
//	}
//	taskJson := marshalTasks(tasks)
//	log.Printf("Tasks: %s", taskJson)
func initializeRouter() *handle.TaskControl {
	config, err := initializeConfiguration("config.json")
	if err != nil {
		log.Fatalf("failed to initialize configuration: %v", err)
	}
	log.Print("Configuration initialized -- success")

	taskStore, err := config.selectDatabase()
	if err != nil {
		log.Fatalf("failed to select database: %v", err)
	}
	log.Print("Database initialized -- success")

	taskService := services.NewTaskService(taskStore)
	router := handle.NewTaskControl(taskService)
	log.Print("Router initialized -- success")

	return router
}

func getAllTasks(r *handle.TaskControl) string {
	// get all tasks db, pass initialized router
	tasks, err := r.ListTasks()
	if err != nil {
		log.Fatalf("Failed to list tasks: %v", err)
	}

	taskJson := marshalTasks(tasks)
	log.Printf("Tasks: %s", taskJson)

	return taskJson
}

func main() {
	router := initializeRouter()
	getAllTasks(router)
}
