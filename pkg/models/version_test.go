package models

import (
	"testing"
	"time"
)

func TestVersionCreation(t *testing.T) {
	version := createVersion()

	if version.GoalID != "test-goalId" {
		t.Errorf("Expected GoalID test-goalId, but got %s", version.GoalID)
	}

	if version.PlanID != "test-planId" {
		t.Errorf("Expected PlanID test-planId, but got %s", version.PlanID)
	}

	if version.TaskID != "test-taskId" {
		t.Errorf("Expected TaskID test-taskId, but got %s", version.TaskID)
	}

	if version.No.Major != 0 {
		t.Errorf("Expected Major 0, but got %d", version.No.Major)
	}

	if version.No.Minor != 0 {
		t.Errorf("Expected Minor 0, but got %d", version.No.Minor)
	}

	if version.No.Patch != 0 {
		t.Errorf("Expected Patch 0, but got %d", version.No.Patch)
	}

	if version.CreatedBy != "test-userId" {
		t.Errorf("Expected CreatedBy test-userId, but got %s", version.CreatedBy)
	}
}

func TestVersionWithPreviousVersion(t *testing.T) {
	version := createVersion()
	previousVersion := createVersion()

	version.PreviousVersion = previousVersion

	if version.PreviousVersion != previousVersion {
		t.Errorf("Expected PreviousVersion to be %v, but got %v", previousVersion, version.PreviousVersion)
	}
}

// utility for creating a version
func createVersion() *Version {
	return &Version{
		GoalID:    "test-goalId",
		PlanID:    "test-planId",
		TaskID:    "test-taskId",
		No:        VersionInfo{Major: 0, Minor: 0, Patch: 0},
		CreatedAt: time.Now(),
		CreatedBy: "test-userId",
	}
}
