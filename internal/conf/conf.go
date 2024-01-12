package conf

import (
	"github.com/asdine/storm"
	"goworkflow/internal/inmemory"
	"goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	"sync"
)

var (
	once       sync.Once
	dbPath     string = "internal/inmemory/goworkflow.db"
	taskRouter *handle.TaskControl
)

func initializeRouter() *handle.TaskControl {
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		panic(err)
	}

	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter = handle.NewTaskControl(taskService)

	return taskRouter
}

func GetTaskRouter() *handle.TaskControl {
	once.Do(func() {
		taskRouter = initializeRouter()
	})
	return taskRouter
}

func GetDBPath() string {
	return dbPath
}
