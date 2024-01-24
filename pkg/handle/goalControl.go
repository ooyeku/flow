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
//
// Fields:
// - Objective: the objective of the goal.
// - Deadline: the deadline of the goal in the format "YYYY-MM-DD".
// - PlannerId: the ID of the planner associated with the goal.
type CreateGoalRequest struct {
	Objective string `json:"objective"`
	Deadline  string `json:"deadline"`
	PlannerId string `json:"planner_id"`
}

// CreateGoalResponse represents the response returned by the CreateGoal method in the GoalControl struct. It contains the ID of the created goal.
type CreateGoalResponse struct {
	ID string `json:"id"`
}

// CreateGoal creates a new goal based on the provided request.
// It generates a unique ID for the goal, converts the deadline to time.Time,
// and invokes the CreateGoal method of the underlying GoalService.
// If an error occurs during any step, it returns nil and the error.
// Otherwise, it returns a CreateGoalResponse containing the ID of the created goal.
func (c *GoalControl) CreateGoal(req *CreateGoalRequest) (*CreateGoalResponse, error) {
	id, err := generateGoalUUID()
	if err != nil {
		return nil, err
	}
	m := &models.Goal{}
	// convert deadline to time.Time
	deadline, err := m.ConvertDeadline(req.Deadline)
	goal := m.GenerateGoalInstance(id, req.Objective, deadline)
	goal.PlannerId = req.PlannerId
	err = c.Service.CreateGoal(goal)
	if err != nil {
		return nil, err
	}
	return &CreateGoalResponse{
		ID: goal.Id,
	}, nil
}

// UpdateGoalRequest represents a request to update a goal.
// It contains the ID, objective, deadline, and planner ID of the goal to be updated.
type UpdateGoalRequest struct {
	Id        string `json:"id"`
	Objective string `json:"objective"`
	Deadline  string `json:"deadline"`
	PlannerId string `json:"planner_id"`
}

// UpdateGoalResponse represents the response object of the UpdateGoal API.
// It contains the ID of the updated goal in the json field "id".
type UpdateGoalResponse struct {
	ID string `json:"id"`
}

