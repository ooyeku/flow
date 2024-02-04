package inmemory

import (
	"flow/internal/models"
	"github.com/asdine/storm"
	"github.com/google/uuid"
)

// BoltGoalStore represents a goal store implementation that uses BoltDB as the underlying database.
// CreateGoal creates a new goal in the store.
type BoltGoalStore struct {
	db *storm.DB
}

// NewInMemoryGoalStore is a function that returns a new instance of BoltGoalStore
// initialized with the given storm.DB.
// It takes a pointer to a storm.DB as a parameter.
// It returns a pointer to BoltGoalStore.
func NewInMemoryGoalStore(db *storm.DB) *BoltGoalStore {
	return &BoltGoalStore{
		db: db,
	}
}

// CreateGoal creates a new goal in the BoltGoalStore.
// It takes a Goal object as a parameter and saves it to the database.
// The Goal object should have the following fields:
// - Id             : string
// - Objective      : string
// - Plans          : []Plan
// - GoalStatus     : string
// - GoalCreatedAt  : time.Time
// - GoalUpdatedAt  : time.Time
// - Deadline       : time.Time
// - PlannerID      : string
// Returns an error if the save operation fails.
func (s *BoltGoalStore) CreateGoal(goal *models.Goal) error {
	return s.db.Save(goal)
}

// UpdateGoal takes a Goal object and updates it in the database.
// An error is returned if the update operation fails.
func (s *BoltGoalStore) UpdateGoal(goal *models.Goal) error {
	return s.db.Update(goal)
}

// DeleteGoal takes an ID string and deletes the goal with that ID from the database.
// It creates a new instance of models.Goal with the given ID, sets it as the ID of the goal to be deleted,
// and then calls the DeleteStruct method of the s.db (storm.DB) object to delete the goal from the database.
// An error is returned if the delete operation fails.
func (s *BoltGoalStore) DeleteGoal(id string) error {
	goal := new(models.Goal)
	goal.Id = id
	return s.db.DeleteStruct(goal)
}

// GetGoal takes an id string and returns the goal with that id from the database.
// If the goal is not found, it returns nil and an error.
// If an error occurs during the database query, it returns nil and the error.
func (s *BoltGoalStore) GetGoal(id string) (*models.Goal, error) {
	goal := new(models.Goal)
	if err := s.db.One("Id", id, goal); err != nil {
		return nil, err
	}
	return goal, nil
}

// GetGoalByObjective takes an objective string and retrieves the corresponding goal from the database.
// If the goal is found, it is returned along with nil error.
// If the goal is not found or an error occurs during the retrieval, nil goal and the error are returned.
func (s *BoltGoalStore) GetGoalByObjective(objective string) (*models.Goal, error) {
	goal := new(models.Goal)
	if err := s.db.One("Objective", objective, goal); err != nil {
		return nil, err
	}
	return goal, nil
}

// GetGoalsByPlannerId takes a plannerId string as input and retrieves
// all goals associated with the specified plannerId from the database.
// It returns a slice of Goal objects and an error. If the database
// retrieval fails, it returns nil and the error. If the retrieval
// is successful, it returns the slice of Goal objects and nil error.
func (s *BoltGoalStore) GetGoalsByPlannerId(plannerId string) ([]*models.Goal, error) {
	var goals []*models.Goal
	if err := s.db.Find("PlannerId", plannerId, &goals); err != nil {
		return nil, err
	}
	return goals, nil
}

// ListGoals retrieves all goals from the database and returns them as a slice of Goal objects.
func (s *BoltGoalStore) ListGoals() ([]*models.Goal, error) {
	var goals []*models.Goal
	if err := s.db.All(&goals); err != nil {
		return nil, err
	}
	return goals, nil
}

// generateGoalUUID generates a new unique UUID for a goal.
// It uses the uuid.New().String() function from the "github.com/google/uuid" package to generate the UUID.
// It returns the generated UUID as a string and returns nil error.
func generateGoalUUID() (string, error) {
	return uuid.New().String(), nil
}
