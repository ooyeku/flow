package inmemory

import (
	"github.com/asdine/storm"
	"github.com/google/uuid"
	"goworkflow/models"
)

// BoltGoalStore represents a goal store implementation that uses BoltDB as the underlying database.
// CreateGoal creates a new goal in the store.
type BoltGoalStore struct {
	db *storm.DB
}

// NewInMemoryGoalStore creates a new instance of BoltGoalStore with the given storm.DB instance as its dependency.
// It returns a pointer to the BoltGoalStore.
func NewInMemoryGoalStore(db *storm.DB) *BoltGoalStore {
	return &BoltGoalStore{
		db: db,
	}
}

// CreateGoal takes a Goal object and saves it to the database.
// An error is returned if the save operation fails.
func (s *BoltGoalStore) CreateGoal(goal *models.Goal) error {
	return s.db.Save(goal)
}

// UpdateGoal updates the specified Goal in the BoltGoalStore.
// It takes a Goal object as a parameter and updates the corresponding record in the database.
// The Goal object should have the following fields:
// - Id         : string
// - Objective  : string
// - Plans      : []Plan
// - GoalStatus : string
// - GoalCreatedAt : time.Time
// - GoalUpdatedAt : time.Time
// - Deadline   : time.Time
// - PlannerId  : string
// Returns an error if the update operation fails.
func (s *BoltGoalStore) UpdateGoal(goal *models.Goal) error {
	return s.db.Update(goal)
}

// DeleteGoal deletes a goal with the specified ID from the BoltGoalStore.
// It creates a new Goal instance with the given ID, and calls the DeleteStruct
// method of the BoltDB client to delete the goal from the database.
// If successful, it returns nil. Otherwise, an error is returned.
func (s *BoltGoalStore) DeleteGoal(id string) error {
	goal := new(models.Goal)
	goal.Id = id
	return s.db.DeleteStruct(goal)
}

// GetGoal retrieves a goal from the BoltGoalStore based on its ID.
//
// Parameters:
// - id: The ID of the goal to retrieve.
//
// Returns:
// - *models.Goal: The goal object that matches the given ID.
// - error: An error if the goal could not be retrieved.
func (s *BoltGoalStore) GetGoal(id string) (*models.Goal, error) {
	goal := new(models.Goal)
	if err := s.db.One("ID", id, goal); err != nil {
		return nil, err
	}
	return goal, nil
}

// ListGoals returns a list of all goals stored in BoltGoalStore.
// It retrieves the goals from the BoltDB database and returns them as a slice of Goal pointers.
// If an error occurs during the retrieval process, it returns an empty slice and the error.
func (s *BoltGoalStore) ListGoals() ([]*models.Goal, error) {
	var goals []*models.Goal
	if err := s.db.All(&goals); err != nil {
		return nil, err
	}
	return goals, nil
}

// generateGoalUUID is a function that generates a new UUID for a goal.
// It returns a string representation of the generated UUID and an error, if any.
// The function uses the uuid package to generate the UUID.
func generateGoalUUID() (string, error) {
	return uuid.New().String(), nil
}
