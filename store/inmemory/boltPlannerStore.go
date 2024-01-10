package inmemory

import (
	"github.com/asdine/storm"
	"goworkflow/models"
)

type BoltPlannerStore struct {
	db *storm.DB
}

func NewInMemoryPlannerStore(db *storm.DB) *BoltPlannerStore {
	return &BoltPlannerStore{
		db: db,
	}
}

func (s *BoltPlannerStore) CreatePlanner(planner *models.Planner) error {
	return s.db.Save(planner)
}

func (s *BoltPlannerStore) UpdatePlanner(planner *models.Planner) error {
	return s.db.Update(planner)
}

func (s *BoltPlannerStore) DeletePlanner(id string) error {
	planner := new(models.Planner)
	planner.Id = id
	return s.db.DeleteStruct(planner)
}

func (s *BoltPlannerStore) GetPlanner(id string) (*models.Planner, error) {
	planner := new(models.Planner)
	if err := s.db.One("ID", id, planner); err != nil {
		return nil, err
	}
	return planner, nil
}

func (s *BoltPlannerStore) ListPlanners() ([]*models.Planner, error) {
	var planners []*models.Planner
	if err := s.db.All(&planners); err != nil {
		return nil, err
	}
	return planners, nil
}
