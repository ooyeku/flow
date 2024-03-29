package store

import "github.com/ooyeku/flow/pkg/models"

// PlanStore is an interface that defines the methods to interact with plans.
// CreatePlan creates a new plan.
type PlanStore interface {
	CreatePlan(plan *models.Plan) error
	UpdatePlan(plan *models.Plan) error
	DeletePlan(id string) error
	GetPlan(id string) (*models.Plan, error)
	ListPlans() ([]*models.Plan, error)
	GetPlanByName(name string) (*models.Plan, error)
	GetPlansByGoal(id string) ([]*models.Plan, error)
}
