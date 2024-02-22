package models

import (
	"testing"
	"time"
)

func TestGeneratePlanInstance(t *testing.T) {
	p := createPlan()

	if p.Id != pid {
		t.Errorf("Expected Id %s, but got %s", pid, p.Id)
	}

	if p.PlanName != planName {
		t.Errorf("Expected PlanName %s, but got %s", planName, p.PlanName)
	}

	if p.PlanDescription != planDescription {
		t.Errorf("Expected PlanDescription %s, but got %s", planDescription, p.PlanDescription)
	}

	if p.PlanStatus != NotStarted {
		t.Errorf("Expected PlanStatus %s, but got %s", NotStarted, p.PlanStatus)
	}

	if len(p.Tasks) != 0 {
		t.Errorf("Expected Tasks to be empty, but got %v", p.Tasks)
	}

	if !p.PlanCreatedAt.Before(time.Now()) {
		t.Errorf("Expected PlanCreatedAt to be before time.Now(), but got %v", p.PlanCreatedAt)
	}
}

func TestConvertPlanDate(t *testing.T) {
	p := createPlan()
	dateStr := "2022-12-31"
	expectedTime, _ := time.Parse("2006-01-02", dateStr)

	result, err := p.ConvertPlanDate(dateStr)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !result.Equal(expectedTime) {
		t.Errorf("Expected time %v, but got %v", expectedTime, result)
	}

	invalidDate := "invalid-date"
	_, err = p.ConvertPlanDate(invalidDate)

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// utility for creating a plan
func createPlan() *Plan {
	p := Plan{}
	id := "test-id"
	planName := "test-plan"
	planDescription := "test-description"
	planDate := time.Now()
	planTime := time.Now()
	goal_id := "test-goal-id"

	return p.GeneratePlanInstance(id, planName, planDescription, planDate, planTime, goal_id)
}

// dummy data for testing
var (
	pid             = "test-id"
	planName        = "test-plan"
	planDescription = "test-description"
	planDate        = time.Now()
	planTime        = time.Now()
	goal_id         = "test-goal-id"
)
