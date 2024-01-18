package models

import (
	"testing"
	"time"
)

func TestGenerateGoalInstance(t *testing.T) {
	g := Goal{}
	id := "test-id"
	objective := "test-objective"
	deadline := time.Now()

	result := g.GenerateGoalInstance(id, objective, deadline)

	if result.Id != id {
		t.Errorf("Expected Id %s, but got %s", id, result.Id)
	}

	if result.Objective != objective {
		t.Errorf("Expected Objective %s, but got %s", objective, result.Objective)
	}

	if !result.Deadline.Equal(deadline) {
		t.Errorf("Expected Deadline %v, but got %v", deadline, result.Deadline)
	}

	if result.GoalStatus != NotStarted {
		t.Errorf("Expected GoalStatus %s, but got %s", NotStarted, result.GoalStatus)
	}
}

func TestConvertDeadline(t *testing.T) {
	g := Goal{}
	date := "2022-12-31"
	expectedTime, _ := time.Parse("2006-01-02", date)

	result, err := g.ConvertDeadline(date)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !result.Equal(expectedTime) {
		t.Errorf("Expected time %v, but got %v", expectedTime, result)
	}

	invalidDate := "invalid-date"
	_, err = g.ConvertDeadline(invalidDate)

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
