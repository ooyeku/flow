package services

import (
	"flow/internal/models"
	"flow/pkg/store"
)

// VersionService is a struct that provides methods for managing versions of goals.
// AddVersion adds a new version of a goal to the version store.
type VersionService struct {
	versionStore store.VersionStores
}

// NewVersionService creates a new instance of VersionService with the provided versionStore.
// It returns a pointer to the VersionService.
func NewVersionService(versionStore store.VersionStores) *VersionService {
	return &VersionService{
		versionStore: versionStore,
	}
}

// AddVersion adds a new version to the VersionService's versionStore.
//
// Parameters:
//   - v: the Version instance to be added
func (service *VersionService) AddVersion(v *models.Version) {
	service.versionStore.AddVersion(v)
}

// UpdateVersion updates the version of a goal identified by the given goal ID
// with the provided version information. The updateType specifies the new version
// details such as Major, Minor, and Patch numbers.
//
// Parameters:
// - goalId: The ID of the goal to update the version for.
// - updateType: The new version information to set for the goal.
//
// Note: The UpdateVersion method internally calls the UpdateVersion method of the
// versionStore to update the version information in the data store.
//
// Example:
//
//	goalId := "12345"
//	updateType := &models.VersionInfo{
//	    Major: 1,
//	    Minor: 2,
//	    Patch: 0,
//	}
//	service.UpdateVersion(goalId, updateType)
func (service *VersionService) UpdateVersion(goalId *models.EntityID, updateType *models.VersionInfo) {
	service.versionStore.UpdateVersion(goalId, updateType)
}

// GetCurrentVersion retrieves the current version of a goal from the VersionService.
// It takes a goalId of type *models.EntityID as a parameter and returns the current version of the goal as *models.Version and an error if any.
// It internally calls the GetCurrentVersion method of the versionStore field to fetch the current version from the storage.
func (service *VersionService) GetCurrentVersion(goalId *models.EntityID) (*models.Version, error) {
	return service.versionStore.GetCurrentVersion(goalId)
}

// ListVersions retrieves a list of versions associated with a goal ID.
// It returns a slice of Version pointers and an error, if any.
func (service *VersionService) ListVersions(goalId *models.EntityID) ([]*models.Version, error) {
	return service.versionStore.ListVersions(goalId)
}

// GetImage retrieves the image of a specific version of a goal. It takes in the goal ID and version number as parameters.
// The method returns a Snapshot, which contains the goal, plans, and tasks related to the specified version, along with an error (if any).
// Example usage:
//
//	  goalId := models.EntityID("goal-1")
//	  versionNo := &models.VersionInfo{Major: 1, Minor: 0, Patch: 0}
//	  snapshot, err := versionService.GetImage(&goalId, versionNo)
//	  if err != nil {
//	     fmt.Println("Error:", err)
//		} else {
//	  	fmt.Println("Snapshot:", snapshot)
//		}
func (service *VersionService) GetImage(goalId *models.EntityID, versionNo *models.VersionInfo) (*models.Snapshot, error) {
	return service.versionStore.GetImage(goalId, versionNo)
}
