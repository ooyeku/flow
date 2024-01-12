package conf

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	"goworkflow/internal/inmemory"
	handle2 "goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	store2 "goworkflow/pkg/store"
	"log"
	"os"
)

// ConfigFileName is the name of the configuration file.
const (
	ConfigFileName   = "config.json"
	InMemoryDatabase = "internal/inmemory/goworkflow.db"
)

// LogAndExitOnError logs the given error message and exits the program if the error is not nil.
func LogAndExitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

// Config represents the configuration structure.
type Config struct {
	DatabaseType string `json:"database_type"`
}

// decodeConfig decodes the JSON data from the specified file and populates the fields of the Config structure accordingly.
// Parameters:
//
//	file: a pointer to the os.File object representing the JSON file to be decoded
//
// Returns:
//
//	nil if the decoding is successful, otherwise an error
func (c *Config) decodeConfig(file *os.File) error {
	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}

// createInMemoryDB creates a new in-memory StormDB instance.
// It returns a pointer to the StormDB instance and an error if any.
func createInMemoryDB() (*storm.DB, error) {
	db, err := storm.Open(InMemoryDatabase, storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// selectTaskDatabase selects the appropriate task database based on the configured database type in the Config struct.
// It returns an instance of TaskStore interface that corresponds to the selected database type.
//
// Parameters:
// - db: Pointer to the storm.DB instance to be passed to the task database constructor.
//
// Returns:
// - store2.TaskStore: The initialized task store implementation based on the configured database type.
// - error: An error if an invalid task database type is specified in the Config struct.
//
// Example usage:
//
//	// Assume config is an instance of Config struct
//	db, err := createInMemoryDB()
//	if err != nil {
//	  // Handle error
//	}
//
//	taskStore, err := config.selectTaskDatabase(db)
//	if err != nil {
//	  // Handle error
//	}
//	// Use the task store for performing CRUD operations on tasks
func (c *Config) selectTaskDatabase(db *storm.DB) (store2.TaskStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryTaskStore(db), nil
	}
	return nil, errors.New("invalid task database type specified in config")
}

// selectPlannerDatabase selects the appropriate PlannerStore based on the DatabaseType specified in the Config.
// It returns an instance of store2.PlannerStore and an error.
// If the DatabaseType is "bolt", it creates a new instance of inmemory.NewInMemoryPlannerStore and returns it as the PlannerStore.
// If the DatabaseType is not "bolt", it returns an error indicating that an invalid planner database type is specified.
func (c *Config) selectPlannerDatabase(db *storm.DB) (store2.PlannerStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryPlannerStore(db), nil
	}
	return nil, errors.New("invalid planner database type specified in config")
}

// selectGoalDatabase selects the appropriate goal store based on the database type specified in the configuration.
func (c *Config) selectGoalDatabase(db *storm.DB) (store2.GoalStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryGoalStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

// selectPlanDatabase selects the appropriate `PlanStore` implementation based on the `DatabaseType` specified in the `Config`.
//
// Parameters:
// - db: A storm.DB instance used for database operations.
//
// Returns:
// - store2.PlanStore: The selected PlanStore implementation.
// - error: An error if the specified database type is invalid.
//
// Possible errors:
// - "invalid database type specified in config": If the DatabaseType specified in the Config is invalid.
//
// Example usage:
// ```
// planStore, err := config.selectPlanDatabase(db)
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// // Use the planStore instance for further operations.
// ```
func (c *Config) selectPlanDatabase(db *storm.DB) (store2.PlanStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryPlanStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

// initializeConfiguration opens the configuration file at the specified file path, decodes its contents into a Config object,
// and returns the Config object or an error if any occurs.
func initializeConfiguration(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(file)
	var config Config
	err = config.decodeConfig(file)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// initializeRouter initializes the routers for task, planner, goal, and plan.
// It returns the taskRouter, plannerRouter, goalRouter, and planRouter.
func initializeRouter() (*handle2.TaskControl, *handle2.PlannerControl, *handle2.GoalControl, *handle2.PlanControl) {
	config, err := initializeConfiguration(ConfigFileName)
	LogAndExitOnError(err, "failed to initialize configuration: %v")

	db, err := createInMemoryDB()
	LogAndExitOnError(err, "failed to create in-memory database: %v")

	taskStore, err := config.selectTaskDatabase(db)
	LogAndExitOnError(err, "failed to select task database: %v")
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle2.NewTaskControl(taskService)

	plannerStore, err := config.selectPlannerDatabase(db)
	LogAndExitOnError(err, "failed to select planner database: %v")
	plannerService := services.NewPlannerService(plannerStore)
	plannerRouter := handle2.NewPlannerControl(plannerService)

	goalStore, err := config.selectGoalDatabase(db)
	LogAndExitOnError(err, "failed to select goal database: %v")
	goalService := services.NewGoalService(goalStore)
	goalRouter := handle2.NewGoalControl(goalService)

	planStore, err := config.selectPlanDatabase(db)
	LogAndExitOnError(err, "failed to select plan database: %v")
	planService := services.NewPlanService(planStore)
	planRouter := handle2.NewPlanControl(planService)

	log.Print("Routers initialized -- success")
	return taskRouter, plannerRouter, goalRouter, planRouter
}

// Setup initializes the router and returns the necessary controller instances.
// It returns the following:
// - TaskControl: The controller for handling task-related operations.
// - PlannerControl: The controller for handling planner-related operations.
// - GoalControl: The controller for handling goal-related operations.
// - PlanControl: The controller for handling plan-related operations.
func Setup() (*handle2.TaskControl, *handle2.PlannerControl, *handle2.GoalControl, *handle2.PlanControl) {
	return initializeRouter()
}
