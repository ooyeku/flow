package store

import "goworkflow/models"

// PlannerStore is an interface that defines the methods for managing planners in a store.
// It provides functionality for creating, updating, deleting, retrieving, and listing planners.
type PlannerStore interface {
	CreatePlanner(planner *models.Planner) error
	UpdatePlanner(planner *models.Planner) error
	DeletePlanner(id string) error
	GetPlanner(id string) (*models.Planner, error)
	ListPlanners() ([]*models.Planner, error)
}
