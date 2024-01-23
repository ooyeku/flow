package inmemory

import (
	"errors"
	"flow/internal/models"
	"github.com/asdine/storm"
)

// BoltVersionStore is a type that represents a version store implemented using BoltDB.
// AddVersion adds a new version to the store.
//
// Example usage:
//
//	s := NewBoltVersionStore(db)
//	err := s.AddVersion(v)
type BoltVersionStore struct {
	db *storm.DB
}

// NewBoltVersionStore initializes a new instance of BoltVersionStore with the provided storm.DB database handle.
func NewBoltVersionStore(db *storm.DB) *BoltVersionStore {
	return &BoltVersionStore{
		db: db,
	}
}

// AddVersion adds a new version to the BoltVersionStore.
func (s *BoltVersionStore) AddVersion(v *models.Version) error {
	return s.db.Save(v)
}

// UpdateVersion updates the version of a goal in the BoltVersionStore.
//
// It takes the following parameters:
// - goalId: A pointer to a models.EntityID representing the ID of the goal.
// - updateType: A pointer to a models.VersionInfo representing the new version details.
//
// The function performs the following steps:
// 1. Retrieves the existing version of the goal from the database using the goalId.
// 2. Updates the version number of the retrieved version with the new updateType.
// 3. Saves the updated version back into the database.
//
// If any error occurs during the process, it is returned.
//
// Example usage:
// err := store.UpdateVersion(&goalId, &updateType)
//
// Note: The BoltVersionStore instance (store) must be initialized and the database connection (store.db) must be set before calling this method.
func (s *BoltVersionStore) UpdateVersion(goalId *models.EntityID, updateType *models.VersionInfo) error {
	v := &models.Version{}
	if err := s.db.One("ID", goalId, v); err != nil {
		return err
	}
	v.No = *updateType
	return s.db.Save(v)
}

// GetCurrentVersion retrieves the current version of a goal by its ID from the BoltVersionStore.
// It returns the current version of the goal and an error indicating any issues with the retrieval.
func (s *BoltVersionStore) GetCurrentVersion(goalId *models.EntityID) (*models.Version, error) {
	v := &models.Version{}
	if err := s.db.One("GoalID", goalId, v); err != nil {
		return nil, err
	}
	return v, nil
}

// ListVersions returns a list of versions associated with a specific goal ID.
func (s *BoltVersionStore) ListVersions(goalId *models.EntityID) ([]*models.Version, error) {
	var versions []*models.Version
	if err := s.db.Find("GoalID", goalId, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

// GetImage retrieves the image of a specific version of a goal from the BoltVersionStore.
// It takes the goalId and versionNo as parameters and returns the Snapshot and any error encountered.
//
// The goalId is the identifier of the goal.
// The versionNo specifies the version of the goal.
// The Snapshot contains the goal, plans, and tasks associated with the specified version.
//
// If the goal with the specified goalId does not exist, it returns an error.
// If the version specified by versionNo does not exist, it returns an error.
//
// Example usage:
//
//	goalId := models.EntityID("example_goal_id")
//	versionNo := &models.VersionInfo{Major: 1, Minor: 0, Patch: 0}
//	snapshot, err := store.GetImage(&goalId, versionNo)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(snapshot)
func (s *BoltVersionStore) GetImage(goalId *models.EntityID, versionNo *models.VersionInfo) (*models.Snapshot, error) {
	v := &models.Version{}
	if err := s.db.One("GoalID", goalId, v); err != nil {
		return nil, err
	}
	if v.No != *versionNo {
		return nil, errors.New("version not found")
	}
	return &v.Image, nil
}
