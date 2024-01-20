package services

import (
	"flow/internal/models"
	store2 "flow/pkg/store"
)

// PlannerService is a type that provides operations for managing planners.
type PlannerService struct {
	store store2.PlannerStore
}

// NewPlannerService creates a new instance of the PlannerService.
// It takes a store of type PlannerStore as a parameter and returns a pointer to a PlannerService.
// The PlannerService struct has the following fields:
//   - store: PlannerStore (interface for accessing and manipulating planner data)
//
// Example usage:
//
//	plannerStore := inmemory.NewInMemoryPlannerStore(db)
//	plannerService := services.NewPlannerService(plannerStore)
//	newPlannerService := NewPlannerService(plannerStore)
func NewPlannerService(store store2.PlannerStore) *PlannerService {
	return &PlannerService{
		store: store,
	}
}

// CreatePlanner is a method of the PlannerService struct that creates a new planner with the provided data.
// It takes a pointer to a Planner struct as a parameter and returns an error.
// The method uses the CreatePlanner method of the PlannerStore interface to create the planner in the data store.
func (s *PlannerService) CreatePlanner(planner *models.Planner) error {
	return s.store.CreatePlanner(planner)
}

// UpdatePlanner updates an existing planner with the provided information.
// It takes a pointer to a models.Planner and returns an error.
// The planner parameter contains the following fields that can be updated:
// - Id: string (unique identifier for the planner)
// - Title: string (new title for the planner)
// - UserId: string (new user identifier associated with the planner)
// The method calls the UpdatePlanner method of the PlannerStore interface to update the planner in the database.
func (s *PlannerService) UpdatePlanner(planner *models.Planner) error {
	return s.store.UpdatePlanner(planner)
}

// DeletePlanner deletes a planner with the given ID.
// It takes a string parameter 'id' which is the unique identifier of the planner.
// The method calls the DeletePlanner method of the PlannerStore interface to delete the planner from the database.
func (s *PlannerService) DeletePlanner(id string) error {
	return s.store.DeletePlanner(id)
}

// GetPlanner retrieves a planner with the given ID.
// It takes a string argument 'id' representing the unique identifier of the planner to be retrieved.
// The method returns a pointer to a models.Planner and an error.
// The Planner structure has the following fields:
// - Id: string (unique identifier for the planner)
// - Title: string (title of the planner)
// - UserId: string (identifier of the user associated with the planner)
// - Goals: []Goal (list of goals associated with the planner)
// The method calls the GetPlanner method of the PlannerStore interface to fetch the planner from the database.
//
// Example usage:
//
//	func (c *PlannerControl) GetPlanner(req *GetPlannerRequest) (*GetPlannerResponse, error) {
//		planner, err := c.Service.GetPlanner(req.Id)
//		if err != nil {
//			return nil, err
//		}
//		return &GetPlannerResponse{
//			Id:     planner.Id,
//			UserId: planner.UserId,
//		}, nil
//	}
func (s *PlannerService) GetPlanner(id string) (*models.Planner, error) {
	return s.store.GetPlanner(id)
}

// ListPlanners returns a list of all planners.
// It takes no parameters and returns a slice of pointers to models.Planner and an error.
// The method calls the ListPlanners method of the PlannerStore interface to retrieve the list of planners from the database.
// Example usage:
// planners, err := c.service.ListPlanners()
//
//	if err != nil {
//	  return nil, err
//	}
//
//	for _, planner := range planners {
//	  fmt.Println(planner.Title)
//	}
func (s *PlannerService) ListPlanners() ([]*models.Planner, error) {
	return s.store.ListPlanners()
}

func (s *PlannerService) GetPlannerByTitle(title string) (*models.Planner, error) {
	return s.store.GetPlannerByTitle(title)
}

// GetPlannerByOwner retrieves all planners owned by the given user.
// It takes a string parameter "id" which represents the user identifier.
// It returns a slice of pointers to models.Planner and an error.
// Each element of the returned slice represents a planner and contains the following fields:
// - Id: string (unique identifier for the planner)
// - Title: string (title of the planner)
// - UserId: string (identifier of the user associated with the planner)
// - Goals: []Goal (list of goals associated with the planner)
// The method calls the GetPlannerByOwner method of the PlannerStore interface to retrieve the planners from the database.
func (s *PlannerService) GetPlannerByOwner(id string) ([]*models.Planner, error) {
	return s.store.GetPlannerByOwner(id)
}
