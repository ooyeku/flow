package inmemory

import (
	"github.com/asdine/storm"
	"github.com/google/uuid"
	"goworkflow/pkg/models"
)

// BoltPlanStore represents a store for managing plans using BoltDB.
type BoltPlanStore struct {
	db *storm.DB
}

// NewInMemoryPlanStore is a function that creates a new instance of BoltPlanStore
// with the provided Storm database connection.
// It returns a pointer to the newly created BoltPlanStore.
func NewInMemoryPlanStore(db *storm.DB) *BoltPlanStore {
	return &BoltPlanStore{
		db: db,
	}
}

// CreatePlan creates a new plan in the BoltPlanStore.
// It takes a pointer to a Plan object representing the plan to be created.
// It returns an error if there was an issue while saving the plan to the database.
func (s *BoltPlanStore) CreatePlan(plan *models.Plan) error {
	return s.db.Save(plan)
}

// UpdatePlan updates the given plan in the BoltDB database.
// It takes a pointer to a models.Plan as an argument and returns an error.
// The function calls the Update method of the underlying BoltDB connection,
// passing the plan as the argument to update the corresponding record in the database.
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

// GetPlan retrieves a specific plan from the BoltPlanStore based on its ID.
// The plan is returned as a pointer to models.Plan object.
// If the plan is not found, an error is returned.
// Example usage:
//
//	plan, err := store.GetPlan("123")
//	if err != nil {
//	  log.Println("Error:", err)
//	} else {
//	  log.Println("Plan:", plan)
//	}
func (s *BoltPlanStore) GetPlan(id string) (*models.Plan, error) {
	plan := new(models.Plan)
	if err := s.db.One("ID", id, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

// ListPlans returns a list of all plans stored in the BoltPlanStore.
// It retrieves all plans from the underlying BoltDB database and returns them as a slice of pointers to models.Plan.
// If an error occurs during the retrieval process, it returns nil and the corresponding error.
func (s *BoltPlanStore) ListPlans() ([]*models.Plan, error) {
	var plans []*models.Plan
	if err := s.db.All(&plans); err != nil {
		return nil, err
	}
	return plans, nil
}

// generatePlanUUID generates a new UUID string for a plan.
func generatePlanUUID() (string, error) {
	return uuid.New().String(), nil
}
