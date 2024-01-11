package conf

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

// ConfigFileName defines the name of the configuration file.
const (
	ConfigFileName   = "config.json"
	InMemoryDatabase = "store/inmemory/goworkflow.db"
)

// LogAndExitOnError logs an error message and exits the program if the error is not nil.
// It uses the log.Fatalf function to log the error message with the error value formatted in it.
// The function takes an error and a message string as parameters.
// If the error is not nil, it will log the message with the error value and exit the program.
//
// Example usage:
//
//	err := someFunction()
//	LogAndExitOnError(err, "Failed to do something: %v")
//
// Any information on the containing package, return example code, or tags are omitted for brevity.
func LogAndExitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

// Config represents the configuration for the application.
type Config struct {
	DatabaseType string `json:"database_type"`
}

// decodeConfig decodes the contents of the file given as an argument
// using the json.NewDecoder function, and then decodes it into the Config
// object c using the Decode method. It returns an error if the decoding fails.
func (c *Config) decodeConfig(file *os.File) error {
	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}

// createInMemoryDB creates and returns a new in-memory storm database.
// It opens the database with the specified InMemoryDatabase constant and Bolt options.
// Returns the database object and any error encountered during the operation.
func createInMemoryDB() (*storm.DB, error) {
	db, err := storm.Open(InMemoryDatabase, storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// selectTaskDatabase selects the appropriate task database based on the configured database type in the Config struct.
// It returns an instance of type TaskStore, which implements the TaskStore interface, and an error if the database type is invalid.
// The method takes a storm.DB as a parameter, which is the database used for persistence.
// In case the database type is "bolt", it creates and returns an instance of an in-memory task store using the storm.DB passed as an argument.
// If the database type is invalid, it returns an error indicating the invalidity of the database type specified in the Config.
func (c *Config) selectTaskDatabase(db *storm.DB) (store.TaskStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryTaskStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

// selectPlannerDatabase selects the appropriate planner database based on the specified DatabaseType in the Config object.
// If the DatabaseType is "bolt", it returns a new instance of an InMemoryPlannerStore using the provided db object.
// If the DatabaseType is invalid or not supported, it returns an error indicating the invalid database type.
// The returned PlannerStore interface can be used to perform various operations on the planner data.
func (c *Config) selectPlannerDatabase(db *storm.DB) (store.PlannerStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryPlannerStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

// initializeConfiguration initializes the configuration by reading the configuration file and returning a Config object.
//
// It takes a `filePath` parameter which specifies the path of the configuration file.
// It opens the file, decodes the contents into a Config object, and returns the Config object and any error encountered.
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

// initializeRouter initializes the task and planner routers by creating the necessary dependencies and services. It returns the task and planner router instances.
//
// The `initializeRouter` function performs the following steps:
// 1. It initializes the configuration by reading from the specified file path. If there is an error, it logs the error message and exits.
// 2. It creates an in-memory database. If there is an error, it logs the error message and exits.
// 3. It selects the task database based on the configuration. If there is an error, it logs the error message and exits.
// 4. It creates a new task service using the selected task store.
// 5. It creates a new task router using the task service.
// 6. It selects the planner database based on the configuration. If there is an error, it logs the error message and exits.
// 7. It creates a new planner service using the selected planner store.
// 8. It creates a new planner router using the planner service.
// 9. It logs a success message.
// 10. It returns the task and planner router instances.
//
// Example usage:
//
//	taskRouter, plannerRouter := initializeRouter()
func initializeRouter() (*handle.TaskControl, *handle.PlannerControl) {
	config, err := initializeConfiguration(ConfigFileName)
	LogAndExitOnError(err, "failed to initialize configuration: %v")

	db, err := createInMemoryDB()
	LogAndExitOnError(err, "failed to create in-memory database: %v")

	taskStore, err := config.selectTaskDatabase(db)
	LogAndExitOnError(err, "failed to select task database: %v")
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)

	plannerStore, err := config.selectPlannerDatabase(db)
	LogAndExitOnError(err, "failed to select planner database: %v")
	plannerService := services.NewPlannerService(plannerStore)
	plannerRouter := handle.NewPlannerControl(plannerService)

	log.Print("Routers initialized -- success")
	return taskRouter, plannerRouter
}

// Setup initializes the router for task and planner controls and returns their references.
// Returns:
//   - *handle.TaskControl: The task control handle
//   - *handle.PlannerControl: The planner control handle
func Setup() (*handle.TaskControl, *handle.PlannerControl) {
	return initializeRouter()
}
