package models

import (
	"testing"
	"time"
)

func TestVersionInfo(t *testing.T) {
	version := VersionInfo{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	t.Run("GetMajor", func(t *testing.T) {
		if version.GetMajor() != 1 {
			t.Errorf("Expected major version to be 1, but got %d", version.GetMajor())
		}
	})

	t.Run("GetMinor", func(t *testing.T) {
		if version.GetMinor() != 2 {
			t.Errorf("Expected minor version to be 2, but got %d", version.GetMinor())
		}
	})

	t.Run("GetPatch", func(t *testing.T) {
		if version.GetPatch() != 3 {
			t.Errorf("Expected patch version to be 3, but got %d", version.GetPatch())
		}
	})
}

func TestSnapshot(t *testing.T) {
	goal := &Goal{}
	goal.GenerateGoalInstance(
		"12123j;dsf3",
		"Test Objective",
		time.Now(),
	)
	var plans []Plan
	var tasks []Task
	snapshot := Snapshot{
		Goal:  goal,
		Plans: plans,
		Tasks: tasks,
	}

	t.Run("GetGoal", func(t *testing.T) {
		if snapshot.GetGoal() != goal {
			t.Errorf("Expected goal to be %v, but got %v", goal, snapshot.GetGoal())
		}
	})

}
