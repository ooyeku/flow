package services_test

import (
	"errors"
	"github.com/ooyeku/flow/pkg/models"
	"github.com/ooyeku/flow/pkg/services"
	_ "github.com/ooyeku/flow/pkg/store"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockVersionStore struct {
	mock.Mock
}

func (m *MockVersionStore) CreateVersion(version *models.Version) error {
	args := m.Called(version)
	return args.Error(0)
}

func (m *MockVersionStore) UpdateVersion(id string, version *models.Version) error {
	args := m.Called(id, version)
	return args.Error(0)
}

func (m *MockVersionStore) GetVersion(id string) (*models.Version, error) {
	args := m.Called(id)
	val, ok := args.Get(0).(*models.Version)
	if !ok {
		// Check if the value is nil, if it is, return nil
		if args.Get(0) == nil {
			return nil, args.Error(1)
		}
		// If it is not nil, but it is also not *models.Version, panic
		panic("Value not of type *models.Version")
	}
	return val, args.Error(1)
}

func (m *MockVersionStore) ListVersions() ([]*models.Version, error) {
	args := m.Called()
	return args.Get(0).([]*models.Version), args.Error(1)
}

func (m *MockVersionStore) GetPreviousVersion(id string) (*models.Version, error) {
	args := m.Called(id)
	val, ok := args.Get(0).(*models.Version)
	if !ok {
		// Check if the value is nil, if it is, return nil
		if args.Get(0) == nil {
			return nil, args.Error(1)
		}
		// If it is not nil, but it is also not *models.Version, panic
		panic("Value not of type *models.Version")
	}
	return val, args.Error(1)
}

func TestCreateVersion(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	version := &models.Version{GoalID: "1", PlanID: "1", TaskID: "1"}

	mockStore.On("CreateVersion", version).Return(nil)

	err := versionService.CreateVersion(version)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockStore.AssertExpectations(t)
}

func TestUpdateVersion(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	version := &models.Version{GoalID: "1", PlanID: "1", TaskID: "1"}

	mockStore.On("UpdateVersion", "1", version).Return(nil)

	err := versionService.UpdateVersion("1", version)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockStore.AssertExpectations(t)
}

func TestGetVersion(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	version := &models.Version{GoalID: "1", PlanID: "1", TaskID: "1"}

	mockStore.On("GetVersion", "1").Return(version, nil)

	result, err := versionService.GetVersion("1")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != version {
		t.Errorf("Expected version %v, got %v", version, result)
	}

	mockStore.AssertExpectations(t)
}

func TestListVersions(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	versions := []*models.Version{
		{GoalID: "1", PlanID: "1", TaskID: "1"},
		{GoalID: "2", PlanID: "2", TaskID: "2"},
	}

	mockStore.On("ListVersions").Return(versions, nil)

	result, err := versionService.ListVersions()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != len(versions) {
		t.Errorf("Expected %d versions, got %d", len(versions), len(result))
	}

	mockStore.AssertExpectations(t)
}

func TestGetPreviousVersion(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	version := &models.Version{GoalID: "1", PlanID: "1", TaskID: "1"}

	mockStore.On("GetPreviousVersion", "1").Return(version, nil)

	result, err := versionService.GetPreviousVersion("1")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != version {
		t.Errorf("Expected version %v, got %v", version, result)
	}

	mockStore.AssertExpectations(t)
}

func TestGetVersionNotFound(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	mockStore.On("GetVersion", "1").Return(nil, errors.New("version not found"))

	_, err := versionService.GetVersion("1")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	mockStore.AssertExpectations(t)
}

func TestGetPreviousVersionNotFound(t *testing.T) {
	mockStore := new(MockVersionStore)
	versionService := services.NewVersionService(mockStore)

	mockStore.On("GetPreviousVersion", "1").Return(nil, errors.New("version not found"))

	_, err := versionService.GetPreviousVersion("1")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	mockStore.AssertExpectations(t)
}
