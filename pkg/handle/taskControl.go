package handle

import (
	_ "encoding/json"
	"github.com/google/uuid"
	"goworkflow/internal/models"
	"goworkflow/pkg/services"
	"time"
)

// TaskControl represents a controller for managing tasks.
type TaskControl struct {
	service *services.TaskService
}

// NewTaskControl creates a new instance of TaskControl.
// It takes a pointer to a `TaskService` as input and returns a pointer to a `TaskControl` struct.
// The `TaskControl` struct has methods to create, update, delete, get, and list tasks.
// To create a new `TaskControl`, pass a pointer to a `TaskService` to the `NewTaskControl` function.
// Usage example:
func NewTaskControl(service *services.TaskService) *TaskControl {
	return &TaskControl{
		service: service,
	}
}

// CreateTaskResponse represents the response model for the CreateTask method in the TaskControl type.
// ID represents the unique identifier of the newly created task.
type CreateTaskResponse struct {
	ID string `json:"id"`
}

// CreateTaskRequest represents the request payload for creating a task. It contains the following fields:
// - Title: The title of the task
// - Description: The description of the task
// - Owner: The owner of the task
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

// CreateTask generates a unique id for the task and creates a new task with the provided request. It saves the task using the service's store and returns the task id in the response
func (c *TaskControl) CreateTask(req CreateTaskRequest) (*CreateTaskResponse, error) {

	// generate unique id
	id, err := generateTaskUUID()
	if err != nil {
		return nil, err
	}
	task := &models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Owner:       req.Owner,
		Started:     false,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = c.service.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return &CreateTaskResponse{
		ID: task.ID,
	}, nil
}

// UpdateTaskRequest represents the request to update a task with the specified ID.
type UpdateTaskRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Started     bool   `json:"started"`
	Completed   bool   `json:"completed"`
}

// UpdateTask updates a task with the specified request parameters.
// It retrieves the task using the provided ID, updates its properties with the values from the request,
// and calls the UpdateTask method on the TaskService's Store to persist the changes.
// If any error occurs during the process, it will be returned.
func (c *TaskControl) UpdateTask(req *UpdateTaskRequest) error {
	task, err := c.service.GetTask(req.ID)
	if err != nil {
		return err
	}
	task.Title = req.Title
	task.Description = req.Description
	task.Owner = req.Owner
	task.Started = req.Started
	task.Completed = req.Completed
	task.UpdatedAt = time.Now()
	if err := c.service.Store.UpdateTask(req.ID, task); err != nil {
		return err
	}
	return nil
}

// DeleteTaskRequest represents the request to delete a task.
type DeleteTaskRequest struct {
	ID string `json:"id"`
}

// DeleteTask deletes a task from the task store.
func (c *TaskControl) DeleteTask(req *DeleteTaskRequest) error {
	if err := c.service.Store.DeleteTask(req.ID); err != nil {
		return err
	}
	return nil
}

// GetTaskRequest represents a request to get a task.
type GetTaskRequest struct {
	ID string `json:"id"`
}

// GetTaskResponse represents the response structure for the GetTask API.
// Field ID represents the ID of the task.
// Field Title represents the title of the task.
// Field Description represents the description of the task.
// Field Owner represents the owner of the task.
// Field Started represents whether the task has been started.
// Field Completed represents whether the task has been completed.
// Field CreatedAt represents the creation time of the task.
// Field UpdatedAt represents the last update time of the task.
type GetTaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Owner       string    `json:"owner"`
	Started     bool      `json:"started"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetTask retrieves a task from the task service based on the provided task ID. It returns the task information in the form of a GetTaskResponse struct or an error if the task is not
func (c *TaskControl) GetTask(req *GetTaskRequest) (*GetTaskResponse, error) {
	task, err := c.service.GetTask(req.ID)
	if err != nil {
		return nil, err
	}
	return &GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Owner:       task.Owner,
		Started:     task.Started,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

// ListTasksResponse represents the response structure for a list of tasks.
// It contains an array of GetTaskResponse, which represents the individual tasks.
type ListTasksResponse struct {
	Tasks []*GetTaskResponse `json:"tasks"`
}

// ListTasks returns a list of tasks.
//
// The method retrieves all tasks using the ListTasks method of the TaskService.
// It then maps the tasks to GetTaskResponse objects and returns a ListTasksResponse
// containing the list of mapped tasks.
//
// Returns:
// - *ListTasksResponse: The list of tasks.
// - error: An error if the retrieval fails.
func (c *TaskControl) ListTasks() (*ListTasksResponse, error) {
	tasks, err := c.service.ListTasks()
	if err != nil {
		return nil, err
	}
	var taskResponses []*GetTaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, &GetTaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Owner:       task.Owner,
			Started:     task.Started,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return &ListTasksResponse{
		Tasks: taskResponses,
	}, nil
}

// generateUUID generates a new UUID (Universally Unique Identifier).
// It returns the UUID as a string and any error encountered.
func generateTaskUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
