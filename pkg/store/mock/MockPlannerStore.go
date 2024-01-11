package mock

import (
	"errors"
	"goworkflow/pkg/models"
	"sync"
)

type MockPlannerStore struct {
	Planner map[string]*models.Planner
	mu      sync.Mutex
}

func NewMockPlannerStore() *MockPlannerStore {
	return &MockPlannerStore{
		Planner: make(map[string]*models.Planner),
	}
}

func (s *MockPlannerStore) CreatePlanner(planner *models.Planner) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Planner[planner.Id]; exists {
		return errors.New("planner already exists")
	}
	s.Planner[planner.Id] = planner
	return nil
}

func (s *MockPlannerStore) GetPlanner(id string) (*models.Planner, error) {
	planner, exists := s.Planner[id]
	if !exists {
		return nil, errors.New("planner referenced by ID not found")
	}
	return planner, nil
}

func (s *MockPlannerStore) UpdatePlanner(planner *models.Planner) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Planner[planner.Id]; !exists {
		return errors.New("planner not found")
	}
	s.Planner[planner.Id] = planner
	return nil
}

func (s *MockPlannerStore) DeletePlanner(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Planner[id]; !exists {
		return errors.New("planner not found")
	}
	delete(s.Planner, id)
	return nil
}

func (s *MockPlannerStore) ListPlanners() ([]*models.Planner, error) {
	planners := make([]*models.Planner, 0, len(s.Planner))
	for _, planner := range s.Planner {
		planners = append(planners, planner)
	}
	return planners, nil
}
