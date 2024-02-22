package models

import (
	"testing"
)

func TestGenerateTaskInstance(t *testing.T) {
	task := Task{}
	id := "test-id"
	title := "test-title"
	description := "test-description"
	owner := "test-owner"

	result := task.GenerateTaskInstance(id, title, description, owner)

	if result.ID != id {
		t.Errorf("Expected ID %s, but got %s", id, result.ID)
	}

	if result.Title != title {
		t.Errorf("Expected Title %s, but got %s", title, result.Title)
	}

	if result.Description != description {
		t.Errorf("Expected Description %s, but got %s", description, result.Description)
	}

	if result.Owner != owner {
		t.Errorf("Expected Owner %s, but got %s", owner, result.Owner)
	}

	if result.Started != false {
		t.Errorf("Expected Started to be false, but got %v", result.Started)
	}

	if result.Completed != false {
		t.Errorf("Expected Completed to be false, but got %v", result.Completed)
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
