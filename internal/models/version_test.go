package models

import (
	"testing"
	"time"
)

func TestVersionCreation(t *testing.T) {
	version := Version{
		GoalID:    "goal1",
		PlanID:    "plan1",
		TaskID:    "task1",
		No:        VersionInfo{Major: 1, Minor: 0, Patch: 0},
		CreatedAt: time.Now(),
		CreatedBy: "user1",
	}

	if version.GoalID != "goal1" {
		t.Errorf("Expected GoalID to be 'goal1', got %s", version.GoalID)
	}

	if version.PlanID != "plan1" {
		t.Errorf("Expected PlanID to be 'plan1', got %s", version.PlanID)
	}

	if version.TaskID != "task1" {
		t.Errorf("Expected TaskID to be 'task1', got %s", version.TaskID)
	}

	if version.No.Major != 1 || version.No.Minor != 0 || version.No.Patch != 0 {
		t.Errorf("Expected VersionInfo to be 1.0.0, got %d.%d.%d", version.No.Major, version.No.Minor, version.No.Patch)
	}

	if version.CreatedBy != "user1" {
		t.Errorf("Expected CreatedBy to be 'user1', got %s", version.CreatedBy)
	}
}

func TestVersionWithPreviousVersion(t *testing.T) {
	previousVersion := &Version{
		GoalID:    "goal1",
		PlanID:    "plan1",
		TaskID:    "task1",
		No:        VersionInfo{Major: 1, Minor: 0, Patch: 0},
		CreatedAt: time.Now(),
		CreatedBy: "user1",
	}

	version := Version{
		GoalID:          "goal2",
		PlanID:          "plan2",
		TaskID:          "task2",
		No:              VersionInfo{Major: 2, Minor: 0, Patch: 0},
		PreviousVersion: previousVersion,
		CreatedAt:       time.Now(),
		CreatedBy:       "user2",
	}

	if version.PreviousVersion == nil {
		t.Errorf("Expected PreviousVersion to not be nil")
	} else if version.PreviousVersion.GoalID != "goal1" {
		t.Errorf("Expected PreviousVersion GoalID to be 'goal1', got %s", version.PreviousVersion.GoalID)
	}
}
