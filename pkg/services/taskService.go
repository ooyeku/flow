package services

import (
	"flow/internal/models"
	"flow/pkg/store"
	"log"
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

// ListTasks retrieves all tasks from the store.
func (s *TaskService) ListTasks() ([]*models.Task, error) {
	tasks, err := s.Store.ListTasks()
	if err != nil {
		log.Printf("Error in TaskService.ListTasks: %s", err)
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByTitle(title string) (*models.Task, error) {
	return s.Store.GetTaskByTitle(title)
}

func (s *TaskService) GetTaskByOwner(owner string) (*models.Task, error) {
	return s.Store.GetTaskByOwner(owner)
}
