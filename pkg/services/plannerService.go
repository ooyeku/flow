package services

import (
	"goworkflow/internal/models"
	store2 "goworkflow/pkg/store"
)

// PlannerService is a type that provides operations for managing planners.
type PlannerService struct {
	store store2.PlannerStore
}

// NewPlannerService creates a new PlannerService with the given PlannerStore.
// It returns a pointer to the PlannerService.
func NewPlannerService(store store2.PlannerStore) *PlannerService {
	return &PlannerService{
		store: store,
	}
}

// CreatePlanner creates a new planner with the given information.
// It takes a pointer to a models.Planner and returns an error.
// The planner parameter contains the following fields:
// - Id: string (unique identifier for the planner)
// - UserId: string (identifier of the user associated with the planner)
// - Goals: []Goal (list of goals associated with the planner)
// The method calls the CreatePlanner method of the PlannerStore interface to store the planner in the database.
// Example usage:
//
//	id, err := generateUUID()
//	if err != nil {
//	  return nil, err
//	}
//	planner := &models.Planner{
//	  Id:     id,
//	  UserId: req.UserId,
//	}
//	err = c.service.CreatePlanner(planner)
//	if err != nil {
//	  return nil, err
//	}
//	return &CreatePlannerResponse{
//	  ID: planner.Id,
//	}, nil
func (s *PlannerService) CreatePlanner(planner *models.Planner) error {
	return s.store.CreatePlanner(planner)
}

// UpdatePlanner is a method of the PlannerService struct that updates a planner with the provided data.
// It takes a pointer to a Planner struct as a parameter and returns an error.
// The method uses the UpdatePlanner method of the PlannerStore interface to update the planner in the data store.
func (s *PlannerService) UpdatePlanner(planner *models.Planner) error {
	return s.store.UpdatePlanner(planner)
}

// DeletePlanner deletes a planner with the specified ID.
func (s *PlannerService) DeletePlanner(id string) error {
	return s.store.DeletePlanner(id)
}

// GetPlanner retrieves a planner with the given ID from the PlannerService.
// It returns the planner and an error if any occurred.
func (s *PlannerService) GetPlanner(id string) (*models.Planner, error) {
	return s.store.GetPlanner(id)
}

// ListPlanners is a method of PlannerService that retrieves a list of planners and returns them along with any errors encountered.
// Signature:
// func (s *PlannerService) ListPlanners() ([]*models.Planner, error)
// Usage example:
// planners, err := c.service.ListPlanners()
//
//	if err != nil {
//	    return nil, err
//	}
//
// var plannerResponses []*GetPlannerResponse
//
//	for _, planner := range planners {
//	    plannerResponses = append(plannerResponses, &GetPlannerResponse{
//	        Id:     planner.Id,
//	        UserId: planner.UserId,
//	    })
//	}
//
//	return &ListPlannersResponse{
//	    Planners: plannerResponses,
//	}, nil
func (s *PlannerService) ListPlanners() ([]*models.Planner, error) {
	return s.store.ListPlanners()
}
