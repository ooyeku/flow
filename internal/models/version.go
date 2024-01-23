package models

type EntityID string

type VersionInfo struct {
	Major int // increment for major stuff
	Minor int // increment change number for changes to child models
	Patch int // increment for status/state changes
}

func (i VersionInfo) GetMajor() int {
	return i.Major
}

func (i VersionInfo) GetMinor() int {
	return i.Minor
}

func (i VersionInfo) GetPatch() int {
	return i.Patch
}

type Snapshot struct {
	Goal  *Goal
	Plans []Plan
	Tasks []Task
}

func (s Snapshot) GetGoal() *Goal {
	return s.Goal
}

func (s Snapshot) GetPlans() []Plan {
	return s.Plans
}

func (s Snapshot) GetTasks() []Task {
	return s.Tasks
}

type Version struct {
	GoalID EntityID
	No     VersionInfo
	Image  Snapshot
}
