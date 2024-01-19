package inmemory

import (
	"flow/internal/models"
	"github.com/asdine/storm"
	"testing"
)

func TestBoltPlanStore(t *testing.T) {

	db, _ := storm.Open("test.db")
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	store := NewInMemoryPlanStore(db)

	t.Run("GetPlanByName", func(t *testing.T) {
		// test case here
	})

	t.Run("GetPlansByGoal", func(t *testing.T) {
		// test cases here
	})

	t.Run("CreatePlan", func(t *testing.T) {
		plan := &models.Plan{Id: "plan1", PlanName: "desc1", PlanDescription: "goal1"}
		err := store.CreatePlan(plan)
		if err != nil {
			t.Fatalf("Error creating plan: %v", err)
		}

		// Simulate error case
		db.Close()
		err = store.CreatePlan(plan)
		if err == nil {
			t.Fatal("Expecting an error, but it did not occur")
		}

	})

	t.Run("UpdatePlan", func(t *testing.T) {
		// test cases here
	})

	t.Run("DeletePlan", func(t *testing.T) {
		// test cases here
	})

	t.Run("GetPlan", func(t *testing.T) {
		// test cases here
	})

	t.Run("ListPlans", func(t *testing.T) {
		db.Drop("Plan")

		err := store.CreatePlan(&models.Plan{Id: "plan1", PlanName: "desc1", PlanDescription: "goal1"})
		if err != nil {
			return
		}
		plans, err := store.ListPlans()
		if err != nil || len(plans) != 1 {
			t.Fatalf("Error listing plans: %v", err)
		}

		// Simulate error case
		err = db.Close()
		if err != nil {
			return
		}
		_, err = store.ListPlans()
		if err == nil {
			t.Fatal("Expecting an error, but it did not occur")
		}
	})

}
