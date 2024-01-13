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
// It takes in a pointer to a TaskService and returns a pointer to TaskControl.
// TaskControl is a struct that provides control functions for tasks, such as creating, updating, deleting, getting, and listing tasks.
// Usage example:
//
//	taskService := &services.TaskService{
//	    Store: &store.TaskStore{}, // replace with your TaskStore implementation
//	}
//	taskControl := NewTaskControl(taskService)
//	// Use taskControl to perform task operations
func NewTaskControl(service *services.TaskService) *TaskControl {
	return &TaskControl{
		service: service,
	}
}

// CreateTaskResponse represents the response object for creating a task.
type CreateTaskResponse struct {
	ID string `json:"id"`
}

// CreateTaskRequest represents a request to create a new task.
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
	m := &models.Task{}
	task := m.GenerateTaskInstance(id, req.Title, req.Description, req.Owner)
	err = c.service.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return &CreateTaskResponse{
		ID: task.ID,
	}, nil
}

// UpdateTaskRequest represents a request to update a task.
// It contains the ID of the task to be updated, as well as the updated values for title, description, owner, started, and completed.
type UpdateTaskRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Started     bool   `json:"started"`
	Completed   bool   `json:"completed"`
}

// UpdateTask updates an existing task with the provided request. It retrieves the task from the service using req.ID, and updates its properties with the values from the request. Finally
func (c *TaskControl) UpdateTask(req *UpdateTaskRequest) error {
	task, err := c.service.GetTask(req.ID)
	if err != nil {
		return err
	}

	m := &models.Task{}
	task = m.GenerateTaskInstance(req.ID, req.Title, req.Description, req.Owner)
	if err != nil {
		return err
	}

	// these fields are not updated by GenerateTaskInstance
	task.Started = req.Started
	task.Completed = req.Completed
	task.UpdatedAt = time.Now()

	if err := c.service.UpdateTask(req.ID, task); err != nil {
		return err
	}
	return nil
}

// DeleteTaskRequest represents a request to delete a task.
type DeleteTaskRequest struct {
	ID string `json:"id"`
}

// DeleteTask deletes a task with the provided ID from the store. It returns an error if the task deletion fails.
func (c *TaskControl) DeleteTask(req *DeleteTaskRequest) error {
	if err := c.service.Store.DeleteTask(req.ID); err != nil {
		return err
	}
	return nil
}

// TaskRequest is the interface that wraps the GetID method so that it can
// be used by GetTaskRequest, GetTaskByTitleRequest, and GetTaskByOwnerRequest.
type TaskRequest interface {
	GetID() string
}

// GetTaskRequest represents a request to get a task by its ID.
type GetTaskRequest struct {
	ID string `json:"id"`
}

// GetID implements the TaskRequest interface for GetTaskRequest.
func (req *GetTaskRequest) GetID() string {
	return req.ID
}

// GetTaskByTitleRequest represents a request to get a task by its title.
type GetTaskByTitleRequest struct {
	Title string `json:"title"`
}

// GetID implements the TaskRequest interface for GetTaskByTitleRequest.
func (req *GetTaskByTitleRequest) GetID() string {
	return req.Title
}

// GetTaskByOwnerRequest represents a request to get a task by its owner.
type GetTaskByOwnerRequest struct {
	Owner string `json:"owner"`
}

// GetID implements the TaskRequest interface for GetTaskByOwnerRequest.
func (req *GetTaskByOwnerRequest) GetID() string {
	// gets the id of the task by its owner
	return req.Owner
}

// GetTaskResponse represents the response data structure for retrieving a task.
// ID represents the unique identifier of the task.
// Title represents the title of the task.
// Description represents the description of the task.
// Owner represents the owner of the task.
// Started indicates whether the task has been started or not.
// Completed indicates whether the task has been completed or not.
// CreatedAt represents the timestamp when the task was created.
// UpdatedAt represents the timestamp when the task was last updated.
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

// GetTask retrieves information about a task based on the provided request ID. It calls the service's GetTask method and returns a GetTaskResponse with the task details. If an error
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

func (c *TaskControl) GetTaskByTitle(req *GetTaskByTitleRequest) (*GetTaskResponse, error) {
	task, err := c.service.GetTaskByTitle(req.Title)
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

func (c *TaskControl) GetTaskByOwner(req *GetTaskByOwnerRequest) (*GetTaskResponse, error) {
	task, err := c.service.GetTaskByOwner(req.Owner)
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

// ListTasksResponse is a struct that represents the response for listing tasks. It contains an array of GetTaskResponse objects, named Tasks, which are the tasks retrieved.
// The JSON field tag for the Tasks array is "tasks", indicating how it should be serialized and deserialized when encoding and decoding JSON data.
type ListTasksResponse struct {
	Tasks []*GetTaskResponse `json:"tasks"`
}

// ListTasks returns a list of task responses by calling the `ListTasks` method of the `c.service` service.
// It transforms each task into a `GetTaskResponse` and appends it to the `taskResponses` slice.
// It returns the `taskResponses` slice and nil error if successful.
// If there is an error while listing tasks, it returns nil and the error encountered.
func (c *TaskControl) ListTasks() ([]GetTaskResponse, error) {
	tasks, err := c.service.ListTasks()
	if err != nil {
		return nil, err
	}
	var taskResponses []GetTaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, GetTaskResponse{
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
	return taskResponses, nil
}

// generateTaskUUID generates a unique task UUID using the uuid.NewRandom() function.
// It returns the generated UUID as a string and any error encountered during the process.
func generateTaskUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
