package services

import (
	"flow/internal/models"
	store2 "flow/pkg/store"
)

// PlanService is a type that provides operations for managing plans.
type PlanService struct {
	store store2.PlanStore
}

// NewPlanService creates a new PlanService with the given PlanStore.
// It returns a pointer to the PlanService.
func NewPlanService(store store2.PlanStore) *PlanService {
	return &PlanService{
		store: store,
	}
}

// CreatePlan creates a new plan with the given information.
// It takes a pointer to a models.Plan and returns an error.
// The plan parameter contains the following fields:
// - Id: string (unique identifier for the plan)
// - Name: string (name of the plan)
// - Description: string (description of the plan)
// - Tasks: []Task (list of tasks associated with the plan)
// The method calls the CreatePlan method of the PlanStore interface to store the plan in the database.
// Example usage:
//
//	id, err := generateUUID()

//	if err != nil {
//	    return nil, err
//	}
//
//	plan := &models.Plan{
//	    Id:          id,
//	    Name:        req.Name,
//	    Description: req.Description,
//	}
//
// err = c.service.CreatePlan(plan)
//
//	if err != nil {
//	    return nil, err
//	}
//
//	return &CreatePlanResponse{
//	    ID: plan.Id,
//	}, nil
func (s *PlanService) CreatePlan(plan *models.Plan) error {
	return s.store.CreatePlan(plan)
}

// UpdatePlan is a method of the PlanService struct that updates a plan with the provided data.
// It takes a pointer to a Plan struct as a parameter and returns an error.
// The method uses the UpdatePlan method of the PlanStore interface to update the plan in the data store.
func (s *PlanService) UpdatePlan(plan *models.Plan) error {
	return s.store.UpdatePlan(plan)
}

// DeletePlan deletes a plan with the specified ID.
func (s *PlanService) DeletePlan(id string) error {
	return s.store.DeletePlan(id)
}

// GetPlan retrieves a plan with the given ID from the PlanService.
// It returns the plan and an error if any occurred.
func (s *PlanService) GetPlan(id string) (*models.Plan, error) {
	return s.store.GetPlan(id)
}

// ListPlans is a method of PlanService that retrieves a list of plans and returns them along with any errors encountered.
// Signature:
// func (s *PlanService) ListPlans() ([]*models.Plan, error)
// Usage example:
// plans, err := c.service.ListPlans()
//
//	if err != nil {
//	    return nil, err
//	}
//
//	var planResponses []*GetPlanResponse
//
//	for _, plan := range plans {
//	    planResponses = append(planResponses, &GetPlanResponse{
//	        Id:          plan.Id,
//	        Name:        plan.Name,
//	        Description: plan.Description,
//	    })
//	}
//
//	return &ListPlansResponse{
//	    Plans: planResponses,
//	}, nil
func (s *PlanService) ListPlans() ([]*models.Plan, error) {
	return s.store.ListPlans()
}
