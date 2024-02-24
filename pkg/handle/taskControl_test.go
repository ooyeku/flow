package handle

import (
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/internal/inmemory"
	"github.com/ooyeku/flow/pkg/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func SetupTaskT(t *testing.T) (*TaskControl, *storm.DB) {
	db, _ := storm.Open("test.db")
	tStore := inmemory.NewInMemoryTaskStore(db)
	service := services.NewTaskService(tStore)
	taskControl := NewTaskControl(service)
	return taskControl, db
}

func TeardownTaskT(t *testing.T, db *storm.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("failed to close db: %v", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Fatalf("failed to remove db: %v", err)
	}
}

func TestTaskControl_CreateTask(t *testing.T) {
	taskControl, db := SetupTaskT(t)
	defer TeardownTaskT(t, db)
	req := &CreateTaskRequest{
		Title:       "My Task",
		Description: "This is a test task",
	}

	res, err := taskControl.CreateTask(*req)
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	assert.NotEmpty(t, res.ID)
}
