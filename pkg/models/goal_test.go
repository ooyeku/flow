package models

import (
	"testing"
	"time"
)

func TestGenerateGoalInstance(t *testing.T) {
	g := createGoal()

	if g.Id != gid {
		t.Errorf("Expected Id %s, but got %s", gid, g.Id)

	}

	if g.Objective != objective {
		t.Errorf("Expected Objective %s, but got %s", objective, g.Objective)
	}

	if g.GoalStatus != NotStarted {
		t.Errorf("Expected GoalStatus %s, but got %s", NotStarted, g.GoalStatus)
	}

	if !g.GoalCreatedAt.Before(time.Now()) {
		t.Errorf("Expected GoalCreatedAt to be before time.Now(), but got %v", g.GoalCreatedAt)
	}

	if !g.GoalUpdatedAt.Before(time.Now()) {
		t.Errorf("Expected GoalUpdatedAt to be before time.Now(), but got %v", g.GoalUpdatedAt)
	}

	if !g.Deadline.Before(time.Now()) {
		t.Errorf("Expected Deadline to be before time.Now(), but got %v", g.Deadline)
	}

	if g.PlannerId != "test-plannerId" {
		t.Errorf("Expected PlannerId to be 'test-plannerId', but got %s", g.PlannerId)
	}
}

func TestConvertDeadline(t *testing.T) {
	g := Goal{}
	date := "2022-12-31"
	expectedTime, _ := time.Parse("2006-01-02", date)

	result, err := g.ConvertDeadtime(date)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !result.Equal(expectedTime) {
		t.Errorf("Expected time %v, but got %v", expectedTime, result)
	}

	invalidDate := "invalid-date"
	_, err = g.ConvertDeadtime(invalidDate)

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// utility for creating a goal
func createGoal() *Goal {
	return &Goal{
		Id:            "test-id",
		Objective:     "test-objective",
		GoalStatus:    NotStarted,
		GoalCreatedAt: time.Now(),
		GoalUpdatedAt: time.Now(),
		Deadline:      time.Now(),
		PlannerId:     "test-plannerId",
	}
}

// dummy data for testing
var (
	gid        = "test-id"
	objective  = "test-objective"
	deadline   = time.Now()
	id2        = "test-id2"
	objective2 = "test-objective2"
	deadline2  = time.Now()
	id3        = "test-id3"
	objective3 = "test-objective3"
	deadline3  = time.Now()
	id4        = "test-id4"
	objective4 = "test-objective4"
	deadline4  = time.Now()
	id5        = "test-id5"
	objective5 = "test-objective5"
	deadline5  = time.Now()
	id6        = "test-id6"
)
