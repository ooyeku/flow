package models

import (
	"testing"
)

func TestGeneratePlannerInstance(t *testing.T) {
	p := createPlanner()

	if p.Id != "test-id" {
		t.Errorf("Expected Id test-id, but got %s", p.Id)
	}

	if p.Title != "test-title" {
		t.Errorf("Expected Title test-title, but got %s", p.Title)
	}

	if p.UserId != "test-userId" {
		t.Errorf("Expected UserId test-userId, but got %s", p.UserId)
	}

	if len(p.Goals) != 0 {
		t.Errorf("Expected Goals to be empty, but got %v", p.Goals)
	}
}

// utility for creating a planner
func createPlanner() *Planner {
	return &Planner{
		Id:     "test-id",
		Title:  "test-title",
		UserId: "test-userId",
		Goals:  []Goal{},
	}
}
