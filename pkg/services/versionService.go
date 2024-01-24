package services

import (
	"flow/internal/models"
	"flow/pkg/store"
)

type VersionService struct {
	versionStore store.VersionStore
}

func NewVersionService(versionStore store.VersionStore) *VersionService {
	return &VersionService{
		versionStore: versionStore,
	}
}

func (service *VersionService) CreateVersion(version *models.Version) error {
	return service.versionStore.CreateVersion(version)
}

func (service *VersionService) UpdateVersion(id string, version *models.Version) error {
	return service.versionStore.UpdateVersion(id, version)
}

func (service *VersionService) GetVersion(id string) (*models.Version, error) {
	return service.versionStore.GetVersion(id)
}

func (service *VersionService) ListVersions() ([]*models.Version, error) {
	return service.versionStore.ListVersions()
}

func (service *VersionService) GetPreviousVersion(id string) (*models.Version, error) {
	return service.versionStore.GetPreviousVersion(id)
}
