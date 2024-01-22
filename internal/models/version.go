package models

// ID type for identifiers
type ID string

// Struct for the version info of the app
type VersionInfo struct {
	Major int // increment for major stuff
	Minor int // increment change number for changes to child models
	Patch int // increment for status/state changes
}

// Snapshot interface has GetData() method
type Snapshot interface {
	GetData() string // each model that implements GetData can be versioned into an an image
}

// SnapshotImpl struct implements Snapshot interface
type SnapshotImpl struct {
	data string
}

// Capture method implementation
func (s *SnapshotImpl) GetData() string {
	return s.data
}

// Version contains VersionInfo, PlannerID and GoalID
type Version struct {
	PlannerID ID
	No        VersionInfo
	GoalId    ID
	Image     []Snapshot
}

// Repository struct
type Repository struct {
	versions []Version
}
