package store

import "flow/internal/models"

type RepositoryStore interface {
	CreateRepository(initialVersions []models.Version) models.Repository
	CreateInitialVersion(plannerID *models.ID) // create 0.0.0
	GetVersionState(goalID models.ID) models.Version
	AddVersion(v *models.Version)                                      // add version to repo
	UpdateVersionState(vInfo models.VersionInfo, snap models.Snapshot) // update version no, create and add snapshot to image
}
