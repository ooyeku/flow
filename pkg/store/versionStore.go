package store

import "flow/internal/models"

type VersionStores interface {
	AddVersion(v *models.Version)
	UpdateVersion(goalId *models.EntityID, updateType *models.VersionInfo)
	GetCurrentVersion(goalId *models.EntityID) (*models.Version, error)
	ListVersions(goalId *models.EntityID) ([]*models.Version, error)
	GetImage(goalId *models.EntityID, versionNo *models.VersionInfo) (*models.Snapshot, error)
}
