package store

import "flow/internal/models"

// PlannerStore is an interface that defines the methods for managing planners in a store.
// It provides functionality for creating, updating, deleting, retrieving, and listing planners.
type PlannerStore interface {
	CreatePlanner(planner *models.Planner) error
	UpdatePlanner(planner *models.Planner) error
	DeletePlanner(id string) error
	GetPlanner(id string) (*models.Planner, error)
	ListPlanners() ([]*models.Planner, error)
	GetPlannerByTitle(title string) (*models.Planner, error)
	GetPlannerByOwner(id string) ([]*models.Planner, error)
}
