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

// NewPlanControl creates a new instance of PlanControl with the provided PlanService.
func NewPlanControl(service *services.PlanService) *PlanControl {
	return &PlanControl{
		Service: service,
	}
}

// CreatePlanRequest is a type that represents the request to create a new plan.
// The CreatePlanRequest type has the following fields:
// - PlanName: the name of the plan.
// - PlanDescription: the description of the plan.
// - PlanDate: the date of the plan in the format "YYYY-MM-DD".
// - PlanTime: the time of the plan in the format "HH:MM".
// Example usage:
//
//	req := &CreatePlanRequest{
//	    PlanName:        "My Plan",
//	    PlanDescription: "This is a test plan",
//	    PlanDate:        "2022-01-01",
//	    PlanTime:        "12:00",
//	}
//
// res, err := planControl.CreatePlan(req)
type CreatePlanRequest struct {
	PlanName        string `json:"plan_name"`
	PlanDescription string `json:"plan_description"`
	PlanDate        string `json:"plan_date"`
	PlanTime        string `json:"plan_time"`
	GoalId          string
}

// CreatePlanResponse is a type that represents the response when creating a new plan.
// The CreatePlanResponse type has an ID field of type string, which stores the ID of the created plan.
// Usage Example:
//
//	response := &CreatePlanResponse{
//	    ID: "abcd1234",
//	}
type CreatePlanResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// CreatePlan creates a new plan with the provided request. It generates a unique id using the `generatePlanUUID` function.
// Then it creates a new `models.Plan` struct with the `Id`, `PlanName`, `PlanDescription`, `planDate`, and `planTime` fields
// from the request. After that, it calls the `CreatePlan` method of the `PlanService` stored in the `PlanControl` struct,
// passing the newly created plan as the argument. If the creation is successful, it returns a `CreatePlanResponse` with the created plan ID.
// It returns an error if there was a problem generating the unique id, converting the plan date or time, or creating the plan.
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
	plan := m.GeneratePlanInstance(id, req.PlanName, req.PlanDescription, planDate, planTime, req.GoalId)
	err = c.Service.CreatePlan(plan)
	if err != nil {
		return nil, err
	}
	return &CreatePlanResponse{
		ID: plan.Id,
	}, nil
}

// convert planDate to time.Time
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

// UpdatePlan updates an existing plan with the provided request. It converts the PlanDate and PlanTime strings
// to time.Time values using the ConvertPlanDate and ConvertPlanTime methods of the models.Plan struct.
// Then it creates a new instance of the models.Plan struct with the provided ID, PlanName, PlanDescription,
// GoalId, PlanDate, and PlanTime values.
// It updates the GoalId field of the plan instance with the GoalId value from the request.
// Finally, it calls the UpdatePlan method of the PlanService stored in the PlanControl struct, passing the updated plan as the argument.
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
	plan := m.GeneratePlanInstance(req.Id, req.PlanName, req.PlanDescription, planDate, planTime, req.GoalId)
	plan.GoalId = req.GoalId
	return c.Service.UpdatePlan(plan)
}

// DeletePlanRequest is a type that represents a request to delete a plan.
// The DeletePlanRequest type has a field Id of type string, which represents the ID of the plan to be deleted.
// Usage Example:
//
//	req := DeletePlanRequest{
//	    Id: "example-id",
//	}
type DeletePlanRequest struct {
	Id string `json:"id"`
}

// DeletePlan deletes the plan with the specified ID by calling the DeletePlan method of the PlanService stored in the PlanControl struct and passing the ID as the argument.
func (c *PlanControl) DeletePlan(req *DeletePlanRequest) error {
	return c.Service.DeletePlan(req.Id)
}

// GetPlanRequest is a type used to request the retrieval of a plan based on its ID.
// Fields:
// - Id: a string that represents the ID of the plan to be retrieved.
// Usage Example:
//
//	req := handle.GetPlanRequest{
//	  Id: "12345",
//	}
//
// plan, err := p.GetPlan(&req)
//
//	if err != nil {
//	  fmt.Printf("Error getting plan with id %s: %s\n", req.Id, err)
//	  return
//	}
//
// fmt.Println("Got plan: ", plan.Plan.PlanName)
// fmt.Println("Description: ", plan.Plan.PlanDescription)
type GetPlanRequest struct {
	Id string `json:"id"`
}

