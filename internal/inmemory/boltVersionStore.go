package inmemory

import (
	"errors"
	"flow/internal/models"
	"github.com/asdine/storm"
)

type BoltVersionStore struct {
	db *storm.DB
}

func NewBoltVersionStore(db *storm.DB) *BoltVersionStore {
	return &BoltVersionStore{
		db: db,
	}
}

func (s *BoltVersionStore) AddVersion(v *models.Version) error {
	return s.db.Save(v)
}

func (s *BoltVersionStore) UpdateVersion(goalId *models.EntityID, updateType *models.VersionInfo) error {
	v := &models.Version{}
	if err := s.db.One("ID", goalId, v); err != nil {
		return err
	}
	v.No = *updateType
	return s.db.Save(v)
}

func (s *BoltVersionStore) GetCurrentVersion(goalId *models.EntityID) (*models.Version, error) {
	v := &models.Version{}
	if err := s.db.One("GoalID", goalId, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *BoltVersionStore) ListVersions(goalId *models.EntityID) ([]*models.Version, error) {
	var versions []*models.Version
	if err := s.db.Find("GoalID", goalId, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

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