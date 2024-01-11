package inmemory

import (
	"github.com/asdine/storm"
	"goworkflow/internal/models"
)

// BoltTaskStore is a type that represents a task store backed by a BoltDB database.
type BoltTaskStore struct {
	db *storm.DB
}

// NewInMemoryTaskStore creates a new in-memory task store with the given storm.DB instance.
// It returns a pointer to the BoltTaskStore.
// Usage example:
//
//	db, err := storm.Open("test.db", storm.BoltOptions(0600, nil))
//	if err != nil {
//		log.Fatalf("Failed to open database: %v", err)
//	}
//	defer func(db *storm.DB) {
//		err := db.Close()
//		if err != nil {
//			log.Fatalf("Failed to close database: %v", err)
//		}
//	}(db)
//
//	// Create a new in-memory store
//	inMemoryStore := NewInMemoryTaskStore(db)
//
//	// create a new taskMake service
//	taskService := services.NewTaskService(inMemoryStore)
//
//	// Create a taskMake handler
//	taskHandler := handle.NewTaskControl(taskService)
//
//	log.Printf("Task handler: %v", taskHandler)
//
//	// Create a new taskMake
//	taskMake, err := taskHandler.CreateTask(handle.CreateTaskRequest{
//		Title:       "Task 1",
//		Description: "Task 1 description",
//		Owner:       "Me",
//	})
//
//	if err != nil {
//		log.Fatalf("Failed to create taskMake: %v", err)
//	}
//	log.Printf("Task created: %s", taskMake.ID)
//
//	// Get the taskMake
//	taskGet, err := taskHandler.GetTask(&handle.GetTaskRequest{ID: taskMake.ID})
//
//	if err != nil {
//		log.Fatalf("Failed to get taskMake: %v", err)
//	}
//
//	log.Printf("Task retrieved: %s", taskGet.Title)
//	log.Printf("Description: %s", taskGet.Description)
//	log.Printf("Owner: %s", taskGet.Owner)
//	log.Printf("Started: %t", taskGet.Started)
func NewInMemoryTaskStore(db *storm.DB) *BoltTaskStore {
	return &BoltTaskStore{
		db: db,
	}
}

// CreateTask method creates a new task in the BoltTaskStore.
// It takes a pointer to a models.Task object as a parameter.
// Returns an error if the operation fails.
func (s *BoltTaskStore) CreateTask(task *models.Task) error {
	return s.db.Save(task)
}

// UpdateTask updates a task with the specified ID. It takes the ID string and the task struct as input parameters.
// It assigns the provided ID to the task's ID field and then updates the task in the BoltDB using the db.Update method.
// Returns an error if there was an issue while updating the task in the BoltDB.
func (s *BoltTaskStore) UpdateTask(id string, task *models.Task) error {
	task.ID = id
	return s.db.Update(task)
}

// DeleteTask deletes a task from the BoltTaskStore.
// It takes the ID of the task as a parameter and returns an error if any
// occurred during the deletion process.
// The function first creates a new Task object with the provided ID,
// then calls the DeleteStruct method of the BoltDB instance to remove
// the task from the database.
//
// Example usage:
// err := myTaskStore.DeleteTask("task-123")
//
//	if err != nil {
//	   fmt.Println("Error deleting task:", err)
//	}
func (s *BoltTaskStore) DeleteTask(id string) error {
	task := new(models.Task)
	task.ID = id
	return s.db.DeleteStruct(task)
}

// GetTask retrieves a task from the BoltTaskStore based on the given ID.
// It returns the task and an error if any occurred.
// If the task with the given ID is not found, it returns nil and an error.
func (s *BoltTaskStore) GetTask(id string) (*models.Task, error) {
	task := new(models.Task)
	if err := s.db.One("ID", id, task); err != nil {
		return nil, err
	}
	return task, nil
}

// ListTasks retrieves a list of tasks from the BoltTaskStore.
func (s *BoltTaskStore) ListTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	if err := s.db.All(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
