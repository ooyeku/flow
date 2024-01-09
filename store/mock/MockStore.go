package mock

import (
	"errors"
	"goworkflow/models"
)

type MockStore struct {
	tasks map[string]*models.Task
}

func NewMockStore() *MockStore {
	return &MockStore{
		tasks: make(map[string]*models.Task),
	}
}

func (m *MockStore) CreateTask(task *models.Task) error {
	if _, exists := m.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}
	m.tasks[task.ID] = task
	return nil
}

func (m *MockStore) UpdateTask(id string, task *models.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return errors.New("task not found")
	}
	m.tasks[id] = task
	return nil
}

func (m *MockStore) DeleteTask(id string) error {
	if _, exists := m.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(m.tasks, id)
	return nil
}

func (m *MockStore) GetTask(id string) (*models.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *MockStore) ListTasks() ([]*models.Task, error) {
	tasks := make([]*models.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
