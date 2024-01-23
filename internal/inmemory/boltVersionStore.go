package inmemory

import (
	"flow/internal/models"
	"github.com/asdine/storm"
)

type InMemoryVersionStore struct {
	db     *storm.DB
	goalId *models.EntityID
}

func NewInMemoryVersionStore(db *storm.DB) models.VersionStore {
	return &InMemoryVersionStore{
		db: db,
	}
}

func (s *InMemoryVersionStore) SetGoalId(goalId *models.EntityID) {
	s.goalId = goalId
}

func (s *InMemoryVersionStore) AddVersion(v *models.Version) error {
	return s.db.Save(v)
}

func (s *InMemoryVersionStore) UpdateVersion(updateType *models.VersionInfo) error {
	// Use s.goalId here
}

func (s *InMemoryVersionStore) GetCurrentVersion() (*models.Version, error) {
	// Use s.goalId here
}

func (s *InMemoryVersionStore) ListVersions() ([]*models.Version, error) {
	// Use s.goalId here
}

func (s *InMemoryVersionStore) GetImage(versionNo *models.VersionInfo) (*models.Image, error) {
	// Use s.goalId here
}
