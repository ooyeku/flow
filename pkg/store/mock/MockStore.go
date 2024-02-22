package mock

import (
	"errors"
	"github.com/ooyeku/flow/pkg/models"
	"sync"
)

type Store struct {
	Tasks map[string]*models.Task
	mu    sync.Mutex // protects the Tasks map to prevent data races
}

func NewMockStore() *Store {
	return &Store{
		Tasks: make(map[string]*models.Task),
	}
}

func (m *Store) CreateTask(task *models.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Tasks[task.ID]; exists {
		return errors.New("task referenced by task ID already exists")
	}
	m.Tasks[task.ID] = task
	return nil
}

func (m *Store) UpdateTask(id string, task *models.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Tasks[id]; !exists {
		return errors.New("task referenced by task ID does not exist")
	}
	m.Tasks[id] = task
	return nil
}

func (m *Store) DeleteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(m.Tasks, id)
	return nil
}

func (m *Store) GetTask(id string) (*models.Task, error) {
	task, exists := m.Tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *Store) ListTasks() ([]*models.Task, error) {
	numTasks := len(m.Tasks)
	tasks := make([]*models.Task, 0, numTasks)

	if numTasks == 0 {
		return tasks, nil
	}
	for _, task := range m.Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
