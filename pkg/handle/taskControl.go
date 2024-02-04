package handle

import (
	_ "encoding/json"
	"github.com/google/uuid"
	"github.com/ooyeku/flow/internal/models"
	"github.com/ooyeku/flow/pkg/services"
	"time"
)

// TaskControl represents a controller for managing tasks.
type TaskControl struct {
	service *services.TaskService
}

// NewTaskControl creates a new instance of the TaskControl struct.
// It takes a pointer to a TaskService as a parameter and returns a pointer to a TaskControl struct.
// Example usage:
// service := &services.TaskService{}
// taskControl := NewTaskControl(service)
func NewTaskControl(service *services.TaskService) *TaskControl {
	return &TaskControl{
		service: service,
	}
}

// CreateTaskResponse represents the response from creating a task.
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

// UpdateTaskRequest represents a request for updating a task.
// It contains the ID of the task, along with the updated title, description, owner, started flag, and completed flag.
type UpdateTaskRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Started     bool   `json:"started"`
	Completed   bool   `json:"completed"`
}

// UpdateTask updates an existing task with the provided request.
// It retrieves the task from the service using the provided task ID.
// Then it generates a new task instance with the updated information
// and updates the relevant fields (Started, Completed, UpdatedAt).
// Finally, it calls the UpdateTask method of the service to save the changes.
// Returns an error if any operation fails.
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

// DeleteTaskRequest represents a request to delete a task with a given ID.
type DeleteTaskRequest struct {
	ID string `json:"id"`
}

// DeleteTask deletes a task with the provided ID.
// It calls the DeleteTask method of the service's store and returns any error that occurred.
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

// GetTaskRequest represents a request to get a task by ID.
// It contains a single field ID which is the ID of the task to be retrieved.
// Usage example:
//
//	req := &GetTaskRequest{
//	  ID: "12345",
//	}
//	task, err := t.GetTask(req)
type GetTaskRequest struct {
	ID string `json:"id"`
}

// GetID retrieves the ID of the GetTaskRequest.
func (req *GetTaskRequest) GetID() string {
	return req.ID
}

// GetTaskByTitleRequest represents a request to get a task by its title.
// The Title field is used to specify the title of the task.
type GetTaskByTitleRequest struct {
	Title string `json:"title"`
}

// GetID returns the title of the task as the ID for the GetTaskByTitleRequest
func (req *GetTaskByTitleRequest) GetID() string {
	return req.Title
}

// GetTaskByOwnerRequest represents a request to get tasks by owner.
// It contains the owner of the tasks.
type GetTaskByOwnerRequest struct {
	Owner string `json:"owner"`
}

// GetID gets the id of the task by its owner. It returns the owner as the id of the task.
func (req *GetTaskByOwnerRequest) GetID() string {
	// gets the id of the task by its owner
	return req.Owner
}

// GetTaskResponse represents the response structure for getting a task.
// ID represents the unique identifier of the task.
// Title represents the title of the task.
// Description represents the description of the task.
// Owner represents the owner of the task.
// Started represents whether the task has been started or not.
// Completed represents whether the task has been completed or not.
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

// GetTask retrieves a task with the specified ID from the service's store.
// It returns the task details in a GetTaskResponse object.
// If there is an error retrieving the task, nil and the error are returned.
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

// GetTaskByTitle retrieves a task by its title from the service's store.
// It returns the task information in a GetTaskResponse struct.
// If the task is not found or an error occurs, it returns nil and the error.
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

// GetTaskByOwner retrieves the tasks with the given owner from the service's store. It returns the task details in the response.
// If an error occurs during the retrieval process, it is returned as well.
func (c *TaskControl) GetTaskByOwner(req *GetTaskByOwnerRequest) ([]*GetTaskResponse, error) {
	tasks, err := c.service.GetTaskByOwner(req.Owner)
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
	return taskResponses, nil
}

// ListTasksResponse represents the response structure containing a list of tasks.
type ListTasksResponse struct {
	Tasks []*GetTaskResponse `json:"tasks"`
}

// ListTasks retrieves all tasks using the service's store and returns a list of GetTaskResponse objects representing the tasks.
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

// generateTaskUUID generates a unique task UUID using the uuid package.
// It returns the generated UUID as a string and any error that occurred during the generation process.
func generateTaskUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
