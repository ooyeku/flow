package handle

import (
	"flow/internal/models"
	"flow/pkg/services"
)

type VersionControl struct {
	versionService *services.VersionService
}

func NewVersionControl(versionService *services.VersionService) *VersionControl {
	return &VersionControl{
		versionService: versionService,
	}
}

func (vc *VersionControl) CreateVersion(version *models.Version) error {
	return vc.versionService.CreateVersion(version)
}

func (vc *VersionControl) UpdateVersion(id string, version *models.Version) error {
	return vc.versionService.UpdateVersion(id, version)
}

func (vc *VersionControl) GetVersion(id string) (*models.Version, error) {
	return vc.versionService.GetVersion(id)
}

func (vc *VersionControl) ListVersions() ([]*models.Version, error) {
	return vc.versionService.ListVersions()
}

func (vc *VersionControl) GetPreviousVersion(id string) (*models.Version, error) {
	return vc.versionService.GetPreviousVersion(id)
}
