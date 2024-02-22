package store

import "github.com/ooyeku/flow/pkg/models"

// GoalStore is an interface that defines methods for creating, updating, deleting, retrieving, and listing goals.
// CreateGoal creates a new goal.
type GoalStore interface {
	CreateGoal(goal *models.Goal) error
	UpdateGoal(goal *models.Goal) error
	DeleteGoal(id string) error
	GetGoal(id string) (*models.Goal, error)
	ListGoals() ([]*models.Goal, error)
	GetGoalByObjective(objective string) (*models.Goal, error)
	GetGoalsByPlannerId(id string) ([]*models.Goal, error)
}
