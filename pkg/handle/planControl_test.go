package handle

import (
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/internal/inmemory"
	"github.com/ooyeku/flow/pkg/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupPlanT(t *testing.T) (*PlanControl, *storm.DB) {
	db, _ := storm.Open("test.db")
	tStore := inmemory.NewInMemoryPlanStore(db)
	service := services.NewPlanService(tStore)
	planControl := NewPlanControl(service)
	return planControl, db
}

func TeardownPlanT(t *testing.T, db *storm.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("failed to close db: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Fatalf("failed to remove db: %v", err)
	}
}

func TestPlanControl_CreatePlan(t *testing.T) {
	planControl, db := SetupPlanT(t)
	defer TeardownPlanT(t, db)
	req := &CreatePlanRequest{
		PlanName:        "My Plan",
		PlanDescription: "This is a test plan",
		PlanDate:        "2022-01-01",
		PlanTime:        "12:00",
		GoalId:          "goal1",
	}

	res, err := planControl.CreatePlan(req)
	if err != nil {
		t.Fatalf("failed to create plan: %v", err)
	}

	assert.NotEmpty(t, res.ID)
}