// UpdateGoal updates a goal based on the provided request.
// The deadline provided in the request will be converted to time.Time format.
// The goal will be created using the GenerateGoalInstance method of the Goal model.
// The UpdateGoal method of the GoalService will be called to update the goal.
// If any error occurs during the goal update process, it will be returned.
// Example usage:
//
//	req := &UpdateGoalRequest{
//		Id:        "123456",
//		Objective: "New objective",
//		Deadline:  "2022-12-31",
//		PlannerID: "654321",
//	}
//	err := goalControl.UpdateGoal(req)
//	if err != nil {
//		log.Fatalf("Failed to update goal: %v", err)
//	}
//	fmt.Println("Goal updated successfully")
func (c *GoalControl) UpdateGoal(req *UpdateGoalRequest) error {
	m := &models.Goal{}
	// convert deadline to time.Time
	deadline, err := m.ConvertDeadline(req.Deadline)
	goal := m.GenerateGoalInstance(req.Id, req.Objective, deadline)
	goal.PlannerId = req.PlannerId

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
// The `Id` field specifies the ID of the goal to retrieve.
type GetGoalRequest struct {
	Id string `json:"id"`
}

// GetGoalResponse represents a response object containing a goal.
// It is used in the GetGoal method of the GoalControl struct.
// Example usage:
// resp, err := c.GetGoal(req)
type GetGoalResponse struct {
	Goal *models.Goal `json:"goal"`
}

// GetGoal retrieves a goal based on the provided request's ID.
// The method calls the GetGoal method of the GoalService to fetch the goal from the underlying data store.
// If an error occurs during the retrieval process, it will be returned.
// Otherwise, a GetGoalResponse containing the retrieved goal will be returned along with nil error.
//
// Example:
//
//	req := &GetGoalRequest{
//		Id: "123456",
//	}
//	goalRes, err := goalControl.GetGoal(req)
//	if err != nil {
//		log.Fatalf("Failed to get goal: %v", err)
//	}
//	fmt.Println("Goal retrieved successfully")
//	fmt.Printf("Goal: %+v\n", goalRes.Goal)
func (c *GoalControl) GetGoal(req *GetGoalRequest) (*GetGoalResponse, error) {
	goal, err := c.Service.GetGoal(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetGoalResponse{
		Goal: goal,
	}, nil
}

// GetGoalByObjectiveRequest represents a request for retrieving a goal by its objective.
type GetGoalByObjectiveRequest struct {
	Objective string `json:"objective"`
}

// GetGoalByObjectiveResponse is a type that represents the response of the GetGoalByObjective method in the GoalControl struct.
// It contains the Goal field, which is a pointer to the models.Goal type.
// This type is used to return the goal that matches the given objective.
type GetGoalByObjectiveResponse struct {
	Goal *models.Goal `json:"goal"`
}

// GetGoalByObjective retrieves a goal based on the provided objective.
// The objective is used as an input to the GetGoalByObjective method of the GoalService.
// If the goal is found, it will be returned in a GetGoalByObjectiveResponse.
// If any error occurs during the retrieval process, it will be returned.
//
// Example:
//
//	req := &GetGoalByObjectiveRequest{
//		Objective: "This is a goal objective",
//	}
//
//	res, err := goalControl.GetGoalByObjective(req)
//	if err != nil {
//		log.Fatalf("Failed to get goal: %v", err)
//	}
//
//	fmt.Println("Got goal successfully")
//	fmt.Printf("Goal: %+v\n", res.Goal)
func (c *GoalControl) GetGoalByObjective(req *GetGoalByObjectiveRequest) (*GetGoalByObjectiveResponse, error) {
	goal, err := c.Service.GetGoalByObjective(req.Objective)
	if err != nil {
		return nil, err
	}
	return &GetGoalByObjectiveResponse{
		Goal: goal,
	}, nil
}

// GetGoalsByPlannerIdRequest represents a request to retrieve goals by planner ID.
// The PlannerId field specifies the ID of the planner.
type GetGoalsByPlannerIdRequest struct {
	PlannerId string `json:"planner_id"`
}

// GetGoalsByPlannerIdResponse represents a response containing a list of goals associated with a planner ID.
type GetGoalsByPlannerIdResponse struct {
	Goals []*models.Goal `json:"goals"`
}

// GetGoalsByPlannerId retrieves all goals associated with a specific planner ID.
// It calls the GetGoalsByPlannerId method of the GoalService to fetch the goals.
// If there is an error during the retrieval process, it will be returned.
// Finally, a GetGoalsByPlannerIdResponse containing the retrieved goals will be returned along with a nil error.
//
// Example:
//
//	req := &GetGoalsByPlannerIdRequest{
//		PlannerID: "123456",
//	}
//	res, err := goalControl.GetGoalsByPlannerId(req)
//	if err != nil {
//		log.Fatalf("Failed to get goals: %v", err)
//	}
//	for _, goal := range res.Goals {
//		fmt.Println("Goal ID:", goal.Id)
//		fmt.Println("Objective:", goal.Objective)
//		fmt.Println("Deadline:", goal.Deadline)
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

// ListGoalsResponse represents a response type that contains a list of goals.
// The `Goals` field is a slice of pointers to `models.Goal` objects.
// It is tagged with `"json:"goals""` to specify the JSON key for this field when serializing or deserializing to JSON.
type ListGoalsResponse struct {
	Goals []*models.Goal `json:"goals"`
}

// get goal id and objective of each goal
func (c *GoalControl) ListGoals() (*ListGoalsResponse, error) {
	goals, err := c.Service.ListGoals()
	if err != nil {
		return nil, err
	}
	return &ListGoalsResponse{
		Goals: goals,
	}, nil
}

// generateGoalUUID generates a UUID string and returns it along with any error that occured during the process.
// It uses the "github.com/google/uuid" package to generate a random UUID.
func generateGoalUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
