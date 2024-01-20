package handle

import (
	"flow/internal/models"
	"flow/pkg/services"
	"github.com/google/uuid"
)

// PlanControl is a type that provides control operations for managing plans.
//
// The PlanControl type has a Service field of type *services.PlanService, which is used to interact with the PlanStore.
//
// Usage Example:
// pc := NewPlanControl(service)
//
// pc.CreatePlan(req *CreatePlanRequest) (*CreatePlanResponse, error)
// - Creates a new plan based on the provided CreatePlanRequest.
// - Returns a CreatePlanResponse containing the ID of the created plan, or an error if operation fails.
//
// pc.UpdatePlan(req *UpdatePlanRequest) error
// - Updates an existing plan based on the provided UpdatePlanRequest.
// - Returns an error if operation fails.
//
// pc.DeletePlan(req *DeletePlanRequest) error
// - Deletes a plan based on the provided DeletePlanRequest.
// - Returns an error if operation fails.
//
// pc.GetPlan(req *GetPlanRequest) (*GetPlanResponse, error)
// - Retrieves a plan based on the provided GetPlanRequest.
// - Returns a GetPlanResponse containing the retrieved plan, or an error if operation fails.
//
// pc.GetPlanByName(req *GetPlanByNameRequest) (*GetPlanByNameResponse, error)
// - Retrieves a plan based on the provided GetPlanByNameRequest.
// - Returns a GetPlanByNameResponse containing the retrieved plan, or an error if operation fails.
//
// pc.GetPlansByGoal(req *GetPlansByGoalRequest) (*GetPlansByGoalResponse, error)
// - Retrieves plans based on the provided GetPlansByGoalRequest.
// - Returns a GetPlansByGoalResponse containing the retrieved plans, or an error if operation fails.
//
// pc.ListPlans() (*ListPlansResponse, error)
// - Retrieves all plans.
// - Returns a ListPlansResponse containing the retrieved plans, or an error if operation fails.
type PlanControl struct {
	Service *services.PlanService
}

// NewPlanControl creates a new instance of PlanControl with the given PlanService.
func NewPlanControl(service *services.PlanService) *PlanControl {
	return &PlanControl{
		Service: service,
	}
}

// CreatePlanRequest is a type that represents a request to create a plan.
// It includes fields for the plan's name, description, date, and time.
type CreatePlanRequest struct {
	PlanName        string `json:"plan_name"`
	PlanDescription string `json:"plan_description"`
	PlanDate        string `json:"plan_date"`
	PlanTime        string `json:"plan_time"`
}

// CreatePlanResponse is a type representing the response of creating a plan.
// It contains the ID of the created plan.
type CreatePlanResponse struct {
	ID string `json:"id"`
}

// CreatePlan creates a new plan with the provided request. It generates a new UUID using the generatePlanUUID function.
// Then it creates a new instance of the models.Plan struct with the ID, PlanName, PlanDescription, PlanDate, and PlanTime values
// extracted from the request. It converts the PlanDate and PlanTime strings to time.Time values using the ConvertPlanDate and
// ConvertPlanTime methods of the models.Plan struct. Finally, it calls the CreatePlan method of the PlanService stored in the
// PlanControl struct, passing the newly created plan as the argument. It returns a CreatePlanResponse containing the ID of the
// created plan or an error if there was a problem creating the plan.
func (c *PlanControl) CreatePlan(req *CreatePlanRequest) (*CreatePlanResponse, error) {
	id, err := generatePlanUUID()
	if err != nil {
		return nil, err
	}
	m := &models.Plan{}
	// convert planDate to time.Time
	planDate, err := m.ConvertPlanDate(req.PlanDate)
	if err != nil {
		return nil, err
	}
	// convert planTime to time.Time
	planTime, err := m.ConvertPlanTime(req.PlanTime)
	if err != nil {
		return nil, err
	}
	plan := m.GeneratePlanInstance(id, req.PlanName, req.PlanDescription, planDate, planTime)
	err = c.Service.CreatePlan(plan)
	if err != nil {
		return nil, err
	}
	return &CreatePlanResponse{
		ID: plan.Id,
	}, nil
}

// UpdatePlanRequest represents a request object for updating a plan.
type UpdatePlanRequest struct {
	// updateplanrequest only updates the fields below.
	// updates to status and tasks are handled by other endpoints.
	Id              string `json:"id"`
	PlanName        string `json:"plan_name"`
	PlanDescription string `json:"plan_description"`
	PlanDate        string `json:"plan_date"`
	PlanTime        string `json:"plan_time"`
	GoalId          string `json:"goal_id"`
}

