package conf

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	handle2 "goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	store2 "goworkflow/pkg/store"
	inmemory2 "goworkflow/pkg/store/inmemory"
	"log"
	"os"
)

const (
	ConfigFileName   = "config.json"
	InMemoryDatabase = "pkg/store/inmemory/goworkflow.db"
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

func (c *Config) selectTaskDatabase(db *storm.DB) (store2.TaskStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory2.NewInMemoryTaskStore(db), nil
	}
	return nil, errors.New("invalid task database type specified in config")
}

func (c *Config) selectPlannerDatabase(db *storm.DB) (store2.PlannerStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory2.NewInMemoryPlannerStore(db), nil
	}
	return nil, errors.New("invalid planner database type specified in config")
}

func (c *Config) selectGoalDatabase(db *storm.DB) (store2.GoalStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory2.NewInMemoryGoalStore(db), nil
	}
	return nil, errors.New("invalid database type specified in config")
}

func (c *Config) selectPlanDatabase(db *storm.DB) (store2.PlanStore, error) {
	switch c.DatabaseType {
	case "bolt":
		return inmemory2.NewInMemoryPlanStore(db), nil
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

func Setup() (*handle2.TaskControl, *handle2.PlannerControl, *handle2.GoalControl, *handle2.PlanControl) {
	return initializeRouter()
}
