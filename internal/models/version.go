package models

import "time"

type EntityID string

type Version struct {
	GoalID          EntityID
	PlanID          EntityID
	TaskID          EntityID
	No              VersionInfo
	Image           Snapshot
	PreviousVersion *Version
	CreatedAt       time.Time
	CreatedBy       string
}

type VersionInfo struct {
	Major int
	Minor int
	Patch int
}

type Snapshot struct {
	Goal  *Goal
	Plans []Plan
	Tasks []Task
}
