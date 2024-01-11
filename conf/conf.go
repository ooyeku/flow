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

const (
	ConfigFileName   = "config.json"
	InMemoryDatabase = "store/inmemory/goworkflow.db"
)

func LogAndExitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

type Config struct {
	DatabaseType string `json:"database_type"`
}

func (c *Config) decodeConfig(file *os.File) error {
	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}

func createInMemoryDB() (*storm.DB, error) {
	db, err := storm.Open(InMemoryDatabase, storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *Config) selectTaskDatabase(db *storm.DB) (store.TaskStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryTaskStore(db), nil
	}
	return nil, errors.New("invalid task database type specified in config")
}

func (c *Config) selectPlannerDatabase(db *storm.DB) (store.PlannerStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryPlannerStore(db), nil
	}
	return nil, errors.New("invalid planner database type specified in config")
}

func (c *Config) selectGoalDatabase(db *storm.DB) (store.GoalStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryGoalStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

func (c *Config) selectPlanDatabase(db *storm.DB) (store.PlanStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory.NewInMemoryPlanStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

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

func initializeRouter() (*handle.TaskControl, *handle.PlannerControl, *handle.GoalControl, *handle.PlanControl) {
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

	goalStore, err := config.selectGoalDatabase(db)
	LogAndExitOnError(err, "failed to select goal database: %v")
	goalService := services.NewGoalService(goalStore)
	goalRouter := handle.NewGoalControl(goalService)

	planStore, err := config.selectPlanDatabase(db)
	LogAndExitOnError(err, "failed to select plan database: %v")
	planService := services.NewPlanService(planStore)
	planRouter := handle.NewPlanControl(planService)

	log.Print("Routers initialized -- success")
	return taskRouter, plannerRouter, goalRouter, planRouter
}

func Setup() (*handle.TaskControl, *handle.PlannerControl, *handle.GoalControl, *handle.PlanControl) {
	return initializeRouter()
}
