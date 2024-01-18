package handle

import (
	_ "encoding/json"
	"flow/internal/models"
	"flow/pkg/services"
	"github.com/google/uuid"
	"time"
)

// TaskControl represents a controller for managing tasks.
type TaskControl struct {
	service *services.TaskService
}

// NewTaskControl creates a new instance of TaskControl.
// It takes a pointer to a TaskService as a parameter and returns a pointer to TaskControl.
// TaskControl is a struct that contains a reference to a TaskService.
func NewTaskControl(service *services.TaskService) *TaskControl {
	return &TaskControl{
		service: service,
	}
}

// CreateTaskResponse represents the response object when creating a task.
type CreateTaskResponse struct {
	ID string `json:"id"`
}

// CreateTaskRequest represents a request to create a new task.
// It contains the title, description, and owner of the task.
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

// DeleteTask deletes a task based on the provided request ID. It calls the Store's DeleteTask method.
func (c *TaskControl) DeleteTask(req *DeleteTaskRequest) error {
	if err := c.service.Store.DeleteTask(req.ID); err != nil {
		return err
	}
	return nil
}

// TaskRequest represents an interface for retrieving the ID of a task.
type TaskRequest interface {
	GetID() string
}

// GetTaskRequest represents a request to get a task.
type GetTaskRequest struct {
	ID string `json:"id"`
}

// GetID returns the ID of the GetTaskRequest.
func (req *GetTaskRequest) GetID() string {
	return req.ID
}

// GetTaskByTitleRequest represents a request to get a task by its title.
// The title field is used to specify the title of the task to retrieve.
type GetTaskByTitleRequest struct {
	Title string `json:"title"`
}

// GetID returns the title of the task as the unique identifier.
func (req *GetTaskByTitleRequest) GetID() string {
	return req.Title
}

// GetTaskByOwnerRequest represents a request to get tasks owned by a specific owner.
type GetTaskByOwnerRequest struct {
	Owner string `json:"owner"`
}

// GetID implements the TaskRequest interface for GetTaskByOwnerRequest.
func (req *GetTaskByOwnerRequest) GetID() string {
	// gets the id of the task by its owner
	return req.Owner
}

// GetTaskResponse represents the response structure for getting a task.
// It contains the ID, title, description, owner, started status, completed status, creation timestamp, and update timestamp of the task.
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

// GetTask retrieves the task with the specified ID from the service. It returns the task details in the response.
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

// GetTaskByTitle returns a task with the specified title. It retrieves the task from the service by calling GetTaskByTitle. If the task is found, it is returned in the GetTaskResponse
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

// GetTaskByOwner retrieves a task based on the provided owner. It calls the service's GetTaskByOwner method, passing the owner as a parameter. If an error occurs, it returns nil and
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

// ListTasksResponse represents a response containing a list of tasks.
// It contains an array of GetTaskResponse objects.
// Each GetTaskResponse object represents a single task with its details.
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
// It returns the generated UUID as a string and an error if any.
// Example usage:
// id, err := generateTaskUUID()
func generateTaskUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
