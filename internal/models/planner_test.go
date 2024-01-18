package models

import (
	"testing"
)

func TestGeneratePlannerInstance(t *testing.T) {
	p := Planner{}
	id := "test-id"
	title := "test-title"
	userId := "test-userId"

	result := p.GeneratePlannerInstance(id, title, userId)

	if result.Id != id {
		t.Errorf("Expected Id %s, but got %s", id, result.Id)
	}

	if result.Title != title {
		t.Errorf("Expected Title %s, but got %s", title, result.Title)
	}

	if result.UserId != userId {
		t.Errorf("Expected UserId %s, but got %s", userId, result.UserId)
	}
}
