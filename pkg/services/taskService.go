package services

import (
	"goworkflow/internal/models"
	"goworkflow/pkg/store"
)

// TaskService represents a service for managing tasks.
type TaskService struct {
	Store store.TaskStore
}

// NewTaskService creates a new instance of TaskService using the provided TaskStore.
// It returns a pointer to the TaskService.
func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{
		Store: store,
	}
}

// CreateTask creates a new task using the provided task object.
// It returns an error if there was a problem creating the task.
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Store.CreateTask(task)
}

// UpdateTask updates the task with the specified ID using the provided task data.
func (s *TaskService) UpdateTask(id string, task *models.Task) error {
	return s.Store.UpdateTask(id, task)
}

// DeleteTask deletes a task by its ID.
// It calls the DeleteTask method of the TaskStore to delete the task from the storage.
// If the task is successfully deleted, it returns nil.
// If there is an error during the deletion process, it returns the error.
func (s *TaskService) DeleteTask(id string) error {
	return s.Store.DeleteTask(id)
}

// GetTask retrieves a task with the given ID.
// It returns the task and an error if the task could not be found.
func (s *TaskService) GetTask(id string) (*models.Task, error) {
	return s.Store.GetTask(id)
}

// ListTasks retrieves a list of tasks from the TaskService.
//
// It invokes the ListTasks method of the TaskService's Store to fetch the tasks.
//
// The method returns a slice of Task pointers and an error.
// The slice holds the Task objects retrieved from the Store.
// If there are no tasks, the returned slice will be empty.
// If an error occurs during the retrieval process, it will be returned in the error parameter.
//
// Example usage:
// tasks, err := taskService.ListTasks()
//
//	if err != nil {
//	  // handle error
//	}
//
//	for _, task := range tasks {
//	  // do something with task
//	}
//
// For more information about the Task struct, see the models.Task declaration.
//
// For more information about the TaskService struct and the TaskStore interface, see their respective declarations.
func (s *TaskService) ListTasks() ([]*models.Task, error) {
	return s.Store.ListTasks()
}

func (s *TaskService) GetTaskByTitle(title string) (*models.Task, error) {
	return s.Store.GetTaskByTitle(title)
}

func (s *TaskService) GetTaskByOwner(owner string) (*models.Task, error) {
	return s.Store.GetTaskByOwner(owner)
}
