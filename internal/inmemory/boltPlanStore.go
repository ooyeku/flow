package inmemory

import (
	"flow/internal/models"
	"github.com/asdine/storm"
	"github.com/google/uuid"
)

// BoltPlanStore represents a store for managing plans using BoltDB.
type BoltPlanStore struct {
	db *storm.DB
}

// GetPlanByName retrieves a plan by its name.
// It takes a string parameter 'name' representing the plan name.
// It returns a pointer to a Plan object and an error. If there was an issue while retrieving the plan from the database, an error is returned. Otherwise, the plan is returned.
func (s *BoltPlanStore) GetPlanByName(name string) (*models.Plan, error) {
	plan := new(models.Plan)
	if err := s.db.One("PlanName", name, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

// GetPlansByGoal retrieves plans associated with a specific goal ID.
// It takes a string parameter 'id' representing the goal ID.
// It returns a slice of pointers to Plan objects and an error. If there was an issue while retrieving the plans from the database, an error is returned. Otherwise, the slice of plans
func (s *BoltPlanStore) GetPlansByGoal(id string) ([]*models.Plan, error) {
	var plans []*models.Plan
	if err := s.db.Find("GoalId", id, &plans); err != nil {
		return nil, err
	}
	return plans, nil
}

// NewInMemoryPlanStore creates a new instance of BoltPlanStore by initializing it with the provided DB instance.
func NewInMemoryPlanStore(db *storm.DB) *BoltPlanStore {
	return &BoltPlanStore{
		db: db,
	}
}

// CreatePlan inserts a new plan into the BoltDB database.
// It takes a pointer to a models.Plan as an argument and returns an error.
// The function calls the Save method of the underlying BoltDB connection,
// passing the plan as the argument to save it as a new record in the database.
func (s *BoltPlanStore) CreatePlan(plan *models.Plan) error {
	return s.db.Save(plan)
}

// UpdatePlan updates an existing plan in the BoltPlanStore.
// It takes a pointer to a Plan object representing the plan to be updated.
// It returns an error if there was an issue while updating the plan in the database.
func (s *BoltPlanStore) UpdatePlan(plan *models.Plan) error {
	return s.db.Update(plan)
}

// DeletePlan deletes a plan from the BoltPlanStore.
// It takes in an id string as a parameter.
// It creates a new Plan with the given id.
// It then calls the DeleteStruct method on the BoltPlanStore's database connection, passing in the Plan.
// The DeleteStruct method returns an error if there was an issue deleting the Plan.
// If the deletion was successful, nil is returned.
// Example usage:
// err := store.DeletePlan("123")
//
//	if err != nil {
//	    fmt.Println("Error deleting plan:", err)
//	} else {
//
//	    fmt.Println("Plan deleted successfully")
//	}
func (s *BoltPlanStore) DeletePlan(id string) error {
	plan := new(models.Plan)
	plan.Id = id
	return s.db.DeleteStruct(plan)
}

// GetPlan retrieves a plan from the BoltPlanStore based on the specified ID.
// It takes a string parameter "id" representing the ID of the plan to be retrieved.
// Returns a pointer to a models.Plan object and an error.
// If there was an issue retrieving the plan from the database, an error is returned.
//
// Example usage:
//
//	plan, err := store.GetPlan("123")
//	if err != nil {
//	    log.Println("Failed to get plan:", err)
//	}
//	fmt.Println(plan)
//
// Note: This method is part of the BoltPlanStore struct that provides methods for managing plans.
// The BoltPlanStore struct uses the storm.DB instance for database operations.
// To access other methods in the BoltPlanStore, create an instance of the struct and call the desired method.
//
// See also: BoltPlanStore, models.Plan
func (s *BoltPlanStore) GetPlan(id string) (*models.Plan, error) {
	plan := new(models.Plan)
	if err := s.db.One("Id", id, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

// ListPlans retrieves all plans from the BoltPlanStore.
// It returns a slice of pointers to Plan objects representing the plans,
// and an error if there was an issue while retrieving the plans from the database.
func (s *BoltPlanStore) ListPlans() ([]*models.Plan, error) {
	var plans []*models.Plan
	if err := s.db.All(&plans); err != nil {
		return nil, err
	}
	return plans, nil
}

// generatePlanUUID is a function that generates a new UUID for a plan.
// It uses the uuid.New() function from the "github.com/google/uuid" package
// to generate a new UUID. The generated UUID is then converted to a string
// using the String() method. The function returns the generated UUID as a string
// and a nil error if the generation process is successful.
// Example usage:
// uuid, err := generatePlanUUID()
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(uuid)
// Output: f47ac10b-58cc-4372-a567-0e02b2c3d479
func generatePlanUUID() (string, error) {
	return uuid.New().String(), nil
}
