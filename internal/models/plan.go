package inmemory

import (
	"github.com/asdine/storm"
	"github.com/google/uuid"
	"testing"
)

func TestBoltPlanStore(t *testing.T) {
	db, _ := storm.Open("test.db")
	defer db.Close()

	store := NewInMemoryPlanStore(db)

	// Test CreatePlan
	plan := &models.Plan{
		Id: uuid.New().String(),
	}
	err := store.CreatePlan(plan)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test GetPlan
	retrievedPlan, err := store.GetPlan(plan.Id)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if retrievedPlan.Id != plan.Id {
		t.Errorf("Expected Id %s, but got %s", plan.Id, retrievedPlan.Id)
	}

	// Test UpdatePlan
	plan.Title = "Updated Title"
	err = store.UpdatePlan(plan)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	updatedPlan, _ := store.GetPlan(plan.Id)
	if updatedPlan.Title != "Updated Title" {
		t.Errorf("Expected Title %s, but got %s", "Updated Title", updatedPlan.Title)
	}

	// Test ListPlans
	plans, err := store.ListPlans()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(plans) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(plans))
	}

	// Test DeletePlan
	err = store.DeletePlan(plan.Id)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	deletedPlan, err := store.GetPlan(plan.Id)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if deletedPlan != nil {
		t.Errorf("Expected nil, but got %v", deletedPlan)
	}

	// Test GetPlanByName
	plan2 := &models.Plan{
		Id:    uuid.New().String(),
		Title: "Test Plan",
	}
	err = store.CreatePlan(plan2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	retrievedPlan2, err := store.GetPlanByName(plan2.Title)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if retrievedPlan2.Title != plan2.Title {
		t.Errorf("Expected Title %s, but got %s", plan2.Title, retrievedPlan2.Title)
	}

	// Test GetPlansByGoal
	plan3 := &models.Plan{
		Id:     uuid.New().String(),
		GoalId: "Test Goal",
	}
	err = store.CreatePlan(plan3)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	retrievedPlans, err := store.GetPlansByGoal(plan3.GoalId)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(retrievedPlans) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(retrievedPlans))
	}
}
