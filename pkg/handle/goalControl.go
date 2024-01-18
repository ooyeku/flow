package handle

import (
	"flow/internal/models"
	"flow/pkg/services"
	"github.com/google/uuid"
)

// GoalControl represents a controller that provides methods to manage goals.
type GoalControl struct {
	Service *services.GoalService
}

// NewGoalControl creates a new GoalControl with the given GoalService.
// It returns a pointer to the GoalControl.
func NewGoalControl(service *services.GoalService) *GoalControl {
	return &GoalControl{
		Service: service,
	}
}

// CreateGoalRequest represents a request to create a goal.
// PlannerId is the ID of the planner for which the goal is being created.
type CreateGoalRequest struct {
	Objective string `json:"objective"`
	Deadline  string `json:"deadline"`
	PlannerId string `json:"planner_id"`
}

// CreateGoalResponse represents the response for creating a goal.
// It contains the ID of the newly created goal.
type CreateGoalResponse struct {
	ID string `json:"id"`
}

// CreateGoal creates a new goal for a planner based on the provided request.
func (c *GoalControl) CreateGoal(req *CreateGoalRequest) (*CreateGoalResponse, error) {
	id, err := generateGoalUUID()
	if err != nil {
		return nil, err
	}
	m := &models.Goal{}
	// convert deadline to time.Time
	deadline, err := m.ConvertDeadline(req.Deadline)
	goal := m.GenerateGoalInstance(id, req.Objective, deadline)
	err = c.Service.CreateGoal(goal)
	if err != nil {
		return nil, err
	}
	return &CreateGoalResponse{
		ID: goal.Id,
	}, nil
}

// UpdateGoalRequest represents the data required to update a goal.
// The Id field represents the identifier of the goal to be updated.
// The PlannerId field represents the new planner ID to be assigned to the goal after the update.
type UpdateGoalRequest struct {
	Id        string `json:"id"`
	Objective string `json:"objective"`
	Deadline  string `json:"deadline"`
	PlannerId string `json:"planner_id"`
}

type UpdateGoalResponse struct {
	ID string `json:"id"`
}

// UpdateGoal updates a goal based on the provided request.
func (c *GoalControl) UpdateGoal(req *UpdateGoalRequest) error {
	m := &models.Goal{}
	// convert deadline to time.Time
	deadline, err := m.ConvertDeadline(req.Deadline)
	goal := m.GenerateGoalInstance(req.Id, req.Objective, deadline)

	// for now, updategoal sets the createdat and updatedat fields to the current time
	// this is because the frontend does not have a way to set these fields
	err = c.Service.UpdateGoal(goal)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGoalRequest represents a request to delete a goal.
// The Id field represents the identifier of the goal to be deleted.
type DeleteGoalRequest struct {
	Id string `json:"id"`
}

// DeleteGoal deletes a goal based on the provided request.
func (c *GoalControl) DeleteGoal(req *DeleteGoalRequest) error {
	return c.Service.DeleteGoal(req.Id)
}

// GetGoalRequest represents a request to retrieve a goal.
// The Id field represents the identifier of the goal to be retrieved.
type GetGoalRequest struct {
	Id string `json:"id"`
}

// GetGoalResponse represents the response for retrieving a goal.
// It contains the goal retrieved from the service.
type GetGoalResponse struct {
	Goal *models.Goal `json:"goal"`
}

// GetGoal retrieves a goal based on the provided request.
func (c *GoalControl) GetGoal(req *GetGoalRequest) (*GetGoalResponse, error) {
	goal, err := c.Service.GetGoal(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetGoalResponse{
		Goal: goal,
	}, nil
}

type GetGoalByObjectiveRequest struct {
	Objective string `json:"objective"`
}

type GetGoalByObjectiveResponse struct {
	Goal *models.Goal `json:"goal"`
}

func (c *GoalControl) GetGoalByObjective(req *GetGoalByObjectiveRequest) (*GetGoalByObjectiveResponse, error) {
	goal, err := c.Service.GetGoalByObjective(req.Objective)
	if err != nil {
		return nil, err
	}
	return &GetGoalByObjectiveResponse{
		Goal: goal,
	}, nil
}

type GetGoalsByPlannerIdRequest struct {
	PlannerId string `json:"planner_id"`
}

type GetGoalsByPlannerIdResponse struct {
	Goals []*models.Goal `json:"goals"`
}

func (c *GoalControl) GetGoalsByPlannerId(req *GetGoalsByPlannerIdRequest) (*GetGoalsByPlannerIdResponse, error) {
	goals, err := c.Service.GetGoalsByPlannerId(req.PlannerId)
	if err != nil {
		return nil, err
	}
	return &GetGoalsByPlannerIdResponse{
		Goals: goals,
	}, nil
}

// ListGoalsResponse represents the response for listing goals.
// It contains a slice of goals retrieved from the service.
type ListGoalsResponse struct {
	Goals []*models.Goal `json:"goals"`
}

// ListGoals retrieves a list of goals based on the provided request.
func (c *GoalControl) ListGoals() (*ListGoalsResponse, error) {
	goals, err := c.Service.ListGoals()
	if err != nil {
		return nil, err
	}
	return &ListGoalsResponse{
		Goals: goals,
	}, nil
}

// generateGoalUUID generates a new UUID for a goal.
func generateGoalUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
