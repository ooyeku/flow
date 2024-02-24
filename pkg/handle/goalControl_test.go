package handle

import (
	"github.com/asdine/storm"
	_ "github.com/google/uuid"
	"github.com/ooyeku/flow/internal/inmemory"
	_ "github.com/ooyeku/flow/pkg/models"
	"github.com/ooyeku/flow/pkg/services"
	_ "github.com/ooyeku/flow/pkg/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupGoalT(t *testing.T) (*GoalControl, *storm.DB) {
	db, _ := storm.Open("test.db")
	tStore := inmemory.NewInMemoryGoalStore(db)
	service := services.NewGoalService(tStore)
	goalControl := NewGoalControl(service)
	return goalControl, db
}

func TeardownGoalT(t *testing.T, db *storm.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("failed to close db: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Fatalf("failed to remove db: %v", err)
	}

}

func TestGoalControl_CreateGoal(t *testing.T) {
	goalControl, db := SetupGoalT(t)
	defer TeardownGoalT(t, db)
	req := &CreateGoalRequest{
		Objective: "objective",
		Deadline:  "2021-01-01",
		PlannerId: "planner1",
	}

	res, err := goalControl.CreateGoal(req)
	if err != nil {
		t.Fatalf("failed to create goal: %v", err)
	}

	assert.NotEmpty(t, res.ID)
}
