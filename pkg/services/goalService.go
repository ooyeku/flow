package services

import (
	"goworkflow/internal/models"
	store2 "goworkflow/pkg/store"
)

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
