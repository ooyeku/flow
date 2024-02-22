package inmemory

import (
	"errors"
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/pkg/models"
)

type BoltVersionStore struct {
	db *storm.DB
}

func NewInMemoryVersionStore(db *storm.DB) *BoltVersionStore {
	return &BoltVersionStore{
		db: db,
	}
}

func (s *BoltVersionStore) CreateVersion(v *models.Version) error {
	return s.db.Save(v)
}

func (s *BoltVersionStore) UpdateVersion(id string, v *models.Version) error {
	existingVersion := &models.Version{}
	if err := s.db.One("ID", id, existingVersion); err != nil {
		return err
	}
	*existingVersion = *v
	return s.db.Save(existingVersion)
}

func (s *BoltVersionStore) GetVersion(id string) (*models.Version, error) {
	v := &models.Version{}
	if err := s.db.One("ID", id, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *BoltVersionStore) ListVersions() ([]*models.Version, error) {
	var versions []*models.Version
	if err := s.db.All(&versions); err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *BoltVersionStore) GetPreviousVersion(id string) (*models.Version, error) {
	v := &models.Version{}
	if err := s.db.One("PreviousVersion.ID", id, v); err != nil {
		return nil, err
	}
	return v.PreviousVersion, nil
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