// GetPlanResponse is a type that represents the response of retrieving a plan.
// GetPlanResponse has a Plan field of type *models.Plan, which contains information about the retrieved plan.
// Usage Example:
//
//	func (c *PlanControl) GetPlan(req *GetPlanRequest) (*GetPlanResponse, error) {
//	  plan, err := c.Service.GetPlan(req.Id)
//	  if err != nil {
//	      return nil, err
//	  }
//	  return &GetPlanResponse{
//	      Plan: plan,
//	  }, nil
//	}
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

// GetPlanByNameRequest is a type that represents a request to retrieve a plan by its name.
//
// The GetPlanByNameRequest type has a PlanName field of type string, which specifies the name of the plan to retrieve.
type GetPlanByNameRequest struct {
	PlanName string `json:"plan_name"`
}

// GetPlanByNameResponse is a type that represents the response of retrieving a plan by its name.
// GetPlanByNameResponse has a Plan field of type *models.Plan, which represents the retrieved plan.
type GetPlanByNameResponse struct {
	Plan *models.Plan `json:"plan"`
}

// GetPlanByName retrieves a plan by its name from the PlanService. It calls the GetPlanByName method of the PlanService
// stored in the PlanControl struct, passing the PlanName from the request as the argument. It returns a GetPlanByNameResponse
// containing the retrieved plan or an error if there was a problem getting the plan.
func (c *PlanControl) GetPlanByName(req *GetPlanByNameRequest) (*GetPlanByNameResponse, error) {
	plan, err := c.Service.GetPlanByName(req.PlanName)
	if err != nil {
		return nil, err
	}
	return &GetPlanByNameResponse{
		Plan: plan,
	}, nil
}

// GetPlansByGoalRequest is a type that represents a request to retrieve plans by goal ID.
//
// The GetPlansByGoalRequest type has a GoalId field, which is the ID of the goal for which to retrieve plans.
type GetPlansByGoalRequest struct {
	GoalId string `json:"goal_id"`
}

// GetPlansByGoalResponse is a type that represents the response for retrieving plans by goal.
// The GetPlansByGoalResponse type has a Plans field which is a slice of *models.Plan.
type GetPlansByGoalResponse struct {
	Plans []*models.Plan `json:"plans"`
}

// GetPlansByGoal retrieves plans based on the provided goal ID. It calls the GetPlansByGoal method of the PlanService stored in the PlanControl struct, passing the goal ID as the argument.
// It returns a GetPlansByGoalResponse containing the retrieved plans or an error if there was a problem retrieving the plans.
func (c *PlanControl) GetPlansByGoal(req *GetPlansByGoalRequest) (*GetPlansByGoalResponse, error) {
	plans, err := c.Service.GetPlansByGoal(req.GoalId)
	if err != nil {
		return nil, err
	}
	return &GetPlansByGoalResponse{
		Plans: plans,
	}, nil
}

// ListPlansResponse is a type that represents the response of retrieving all plans.
// The ListPlansResponse type has a Plans field of type []*models.Plan, which contains the retrieved plans.
// Usage Example:
//
//	func (c *PlanControl) ListPlans() (*ListPlansResponse, error) {
//	  plans, err := c.Service.ListPlans()
//	  if err != nil {
//	    return nil, err
//	  }
//	  return &ListPlansResponse{
//	    Plans: plans,
//	  }, nil
//	}
type ListPlansResponse struct {
	Plans []*models.Plan `json:"plans"`
}

// ListPlans returns a ListPlansResponse containing a list of plans retrieved from the PlanService.
// It calls the ListPlans method of the PlanService stored in the PlanControl struct.
// If there is an error retrieving the plans, it returns nil and the error.
// Otherwise, it returns the ListPlansResponse with the plans.
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
// It uses the uuid.NewRandom function from the "github.com/google/uuid" package to generate a random UUID.
// If an error occurs during the generation of the UUID, the function returns an empty string and the error.
// Otherwise, it returns the generated UUID as a string and nil error.
func generatePlanUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
