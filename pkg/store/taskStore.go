package store

import (
	"github.com/ooyeku/flow/internal/models"
)

// TaskStore represents an interface for managing tasks
// Implementations of this interface should provide methods to create, update, delete, get, and list tasks
type TaskStore interface {
	CreateTask(task *models.Task) error
	UpdateTask(id string, task *models.Task) error
	DeleteTask(id string) error
	GetTask(id string) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	GetTaskByTitle(title string) (*models.Task, error)
	GetTaskByOwner(owner string) ([]*models.Task, error)
}
