package store

import "flow/internal/models"

type VersionStore interface {
	CreateVersion(version *models.Version) error
	UpdateVersion(id string, version *models.Version) error
	GetVersion(id string) (*models.Version, error)
	ListVersions() ([]*models.Version, error)
	GetPreviousVersion(id string) (*models.Version, error)
}
