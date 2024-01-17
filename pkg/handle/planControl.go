package handle

import (
	"flow/internal/models"
	"flow/pkg/services"
	"github.com/google/uuid"
)

// PlanControl is a type that provides control operations for managing plans.
type PlanControl struct {
	Service *services.PlanService
}

// NewPlanControl creates a new instance of PlanControl with the given PlanService.
func NewPlanControl(service *services.PlanService) *PlanControl {
	return &PlanControl{
		Service: service,
	}
}

// CreatePlanRequest represents the request body for creating a plan.
type CreatePlanRequest struct {
	GoalId string `json:"goal_id"`
}

// CreatePlanResponse represents the response object for creating a plan.
type CreatePlanResponse struct {
	ID string `json:"id"`
}

// CreatePlan generates a new plan by creating a unique ID using generatePlanUUID function. It initializes a new Plan object with the generated ID and the goal ID from the request. Then
func (c *PlanControl) CreatePlan(req *CreatePlanRequest) (*CreatePlanResponse, error) {
	id, err := generatePlanUUID()
	if err != nil {
		return nil, err
	}
	plan := &models.Plan{
		Id:     id,
		GoalId: req.GoalId,
	}
	err = c.Service.CreatePlan(plan)
	if err != nil {
		return nil, err
	}
	return &CreatePlanResponse{
		ID: plan.Id,
	}, nil
}

// UpdatePlanRequest represents a request to update a plan with new data.
type UpdatePlanRequest struct {
	Id     string `json:"id"`
	GoalId string `json:"goal_id"`
}

// UpdatePlan updates a plan with the provided request. It creates a new `models.Plan` struct
// with the `Id` and `GoalId` fields from the request. Then it calls the `UpdatePlan` method of
// the `PlanService` stored in the `PlanControl` struct, passing the newly created plan as the argument.
// It returns an error if there was a problem updating the plan.
func (c *PlanControl) UpdatePlan(req *UpdatePlanRequest) error {
	plan := &models.Plan{
		Id:     req.Id,
		GoalId: req.GoalId,
	}
	return c.Service.UpdatePlan(plan)
}

// DeletePlanRequest represents a request to delete a plan by its ID.
type DeletePlanRequest struct {
	Id string `json:"id"`
}

// DeletePlan deletes a plan based on the provided DeletePlanRequest.
// It calls the DeletePlan method of the PlanService to delete the plan from the store.
// It returns an error if the deletion fails.
func (c *PlanControl) DeletePlan(req *DeletePlanRequest) error {
	return c.Service.DeletePlan(req.Id)
}

// GetPlanRequest represents a request to retrieve a plan by its ID.
// The ID field specifies the unique identifier of the requested plan.
type GetPlanRequest struct {
	Id string `json:"id"`
}

// GetPlanResponse represents the response structure for the GetPlan API endpoint.
type GetPlanResponse struct {
	Plan *models.Plan `json:"plan"`
}

// GetPlan retrieves a specific plan based on the provided request ID.
// It returns a GetPlanResponse object that contains the requested plan,
// or an error if the plan could not be retrieved.
func (c *PlanControl) GetPlan(req *GetPlanRequest) (*GetPlanResponse, error) {
	plan, err := c.Service.GetPlan(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetPlanResponse{
		Plan: plan,
	}, nil
}

// ListPlansResponse represents the response structure for listing plans.
type ListPlansResponse struct {
	Plans []*models.Plan `json:"plans"`
}

// ListPlans returns a list of plans by calling the `ListPlans` method on the `PlanService` in `PlanControl`.
//
// Sample usage:
// plans, err := r.ListPlans()
//
//	if err != nil {
//	    log.Fatalf("Failed to list plans: %v", err)
//	}
//
// planJson, err := json.Marshal(plans)
//
//	if err != nil {
//	    log.Fatalf("Failed to marshal plans: %v", err)
//	}
//
// log.Printf("Plans: %s", string(planJson))
// return string(planJson)
func (c *PlanControl) ListPlans() (*ListPlansResponse, error) {
	plans, err := c.Service.ListPlans()
	if err != nil {
		return nil, err
	}
	return &ListPlansResponse{
		Plans: plans,
	}, nil
}

// generatePlanUUID generates a new UUID for a plan.
func generatePlanUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
