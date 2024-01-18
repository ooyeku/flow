package inmemory

import (
	"flow/internal/models"
	"github.com/asdine/storm"
)

// BoltPlannerStore represents a store for managing planners using a BoltDB database.
type BoltPlannerStore struct {
	db *storm.DB
}

// GetPlannerByTitle retrieves a planner with the specified title from the BoltPlannerStore.
func (s *BoltPlannerStore) GetPlannerByTitle(title string) (*models.Planner, error) {
	planner := new(models.Planner)
	if err := s.db.One("Title", title, planner); err != nil {
		return nil, err
	}
	return planner, nil
}

// GetPlannerByOwner returns a list of planners owned by the specified user ID.
// It queries the database to find all planners that have a "UserId" field matching the provided ID.
// If an error occurs during the query, it returns nil and the error.
// Otherwise, it returns the list of found planners and nil for the error.
func (s *BoltPlannerStore) GetPlannerByOwner(id string) ([]*models.Planner, error) {
	var planners []*models.Planner
	if err := s.db.Find("UserId", id, &planners); err != nil {
		return nil, err
	}
	return planners, nil
}

// NewInMemoryPlannerStore returns a new instance of the BoltPlannerStore type with the provided storm.DB instance as its database.
func NewInMemoryPlannerStore(db *storm.DB) *BoltPlannerStore {
	return &BoltPlannerStore{
		db: db,
	}
}

// CreatePlanner creates a new planner in the BoltPlannerStore.
// It takes a pointer to a models.Planner object as its argument and returns an error.
// It saves the planner object to the underlying BoltDB database using the db.Save() method.
// If the save operation fails, it returns an error.
func (s *BoltPlannerStore) CreatePlanner(planner *models.Planner) error {
	return s.db.Save(planner)
}

// UpdatePlanner updates the details of a planner in the Bolt DB.
// It takes a *models.Planner as input and returns an error if the update operation fails.
func (s *BoltPlannerStore) UpdatePlanner(planner *models.Planner) error {
	return s.db.Update(planner)
}

// DeletePlanner deletes a planner from the BoltPlannerStore. It takes an ID as a parameter and returns an error.
func (s *BoltPlannerStore) DeletePlanner(id string) error {
	planner := new(models.Planner)
	planner.Id = id
	return s.db.DeleteStruct(planner)
}

// GetPlanner retrieves a planner by its ID from the BoltPlannerStore.
//
// Parameters:
// - id: the ID of the planner to retrieve.
//
// Returns:
// - *models.Planner: the retrieved planner.
// - error: an error if the planner cannot be retrieved.
func (s *BoltPlannerStore) GetPlanner(id string) (*models.Planner, error) {
	planner := new(models.Planner)
	if err := s.db.One("ID", id, planner); err != nil {
		return nil, err
	}
	return planner, nil
}

// ListPlanners returns a list of all planners in the BoltPlannerStore database.
// Each planner is of type *models.Planner.
// If there is an error during the retrieval process, an error is returned along with the empty list of planners.
func (s *BoltPlannerStore) ListPlanners() ([]*models.Planner, error) {
	var planners []*models.Planner
	if err := s.db.All(&planners); err != nil {
		return nil, err
	}
	return planners, nil
}