// UpdatePlan updates a plan with the provided request. It creates a new `models.Plan` struct
// with the `Id` and `GoalId` fields from the request. Then it calls the `UpdatePlan` method of
// the `PlanService` stored in the `PlanControl` struct, passing the newly created plan as the argument.
// It returns an error if there was a problem updating the plan.
func (c *PlanControl) UpdatePlan(req *UpdatePlanRequest) error {
	m := &models.Plan{}
	// convert planDate to time.Time
	planDate, err := m.ConvertPlanDate(req.PlanDate)
	if err != nil {
		return err
	}
	// convert planTime to time.Time
	planTime, err := m.ConvertPlanTime(req.PlanTime)
	if err != nil {
		return err
	}
	plan := m.GeneratePlanInstance(req.Id, req.PlanName, req.PlanDescription, planDate, planTime)
	plan.GoalId = req.GoalId
	return c.Service.UpdatePlan(plan)
}

// DeletePlanRequest is a type that represents a request to delete a plan.
// The `Id` field is used to specify the identifier of the plan to be deleted.
type DeletePlanRequest struct {
	Id string `json:"id"`
}

// DeletePlan deletes a plan with the specified ID by calling the DeletePlan method of the PlanService. It takes a DeletePlanRequest parameter which contains the ID of the plan to be
func (c *PlanControl) DeletePlan(req *DeletePlanRequest) error {
	return c.Service.DeletePlan(req.Id)
}

// GetPlanRequest is a type that represents a request to get a plan by its ID.
// It contains the following field:
// - Id: a string that specifies the ID of the plan to get.
//
// Example usage:
//
//	req := handle.GetPlanRequest{
//		Id: "12345",
//	}
//	plan, err := p.GetPlan(&req)
//	if err != nil {
//		fmt.Printf("Error getting plan with id %s: %s\n", req.Id, err)
//		return
//	}
//	fmt.Println("Got plan: ", plan.Plan.PlanName)
//	fmt.Println("Description: ", plan.Plan.PlanDescription)
//
// In this example, we create a GetPlanRequest instance with an ID and use it to get a plan by calling the GetPlan method on a PlanControl instance.
type GetPlanRequest struct {
	Id string `json:"id"`
}

// GetPlanResponse represents the response data structure for the GetPlan method in the PlanControl type.
// It contains the retrieved plan information.
type GetPlanResponse struct {
	Plan *models.Plan `json:"plan"`
}

// GetPlan retrieves a plan from the service using the provided plan ID. It calls the GetPlan method of the PlanService with the given ID and returns the retrieved plan in a GetPlan
func (c *PlanControl) GetPlan(req *GetPlanRequest) (*GetPlanResponse, error) {
	plan, err := c.Service.GetPlan(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetPlanResponse{
		Plan: plan,
	}, nil
}

// GetPlanByNameRequest is a type that represents a request to get a plan by its name.
// The request contains the name of the plan.
type GetPlanByNameRequest struct {
	PlanName string `json:"plan_name"`
}

// GetPlanByNameResponse is a type that represents the response
// for getting a plan by name.
// It contains the plan details.
type GetPlanByNameResponse struct {
	Plan *models.Plan `json:"plan"`
}

// GetPlanByName retrieves a plan by its name from the PlanService.
func (c *PlanControl) GetPlanByName(req *GetPlanByNameRequest) (*GetPlanByNameResponse, error) {
	plan, err := c.Service.GetPlanByName(req.PlanName)
	if err != nil {
		return nil, err
	}
	return &GetPlanByNameResponse{
		Plan: plan,
	}, nil
}

// GetPlansByGoalRequest represents a request to get plans by goal ID.
type GetPlansByGoalRequest struct {
	GoalId string `json:"goal_id"`
}

// GetPlansByGoalResponse is a type that represents the response for getting plans by goal.
// It contains a list of plans.
type GetPlansByGoalResponse struct {
	Plans []*models.Plan `json:"plans"`
}

// GetPlansByGoal retrieves all plans associated with a goal ID.
func (c *PlanControl) GetPlansByGoal(req *GetPlansByGoalRequest) (*GetPlansByGoalResponse, error) {
	plans, err := c.Service.GetPlansByGoal(req.GoalId)
	if err != nil {
		return nil, err
	}
	return &GetPlansByGoalResponse{
		Plans: plans,
	}, nil
}

// ListPlansResponse is a type that represents a response containing a list of plans.
type ListPlansResponse struct {
	Plans []*models.Plan `json:"plans"`
}

// ListPlans retrieves a list of plans by calling the ListPlans method of the PlanService.
func (c *PlanControl) ListPlans() (*ListPlansResponse, error) {
	plans, err := c.Service.ListPlans()
	if err != nil {
		return nil, err
	}
	return &ListPlansResponse{
		Plans: plans,
	}, nil
}

// generatePlanUUID generates a unique UUID for creating a plan.
//
// Example:
//
//	  id, err := generatePlanUUID()
//	  if err != nil {
//		   return nil, err
//	  }
//	  // Use id to create a new plan
func generatePlanUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
