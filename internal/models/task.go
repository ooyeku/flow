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

func (t *Task) IsStarted() bool {
	return t.Started
}

func (t *Task) IsCompleted() bool {
	return t.Completed
}

func (t *Task) Start() {
	t.Started = true
}
