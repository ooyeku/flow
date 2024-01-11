package mock

import (
	"github.com/stretchr/testify/assert"
	"goworkflow/pkg/models"
	"testing"
)

func TestMockPlannerStore_CreatePlanner(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)

	assert.NoError(t, err)
	assert.Equal(t, planner, store.Planner[planner.Id])
}

func TestMockPlannerStore_CreatePlanner_AlreadyExists(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)
	if err != nil {
		return
	}
	err = store.CreatePlanner(planner)

	assert.Error(t, err)
}

func TestMockPlannerStore_UpdatePlanner(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)
	if err != nil {
		return
	}
	planner.UserId = "2"
	err = store.UpdatePlanner(planner)

	assert.NoError(t, err)
	assert.Equal(t, planner, store.Planner[planner.Id])
}

func TestMockPlannerStore_UpdatePlanner_NotExists(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.UpdatePlanner(planner)
	assert.Error(t, err)
}

func TestMockPlannerStore_DeletePlanner(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)
	if err != nil {
		return
	}
	err = store.DeletePlanner(planner.Id)

	assert.NoError(t, err)
	_, exists := store.Planner[planner.Id]
	assert.False(t, exists)
}

func TestMockPlannerStore_DeletePlanner_NotExists(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.DeletePlanner(planner.Id)

	assert.Error(t, err)
}

func TestMockPlannerStore_GetPlanner(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)
	if err != nil {
		return
	}
	_, err = store.GetPlanner(planner.Id)

	assert.NoError(t, err)
}

func TestMockPlannerStore_GetPlanner_NotExists(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}

	_, err := store.GetPlanner(planner.Id)

	assert.Error(t, err)
}

func TestMockPlannerStore_ListPlanners(t *testing.T) {
	store := NewMockPlannerStore()
	planner := &models.Planner{
		Id:     "1",
		UserId: "1",
		Goals:  []models.Goal{},
	}
	planner2 := &models.Planner{
		Id:     "2",
		UserId: "2",
		Goals:  []models.Goal{},
	}

	err := store.CreatePlanner(planner)
	if err != nil {
		return
	}
	err = store.CreatePlanner(planner2)
	if err != nil {
		return
	}
	planners, err := store.ListPlanners()

	assert.NoError(t, err)
	assert.Equal(t, planner, planners[0])
	assert.Equal(t, planner2, planners[1])
}

func TestMockPlannerStore_ListPlanners_Empty(t *testing.T) {
	store := NewMockPlannerStore()
	planners, err := store.ListPlanners()

	assert.NoError(t, err)
	assert.Empty(t, planners)
}
