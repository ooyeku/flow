package handle

import (
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/internal/inmemory"
	"github.com/ooyeku/flow/pkg/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupVersionT(t *testing.T) (*VersionControl, *storm.DB) {
	db, _ := storm.Open("test.db")
	tStore := inmemory.NewInMemoryVersionStore(db)
	service := services.NewVersionService(tStore)
	versionControl := NewVersionControl(service)
	return versionControl, db
}

func TeardownVersionT(t *testing.T, db *storm.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("failed to close db: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Fatalf("failed to remove db: %v", err)
	}
}

func TestVersionControl_CreateVersion(t *testing.T) {
	versionControl, db := SetupVersionT(t)
	defer TeardownVersionT(t, db)
	req := &CreateVersionRequest{
		GoalID: "goal1",
		PlanID: "plan1",
		TaskID: "task1",
	}

	res, err := versionControl.CreateVersion(req)
	if err != nil {
		t.Fatalf("failed to create version: %v", err)
	}

	assert.NotEmpty(t, res.ID)
}
