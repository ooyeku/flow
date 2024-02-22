package models

import "time"

type EntityID string

type Version struct {
	ID              EntityID    `json:"id" storm:"id,unique"`
	GoalID          EntityID    `json:"goal_id"`
	PlanID          EntityID    `json:"plan_id"`
	TaskID          EntityID    `json:"task_id"`
	No              VersionInfo `json:"version_no"`
	Image           Snapshot    `json:"image"`
	PreviousVersion *Version    `json:"previous_version"`
	CreatedAt       time.Time   `json:"created_at"`
	CreatedBy       string      `json:"created_by"`
}

type VersionInfo struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type Snapshot struct {
	Goal  *Goal   `json:"goal"`
	Plans []*Plan `json:"plans"`
	Tasks []*Task `json:"tasks"`
}
