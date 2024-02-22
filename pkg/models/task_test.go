package models

import (
	"testing"
)

func TestGenerateTaskInstance(t *testing.T) {
	task := createTask()

	if task.ID != "test-id" {
		t.Errorf("Expected ID test-id, but got %s", task.ID)
	}

	if task.Title != "test-title" {
		t.Errorf("Expected Title test-title, but got %s", task.Title)
	}

	if task.Description != "test-description" {
		t.Errorf("Expected Description test-description, but got %s", task.Description)
	}

	if task.Owner != "test" {
		t.Errorf("Expected Owner test, but got %s", task.Owner)
	}

	if task.Started != false {
		t.Errorf("Expected Started to be false, but got true")
	}

	if task.Completed != false {
		t.Errorf("Expected Completed to be false, but got true")
	}
}

func TestIsStarted(t *testing.T) {
	task := Task{}
	task.Started = true

	if task.IsStarted() != true {
		t.Errorf("Expected IsStarted to return true, but got false")
	}
}

func TestIsCompleted(t *testing.T) {
	task := Task{}
	task.Completed = true

	if task.IsCompleted() != true {
		t.Errorf("Expected IsCompleted to return true, but got false")
	}
}

func TestStart(t *testing.T) {
	task := Task{}
	task.Start()

	if task.Started != true {
		t.Errorf("Expected Started to be true, but got false")
	}
}

func TestComplete(t *testing.T) {
	task := Task{}
	task.Completed = true

	if task.Completed != true {
		t.Errorf("Expected Completed to be true, but got false")
	}
}

// utility for creating a task
func createTask() *Task {
	return &Task{
		ID:          "test-id",
		Title:       "test-title",
		Description: "test-description",
		Owner:       "test",
		Started:     false,
		Completed:   false,
	}
}
