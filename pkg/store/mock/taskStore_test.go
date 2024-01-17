package mock

import (
	"flow/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockStore_CreateTask(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.CreateTask(task)

	assert.NoError(t, err)
	assert.Equal(t, task, store.Tasks[task.ID])
}

func TestMockStore_CreateTask_AlreadyExists(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.CreateTask(task)
	if err != nil {
		return
	}
	err = store.CreateTask(task)

	assert.Error(t, err)
}

func TestMockStore_UpdateTask(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.CreateTask(task)
	if err != nil {
		return
	}
	task.Title = "Updated Task"
	err = store.UpdateTask(task.ID, task)

	assert.NoError(t, err)
	assert.Equal(t, task, store.Tasks[task.ID])
}

func TestMockStore_UpdateTask_NotExists(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.UpdateTask(task.ID, task)

	assert.Error(t, err)
}

func TestMockStore_DeleteTask(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.CreateTask(task)
	if err != nil {
		return
	}
	err = store.DeleteTask(task.ID)

	assert.NoError(t, err)
	_, exists := store.Tasks[task.ID]
	assert.False(t, exists)
}

func TestMockStore_DeleteTask_NotExists(t *testing.T) {
	store := NewMockStore()
	err := store.DeleteTask("1")

	assert.Error(t, err)
}

func TestMockStore_GetTask(t *testing.T) {
	store := NewMockStore()
	task := &models.Task{ID: "1", Title: "Task 1"}

	err := store.CreateTask(task)
	if err != nil {
		return
	}
	retrievedTask, err := store.GetTask(task.ID)

	assert.NoError(t, err)
	assert.Equal(t, task, retrievedTask)
}

func TestMockStore_GetTask_NotExists(t *testing.T) {
	store := NewMockStore()
	_, err := store.GetTask("1")

	assert.Error(t, err)
}

func TestMockStore_ListTasks(t *testing.T) {
	store := NewMockStore()
	task1 := &models.Task{ID: "1", Title: "Task 1"}
	task2 := &models.Task{ID: "2", Title: "Task 2"}

	err := store.CreateTask(task1)
	if err != nil {
		return
	}
	err = store.CreateTask(task2)
	if err != nil {
		return
	}
	tasks, err := store.ListTasks()

	assert.NoError(t, err)
	assert.Contains(t, tasks, task1)
	assert.Contains(t, tasks, task2)
}
