package services

import (
	"github.com/ooyeku/flow/pkg/models"
	store2 "github.com/ooyeku/flow/pkg/store"
)

// GoalService is a type that provides operations for managing goals.
type GoalService struct {
	store store2.GoalStore
}

// NewGoalService is a function that creates a new instance of GoalService.
// It takes a GoalStore as a parameter and returns a pointer to a GoalService.
// The GoalStore interface is used to interact with the data store.
// Example usage:
// goalStore := inmemory.NewInMemoryGoalStore(db)
// goalService := services.NewGoalService(goalStore)
// newGoalService := NewGoalService(goalStore)
func NewGoalService(store store2.GoalStore) *GoalService {
	return &GoalService{
		store: store,
	}
}

// CreateGoal creates a new goal using the provided goal object
func (s *GoalService) CreateGoal(goal *models.Goal) error {
	return s.store.CreateGoal(goal)
}

// UpdateGoal is a method of the GoalService struct that updates an existing goal with the provided data.
// It takes a pointer to a Goal struct as a parameter and returns an error.
// The method uses the UpdateGoal method of the GoalStore interface to update the goal in the data store.
func (s *GoalService) UpdateGoal(goal *models.Goal) error {
	return s.store.UpdateGoal(goal)
}

// DeleteGoal deletes a goal with the specified ID.
func (s *GoalService) DeleteGoal(id string) error {
	return s.store.DeleteGoal(id)
}

// GetGoal is a method of the GoalService struct that retrieves a goal with the specified ID.
// It takes a string parameter `id` representing the ID of the goal to be retrieved.
// The method returns a pointer to a Goal struct and an error. If the goal is not found, it returns nil and an error.
// The method uses the GetGoal method of the GoalStore interface to retrieve the goal from the data store.
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

// ListGoals is a method of the GoalService struct that retrieves a list of goals.
// It returns a slice of pointers to Goal structs and an error.
// The method uses the ListGoals method of the GoalStore interface to fetch the goals from the data store.
func (s *GoalService) ListGoals() ([]*models.Goal, error) {
	return s.store.ListGoals()
}

// GetGoalByObjective retrieves a goal by its objective.
// It takes a string parameter `objective` which specifies the objective of the goal.
// It returns a pointer to a models.Goal and an error.
// The method calls the GetGoalByObjective method of the GoalStore interface to fetch the goal with the given objective.
// Example usage:
//
//	goal, err := c.Service.GetGoalByObjective(req.Objective)
//	if err != nil {
//	  return nil, err
//	}
//	return &GetGoalByObjectiveResponse{
//	  Goal: goal,
//	}, nil
func (s *GoalService) GetGoalByObjective(objective string) (*models.Goal, error) {
	return s.store.GetGoalByObjective(objective)
}

// GetGoalsByPlannerId is a method of the GoalService struct that retrieves all goals associated with a given planner ID.
// It takes a plannerId string as a parameter and returns a slice of pointers to Goal structs and an error.
// The method uses the GetGoalsByPlannerId method of the GoalStore interface to fetch the goals from the data store.
// Example usage:
//
//	  req := &GetGoalsByPlannerIdRequest{
//		   PlannerID: "planner-123",
//	  }
//	  resp, err := control.GetGoalsByPlannerId(req)
//	  if err != nil {
//		   log.Fatal(err)
//	  }
//
// Note: The GetGoalsByPlannerIdRequest and GetGoalsByPlannerIdResponse types are not shown here, but they are used in the example above to pass the request and retrieve the response
func (s *GoalService) GetGoalsByPlannerId(plannerId string) ([]*models.Goal, error) {
	return s.store.GetGoalsByPlannerId(plannerId)
}
