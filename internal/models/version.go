package models

type EntityID string

type VersionInfo struct {
	Major int // increment for major stuff
	Minor int // increment change number for changes to child models
	Patch int // increment for status/state changes
}

type Snapshot struct {
	Goal  *Goal
	Plans []Plan
	Tasks []Task
}

type Version struct {
	GoalID EntityID
	No     VersionInfo
	Image  Snapshot
}
