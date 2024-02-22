package handle

import (
	"github.com/google/uuid"
	"github.com/ooyeku/flow/pkg/models"
	"github.com/ooyeku/flow/pkg/services"
	"time"
)

type VersionControl struct {
	versionService *services.VersionService
}

func NewVersionControl(versionService *services.VersionService) *VersionControl {
	return &VersionControl{
		versionService: versionService,
	}
}

type CreateVersionRequest struct {
	GoalID string `json:"goalId"`
	PlanID string `json:"planId"`
	TaskID string `json:"taskId"`
	No     struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
	CreatedBy string `json:"createdBy"`
}

type CreateVersionResponse struct {
	ID string `json:"id"`
}

func (vc *VersionControl) CreateVersion(vr *CreateVersionRequest) (*CreateVersionResponse, error) {
	// generate version id
	id, _ := generateVersionID()
	// create version
	m := &models.Version{
		ID:     models.EntityID(id),
		GoalID: models.EntityID(vr.GoalID),
		PlanID: models.EntityID(vr.PlanID),
		TaskID: models.EntityID(vr.TaskID),
		No: models.VersionInfo{
			Major: vr.No.Major,
			Minor: vr.No.Minor,
			Patch: vr.No.Patch,
		},
		CreatedAt: time.Now(),
		CreatedBy: vr.CreatedBy,
	}

	err := vc.versionService.CreateVersion(m)
	if err != nil {
		return nil, err
	}
	return &CreateVersionResponse{
		ID: string(m.ID),
	}, nil
}

type UpdateVersionRequest struct {
	ID    string `json:"id"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
	Patch int    `json:"patch"`
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

func generateVersionID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
