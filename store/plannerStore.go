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

// GoalStore is an interface that defines methods for creating, updating, deleting, retrieving, and listing goals.
// CreateGoal creates a new goal.
type GoalStore interface {
	CreateGoal(goal *models.Goal) error
	UpdateGoal(goal *models.Goal) error
	DeleteGoal(id string) error
	GetGoal(id string) (*models.Goal, error)
	ListGoals() ([]*models.Goal, error)
}

// PlanStore is an interface that defines the methods to interact with plans.
// CreatePlan creates a new plan.
type PlanStore interface {
	CreatePlan(plan *models.Plan) error
	UpdatePlan(plan *models.Plan) error
	DeletePlan(id string) error
	GetPlan(id string) (*models.Plan, error)
	ListPlans() ([]*models.Plan, error)
}
