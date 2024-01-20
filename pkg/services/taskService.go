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

// CreateTask creates a new task using the provided task data.
// It returns an error if the task creation fails.
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Store.CreateTask(task)
}

// UpdateTask updates the task with the specified ID using the provided task object.
// It returns an error if there was a problem updating the task.
func (s *TaskService) UpdateTask(id string, task *models.Task) error {
	return s.Store.UpdateTask(id, task)
}

// DeleteTask deletes a task with the given ID.
// It returns an error if there was a problem deleting the task.
func (s *TaskService) DeleteTask(id string) error {
	return s.Store.DeleteTask(id)
}

// GetTask retrieves a task with the specified ID from the task store.
// It takes an ID string as the input parameter.
// It returns a pointer to a Task object and an error.
// The Task object contains the details of the retrieved task, including the ID, title, description, owner, started status, completed status, creation timestamp, and update timestamp
func (s *TaskService) GetTask(id string) (*models.Task, error) {
	return s.Store.GetTask(id)
}

// ListTasks retrieves a list of tasks from the TaskService.
// It returns a slice of task objects and an error if there was a problem fetching the tasks.
func (s *TaskService) ListTasks() ([]*models.Task, error) {
	tasks, err := s.Store.ListTasks()
	if err != nil {
		log.Printf("Error in TaskService.ListTasks: %s", err)
		return nil, err
	}
	return tasks, nil
}

// GetTaskByTitle retrieves a task by its title.
// It takes a string parameter `title` representing the title of the task.
// It returns a pointer to a `models.Task` object and an error.
// The returned task object contains the details of the task with the matching title.
// If no task is found with the given title, the function returns `nil` along with an error.
// Example usage:
//
//	req := &GetTaskByTitleRequest{
//	    Title: "Example Task",
//	}
//
// task, err := c.GetTaskByTitle(req)
//
//	if err != nil {
//	    // Handle error
//	}
//
// fmt.Println(task.ID)
// fmt.Println(task.Title)
// fmt.Println(task.Description)
// fmt.Println(task.Owner)
// fmt.Println(task.Started)
// fmt.Println(task.Completed)
// fmt.Println(task.CreatedAt)
// fmt.Println(task.UpdatedAt)
func (s *TaskService) GetTaskByTitle(title string) (*models.Task, error) {
	return s.Store.GetTaskByTitle(title)
}

// GetTaskByOwner returns a task owned by the given owner.
// It takes an owner string as a parameter and returns the task associated with that owner.
// If there is no task found or an error occurs, it returns nil and an error respectively.
func (s *TaskService) GetTaskByOwner(owner string) (*models.Task, error) {
	return s.Store.GetTaskByOwner(owner)
}
