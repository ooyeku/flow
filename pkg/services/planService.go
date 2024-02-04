package services

import (
	"github.com/ooyeku/flow/internal/models"
	store2 "github.com/ooyeku/flow/pkg/store"
)

// PlanService is a type that provides operations for managing plans.
type PlanService struct {
	store store2.PlanStore
}

// NewPlanService initializes a new instance of the PlanService struct.
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

// CreatePlan is a method of the PlanService struct that creates a new plan with the provided data.
// It takes a pointer to a Plan struct as a parameter and returns an error.
// The method uses the CreatePlan method of the PlanStore interface to create the plan in the data store.
func (s *PlanService) CreatePlan(plan *models.Plan) error {
	return s.store.CreatePlan(plan)
}

// UpdatePlan updates the details of a plan.
// Parameters:
// - plan: a pointer to a Plan object representing the updated plan.
// Returns:
// - error: an error, if any.
func (s *PlanService) UpdatePlan(plan *models.Plan) error {
	return s.store.UpdatePlan(plan)
}

// DeletePlan is a method of the PlanService struct that deletes a plan from the store based on the provided ID.
// It calls the DeletePlan method of the PlanStore interface using the provided ID as the parameter.
// The method returns an error if there was a problem deleting the plan.
func (s *PlanService) DeletePlan(id string) error {
	return s.store.DeletePlan(id)
}

// GetPlan returns the plan with the specified ID.
func (s *PlanService) GetPlan(id string) (*models.Plan, error) {
	return s.store.GetPlan(id)
}

// ListPlans returns a list of plans from the PlanStore.
// It calls the ListPlans method of the store to fetch the list of plans.
// It returns a slice of Plan pointers and an error.
func (s *PlanService) ListPlans() ([]*models.Plan, error) {
	return s.store.ListPlans()
}

// GetPlanByName retrieves a plan by its name from the plan store.
// It returns a pointer to the plan and an error if one occurs.
func (s *PlanService) GetPlanByName(name string) (*models.Plan, error) {
	return s.store.GetPlanByName(name)
}

// GetPlansByGoal retrieves a list of plans associated with a specific goal.
//
// Parameters:
// - id: The ID of the goal.
//
// Returns:
// - plans: A list of plans that are associated with the specified goal.
// - error: Any error that occurred during the retrieval process.
func (s *PlanService) GetPlansByGoal(id string) ([]*models.Plan, error) {
	return s.store.GetPlansByGoal(id)
}
