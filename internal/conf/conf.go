package conf

import (
	"github.com/asdine/storm"
	"goworkflow/internal/inmemory"
	"goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	"sync"
)

// The `once` variable is a sync.Once type that guarantees an initialization function is executed only once.
// The `dbPath` variable is a string that stores the file path of the database.
// The `taskRouter` variable is a pointer to a handle.TaskControl struct that handles task-related operations.
var (
	once       sync.Once
	dbPath     string = "internal/inmemory/goworkflow.db"
	taskRouter *handle.TaskControl
)

// initializeRouter initializes the router by creating a new TaskControl instance and returning it.
// It opens the database at `dbPath` using storm, and panics if there is an error.
// It then creates a new InMemoryTaskStore using the opened database,
// creates a new TaskService using the task store, and finally creates a new TaskControl using the task service.
// The created TaskControl is assigned to the global variable taskRouter,
// which is returned at the end of the function.
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

// GetDBPath returns the path of the database used by the application.
func GetDBPath() string {
	return dbPath
}
