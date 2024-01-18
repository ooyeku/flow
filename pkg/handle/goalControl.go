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

// NewGoalControl creates a new instance of GoalControl with the provided GoalService.
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

// CreateGoalResponse contains the ID of the newly created goal.
type CreateGoalResponse struct {
	ID string `json:"id"`
}

// CreateGoal creates a new goal based on the provided request.
// If the generation of a new UUID for the goal fails, an error will be returned.
// The deadline provided in the request will be converted to time.Time format.
// The goal will be created using the GenerateGoalInstance method of the Goal model,
// and then the CreateGoal method of the GoalService will be called to create the goal.
// If any error occurs during the goal creation process, it will be returned.
// Finally, a CreateGoalResponse containing the ID of the newly created goal will be returned along with nil error.
// Example:
//
//	goalReq := &CreateGoalRequest{
//	    Objective: "This is a new goal",
//	    Deadline:  "2022-12-31",
//	    PlannerId: "123456",
//	}
//
// goalRes, err := goalControl.CreateGoal(goalReq)
//
//	if err != nil {
//	    log.Fatalf("Failed to create goal: %v", err)
//	}
//
// fmt.Println("Goal created successfully")
// fmt.Printf("Goal ID: %s\n", goalRes.ID)
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

// UpdateGoalResponse represents a response structure for updating a goal.
// It contains the ID of the updated goal in the json:"id" field.
// Example usage:
//
//	var response UpdateGoalResponse
//	err := json.Unmarshal(data, &response)
//	if err != nil {
//	  fmt.Println("Error unmarshalling response:", err)
//	}
//	fmt.Println("Updated goal ID:", response.ID)
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
// It contains the ID of the goal to be deleted.
type DeleteGoalRequest struct {
	Id string `json:"id"`
}

// DeleteGoal deletes a goal with the specified ID.
// It calls the DeleteGoal method of GoalService to delete the goal from the store.
// The ID of the goal to be deleted is specified in the req parameter of type DeleteGoalRequest.
// If the goal is successfully deleted, nil is returned. Otherwise, an error is returned.
func (c *GoalControl) DeleteGoal(req *DeleteGoalRequest) error {
	return c.Service.DeleteGoal(req.Id)
}

// GetGoalRequest represents a request to get a goal by its ID.
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

// GetGoalByObjectiveRequest represents a request to get a goal by its objective.
type GetGoalByObjectiveRequest struct {
	Objective string `json:"objective"`
}

// GetGoalByObjectiveResponse contains the retrieved goal
type GetGoalByObjectiveResponse struct {
	Goal *models.Goal `json:"goal"`
}

// handle error
func (c *GoalControl) GetGoalByObjective(req *GetGoalByObjectiveRequest) (*GetGoalByObjectiveResponse, error) {
	goal, err := c.Service.GetGoalByObjective(req.Objective)
	if err != nil {
		return nil, err
	}
	return &GetGoalByObjectiveResponse{
		Goal: goal,
	}, nil
}

// GetGoalsByPlannerIdRequest represents a request to get goals by planner ID.
// The PlannerId field specifies the ID of the planner.
type GetGoalsByPlannerIdRequest struct {
	PlannerId string `json:"planner_id"`
}

type GetGoalsByPlannerIdResponse struct {
	Goals []*models.Goal `json:"goals"`
}

// GetGoalsByPlannerId retrieves goals by their planner ID.
//
// Parameters:
//
//	req: The request object containing the planner ID.
//
// Returns:
//
//	A GetGoalsByPlannerIdResponse object containing the retrieved goals, or an error if retrieval fails.
//
// Example usage:
//
//	plannerId := "12345"
//	req := &GetGoalsByPlannerIdRequest{
//	    PlannerId: plannerId,
//	}
//	goals, err := control.GetGoalsByPlannerId(req)
//	if err != nil {
//	    // Handle error
//	}
//	for _, goal := range goals.Goals {
//	    // Do something with the goal
//	}
func (c *GoalControl) GetGoalsByPlannerId(req *GetGoalsByPlannerIdRequest) (*GetGoalsByPlannerIdResponse, error) {
	goals, err := c.Service.GetGoalsByPlannerId(req.PlannerId)
	if err != nil {
		return nil, err
	}
	return &GetGoalsByPlannerIdResponse{
		Goals: goals,
	}, nil
}

// ListGoalsResponse represents the response structure for listing goals.
// It contains a sliced pointer to models.Goal.
// It is used in the ListGoals method of the GoalControl controller.
// Example usage:
//
//	goals, err := c.Service.ListGoals()
//	if err != nil {
//	  return nil, err
//	}
//	return &ListGoalsResponse{
//	  Goals: goals,
//	}, nil
type ListGoalsResponse struct {
	Goals []*models.Goal `json:"goals"`
}

// ListGoals retrieves a list of goals from the GoalService.
// It calls the ListGoals method of the GoalService and returns a ListGoalsResponse.
// If there is an error, it returns nil and the error.
//
// Example usage:
//
//	goals, err := goalControl.ListGoals()
//	if err != nil {
//	  fmt.Println("Error listing goals: ", err)
//	  return
//	}
//	for _, goal := range goals.Goals {
//	  fmt.Printf("Goal id: %s, Objective: %s, Deadline: %s\n", goal.Id, goal.Objective, goal.Deadline)
//	}
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
// It uses the uuid.NewRandom() function to generate a random UUID.
// It returns the generated UUID as a string and any error that occurred during the generation process.
func generateGoalUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
