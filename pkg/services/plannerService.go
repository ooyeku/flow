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

// GoalService is a type that provides operations for managing goals.
type GoalService struct {
	store store2.GoalStore
}

// NewGoalService creates a new GoalService with the given GoalStore.
// It returns a pointer to the GoalService.
func NewGoalService(store store2.GoalStore) *GoalService {
	return &GoalService{
		store: store,
	}
}

// CreateGoal creates a new goal with the given information.
// It takes a pointer to a models.Goal and returns an error.
// The goal parameter contains the following fields:
// - Id: string (unique identifier for the goal)
// - Name: string (name of the goal)
// - Description: string (description of the goal)
// - Tasks: []Task (list of tasks associated with the goal)
// The method calls the CreateGoal method of the GoalStore interface to store the goal in the database.
// Example usage:
//
//	id, err := generateUUID()
//	if err != nil {
//	  return nil, err
//	}
//	goal := &models.Goal{
//	  Id:          id,
//	  Name:        req.Name,
//	  Description: req.Description,
//	}
//	err = c.service.CreateGoal(goal)
//	if err != nil {
//	  return nil, err
//	}
//	return &CreateGoalResponse{
//	  ID: goal.Id,
//	}, nil
func (s *GoalService) CreateGoal(goal *models.Goal) error {
	return s.store.CreateGoal(goal)
}

// UpdateGoal is a method of the GoalService struct that updates a goal with the provided data.
// It takes a pointer to a Goal struct as a parameter and returns an error.
// The method uses the UpdateGoal method of the GoalStore interface to update the goal in the data store.
func (s *GoalService) UpdateGoal(goal *models.Goal) error {
	return s.store.UpdateGoal(goal)
}

// DeleteGoal deletes a goal with the specified ID.
func (s *GoalService) DeleteGoal(id string) error {
	return s.store.DeleteGoal(id)
}

// GetGoal retrieves a goal with the given ID from the GoalService.
// It returns the goal and an error if any occurred.
func (s *GoalService) GetGoal(id string) (*models.Goal, error) {
	return s.store.GetGoal(id)
}

// ListGoals is a method of GoalService that retrieves a list of goals and returns them along with any errors encountered.
// Signature:
// func (s *GoalService) ListGoals() ([]*models.Goal, error)
// Usage example:
// goals, err := c.service.ListGoals()
//
//	if err != nil {

//	    return nil, err
//	}
//
//	var goalResponses []*GetGoalResponse
//
//	for _, goal := range goals {
//	    goalResponses = append(goalResponses, &GetGoalResponse{
//	        Id:          goal.Id,

//	        Name:        goal.Name,
//	        Description: goal.Description,
//	    })
//	}
//
//	return &ListGoalsResponse{
//	    Goals: goalResponses,
//	}, nil
func (s *GoalService) ListGoals() ([]*models.Goal, error) {
	return s.store.ListGoals()
}

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
