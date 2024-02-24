package handle

import (
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/internal/inmemory"
	"github.com/ooyeku/flow/pkg/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupPlannerT(t *testing.T) (*PlannerControl, *storm.DB) {
	db, _ := storm.Open("test.db")
	tStore := inmemory.NewInMemoryPlannerStore(db)
	service := services.NewPlannerService(tStore)
	plannerControl := NewPlannerControl(service)
	return plannerControl, db
}

func TeardownPlannerT(t *testing.T, db *storm.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("failed to close db: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Fatalf("failed to remove db: %v", err)
	}
}

func TestPlannerControl_CreatePlanner(t *testing.T) {
	plannerControl, db := SetupPlannerT(t)
	defer TeardownPlannerT(t, db)
	req := &CreatePlannerRequest{
		Title:  "My Planner",
		UserId: "user1",
	}

	res, err := plannerControl.CreatePlanner(req)
	if err != nil {
		t.Fatalf("failed to create planner: %v", err)
	}

	assert.NotEmpty(t, res.Id)
}
