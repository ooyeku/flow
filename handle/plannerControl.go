package handle

import (
	"github.com/google/uuid"
	"goworkflow/models"
	"goworkflow/services"
)

// PlannerControl represents a controller that provides methods to manage planners.
type PlannerControl struct {
	Service *services.PlannerService
}

// CreatePlannerRequest represents a request to create a planner.
// UserId is the ID of the user for whom the planner is being created.
type CreatePlannerRequest struct {
	UserId string `json:"user_id"`
}

// CreatePlannerResponse represents the response for creating a planner.
// It contains the ID of the newly created planner.
type CreatePlannerResponse struct {
	ID string `json:"id"`
}

// CreatePlanner creates a new planner for a user based on the provided request.
func (c *PlannerControl) CreatePlanner(req *CreatePlannerRequest) (*CreatePlannerResponse, error) {
	id, err := generatePlannerUUID()
	if err != nil {
		return nil, err
	}
	planner := &models.Planner{
		Id:     id,
		UserId: req.UserId,
	}
	err = c.Service.CreatePlanner(planner)
	if err != nil {
		return nil, err
	}
	return &CreatePlannerResponse{
		ID: planner.Id,
	}, nil
}

// UpdatePlannerRequest represents the data required to update a planner.
// The Id field represents the identifier of the planner to be updated.
// The UserId field represents the new user ID to be assigned to the planner after the update.
type UpdatePlannerRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

// UpdatePlanner updates the user ID of a planner specified by the ID in the provided request.
// It retrieves the planner using the ID from the service, updates the user ID, and then calls the service to update the planner.
// If any error occurs during the process, it returns the error.
// Otherwise, it returns nil.
func (c *PlannerControl) UpdatePlanner(req *UpdatePlannerRequest) error {
	planner, err := c.Service.GetPlanner(req.Id)
	if err != nil {
		return err
	}
	planner.UserId = req.UserId
	if err := c.Service.UpdatePlanner(planner); err != nil {
		return err
	}
	return nil
}

// DeletePlannerRequest represents the request for deleting a planner.
type DeletePlannerRequest struct {
	Id string `json:"id"`
}

// DeletePlanner deletes a planner from the planner service based on the provided request.
// If the deletion is successful, nil is returned. Otherwise, an error is returned.
func (c *PlannerControl) DeletePlanner(req *DeletePlannerRequest) error {
	if err := c.Service.DeletePlanner(req.Id); err != nil {
		return err
	}
	return nil
}

// GetPlannerRequest represents a request to get a planner by its ID.
type GetPlannerRequest struct {
	Id string `json:"id"`
}

// GetPlannerResponse is a struct that represents the response object for retrieving a planner.
type GetPlannerResponse struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

// GetPlanner retrieves a planner with the specified ID from the planner service.
// It takes a GetPlannerRequest as input, which contains the ID of the planner to retrieve.
// It returns a GetPlannerResponse, containing the ID and user ID of the retrieved planner, or an error if the retrieval fails.
func (c *PlannerControl) GetPlanner(req *GetPlannerRequest) (*GetPlannerResponse, error) {
	planner, err := c.Service.GetPlanner(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetPlannerResponse{
		Id:     planner.Id,
		UserId: planner.UserId,
	}, nil
}

// ListPlannersResponse is a response type that represents a list of planners.
// It contains an array of GetPlannerResponse structs.
type ListPlannersResponse struct {
	Planners []*GetPlannerResponse `json:"planners"`
}

// ListPlanners is a method that retrieves a list of planners from the PlannerControl service.
// It calls the ListPlanners method of the PlannerService to get the list of planners.
// It then transforms each planner into a GetPlannerResponse object and adds it to the list of plannerResponses.
// Finally, it returns a ListPlannersResponse object containing the list of plannerResponses.
// If an error occurs during the process, it is returned alongside the response.
// Example usage:
// control := &PlannerControl{service: &services.PlannerService{store: &store.MyPlannerStore{}}}
// response, err := control.ListPlanners()
func (c *PlannerControl) ListPlanners() (*ListPlannersResponse, error) {
	planners, err := c.Service.ListPlanners()
	if err != nil {
		return nil, err
	}
	var plannerResponses []*GetPlannerResponse
	for _, planner := range planners {
		plannerResponses = append(plannerResponses, &GetPlannerResponse{
			Id:     planner.Id,
			UserId: planner.UserId,
		})
	}
	return &ListPlannersResponse{
		Planners: plannerResponses,
	}, nil
}

// generateUUID generates a new UUID (Universally Unique Identifier).
// It returns the UUID as a string and any error encountered.
func generatePlannerUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
