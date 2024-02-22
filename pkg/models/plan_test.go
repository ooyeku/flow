package models

import (
	"testing"
	"time"
)

func TestGeneratePlanInstance(t *testing.T) {
	p := Plan{}
	id := "test-id"
	planName := "test-plan"
	planDescription := "test-description"
	planDate := time.Now()
	planTime := time.Now()
	goal_id := "test-goal-id"

	result := p.GeneratePlanInstance(id, planName, planDescription, planDate, planTime, goal_id)

	if result.Id != id {
		t.Errorf("Expected Id %s, but got %s", id, result.Id)
	}

	if result.PlanName != planName {
		t.Errorf("Expected PlanName %s, but got %s", planName, result.PlanName)
	}

	if result.PlanDescription != planDescription {
		t.Errorf("Expected PlanDescription %s, but got %s", planDescription, result.PlanDescription)
	}

	if !result.PlanDate.Equal(planDate) {
		t.Errorf("Expected PlanDate %v, but got %v", planDate, result.PlanDate)
	}

	if !result.PlanTime.Equal(planTime) {
		t.Errorf("Expected PlanTime %v, but got %v", planTime, result.PlanTime)
	}

	if result.PlanStatus != NotStarted {
		t.Errorf("Expected PlanStatus %s, but got %s", NotStarted, result.PlanStatus)
	}
}

func TestConvertPlanDate(t *testing.T) {
	p := Plan{}
	date := "2022-12-31"
	expectedTime, _ := time.Parse("2006-01-02", date)

	result, err := p.ConvertPlanDate(date)

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

func TestConvertPlanTime(t *testing.T) {
	p := Plan{}
	timeStr := "15:04"
	expectedTime, _ := time.Parse("15:04", timeStr)

	result, err := p.ConvertPlanTime(timeStr)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !result.Equal(expectedTime) {
		t.Errorf("Expected time %v, but got %v", expectedTime, result)
	}

	invalidTime := "invalid-time"
	_, err = p.ConvertPlanTime(invalidTime)

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
