package models

import "time"

// Task represents a to-do item
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Owner       string    `json:"owner"`
	Started     bool      `json:"started"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GenerateTaskInstance generates a new instance of the Task struct with the provided parameters.
// The generated Task instance has its ID, Title, Description, and Owner fields set to the provided values.
// The Started and Completed fields are set to false, and the CreatedAt and UpdatedAt fields are set to the current time.
//
// Example usage:
//
//	m := &models.Task{}
//	task := m.GenerateTaskInstance(id, req.Title, req.Description, req.Owner)
//	err = c.service.CreateTask(task)
//	if err != nil {
//		return nil, err
//	}
//
// Parameters:
// - id: The ID of the task instance.
// - title: The title of the task instance.
// - description: The description of the task instance.
// - owner: The owner of the task instance.
//
// Returns:
// - *Task: The generated task instance.
func (t *Task) GenerateTaskInstance(id, title, description, owner string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Owner:       owner,
		Started:     false,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// IsStarted returns the status of the task indicating if it has been started or not.
func (t *Task) IsStarted() bool {
	return t.Started
}

// IsCompleted returns a boolean value indicating whether the task is completed or not.
func (t *Task) IsCompleted() bool {
	return t.Completed
}

// Start marks the task as started.
func (t *Task) Start() {
	t.Started = true
}
