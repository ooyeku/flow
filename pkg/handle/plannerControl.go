package handle

import (
	"flow/internal/models"
	"flow/pkg/services"
	"github.com/google/uuid"
)

// PlannerControl represents a controller that provides methods to manage planners.
type PlannerControl struct {
	Service *services.PlannerService
}

// NewPlannerControl initializes a new PlannerControl struct with the given PlannerService instance as its Service field.
// It returns a pointer to the newly created PlannerControl.
func NewPlannerControl(service *services.PlannerService) *PlannerControl {
	return &PlannerControl{
		Service: service,
	}
}

// CreatePlannerRequest represents a request to create a planner.
// It contains the title and user ID of the planner.
type CreatePlannerRequest struct {
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// CreatePlannerResponse represents the response structure for creating a planner. It contains the ID of the created planner.
type CreatePlannerResponse struct {
	ID string `json:"id"`
}

// CreatePlanner creates a new planner based on the provided request.
// It generates a unique ID for the planner using the generatePlannerUUID function.
// Then, it creates a new instance of the Planner struct with the generated ID, request title, and request user ID.
// Next, it calls the CreatePlanner method of the PlannerService to persist the planner in the database.
// If any error occurs during the process, it returns nil and the error.
// Otherwise, it returns a CreatePlannerResponse with the ID of the created planner and nil error.
func (c *PlannerControl) CreatePlanner(req *CreatePlannerRequest) (*CreatePlannerResponse, error) {
	id, err := generatePlannerUUID()
	if err != nil {
		return nil, err
	}
	m := &models.Planner{}
	planner := m.GeneratePlannerInstance(id, req.Title, req.UserId)
	err = c.Service.CreatePlanner(planner)
	if err != nil {
		return nil, err
	}
	return &CreatePlannerResponse{
		ID: planner.Id,
	}, nil
}

// UpdatePlannerRequest represents a request to update a planner with new title and user ID. It contains the following fields:
// - Id: the ID of the planner to update
// - Title: the new title for the planner
// - UserId: the new user ID for the planner
type UpdatePlannerRequest struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// UpdatePlanner updates the title and user ID of a planner based on the provided request.
func (c *PlannerControl) UpdatePlanner(req *UpdatePlannerRequest) error {
	planner, err := c.Service.GetPlanner(req.Id)
	if err != nil {
		return err
	}
	planner.UserId = req.UserId
	planner.Title = req.Title
	if err := c.Service.UpdatePlanner(planner); err != nil {
		return err
	}
	return nil
}

// DeletePlannerRequest represents a request to delete a planner.
// It contains the ID of the planner to be deleted.
type DeletePlannerRequest struct {
	Id string `json:"id"`
}

// DeletePlanner removes a planner from the system based on the provided request.
// It calls the DeletePlanner method of the Service, passing the Id from the request.
// If an error occurs during the deletion process, it returns the error. Otherwise, it returns nil.
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

// GetPlannerResponse represents the response object when retrieving a planner.
// It contains the ID, title, and user ID of the planner.
// This type is used in the following methods:
// - PlannerControl.GetPlanner
// - PlannerControl.ListPlanners
type GetPlannerResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// GetPlanner retrieves a planner based on the provided request ID.
// It returns a GetPlannerResponse object containing the planner ID, title, and user ID.
// If an error occurs during the retrieval process, it returns nil and the error message.
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

// GetPlannerByTitleRequest represents a request to get a planner by its title.
// It contains the following field:
// - Title: The title of the planner to get.
type GetPlannerByTitleRequest struct {
	Title string `json:"title"`
}

// GetPlannerByTitleResponse represents the response structure for retrieving a planner by title.
type GetPlannerByTitleResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// GetPlannerByTitle retrieves a planner by its title based on the provided request.
// It calls the GetPlannerByTitle method on the PlannerService and returns the corresponding planner
// ID, Title, and UserId as a GetPlannerByTitleResponse object.
// If an error occurs during the retrieval process, it returns nil and the error.
func (c *PlannerControl) GetPlannerByTitle(req *GetPlannerByTitleRequest) (*GetPlannerByTitleResponse, error) {
	planner, err := c.Service.GetPlannerByTitle(req.Title)
	if err != nil {
		return nil, err
	}
	return &GetPlannerByTitleResponse{
		Id:     planner.Id,
		UserId: planner.UserId,
	}, nil
}

// GetPlannerByOwnerRequest represents a request for getting a planner by owner.
type GetPlannerByOwnerRequest struct {
	UserId string `json:"user_id"`
}

// GetPlannerByOwnerResponse represents the response when getting a planner by owner.
type GetPlannerByOwnerResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// GetPlannerByOwner retrieves the planner(s) owned by the specified user based on the provided request.
// Parameters:
//   - req: The request object that contains the user ID.
//
// Returns:
//   - A pointer to GetPlannerByOwnerResponse: The response object that contains the ID and user ID of the retrieved planner(s).
//   - An error if there was an issue retrieving the planner(s).
//
// Example usage:
//
//	plannerControl := &PlannerControl{}
//	request := &GetPlannerByOwnerRequest{
//	  UserId: "123456",
//	}
//	response, err := plannerControl.GetPlannerByOwner(request)
//	if err != nil {
//	  fmt.Println("Error:", err)
//	  return
//	}
//	fmt.Println("ID:", response.Id)
//	fmt.Println("User ID:", response.UserId)
//
//	This will output:
//	ID: abcdefg
//	User ID: 123456
func (c *PlannerControl) GetPlannerByOwner(req *GetPlannerByOwnerRequest) (*GetPlannerByOwnerResponse, error) {
	planners, err := c.Service.GetPlannerByOwner(req.UserId)
	if err != nil {
		return nil, err
	}
	// get all planner from planners
	var plannerResponses []*GetPlannerByOwnerResponse
	for _, planner := range planners {
		plannerResponses = append(plannerResponses, &GetPlannerByOwnerResponse{
			Id:     planner.Id,
			UserId: planner.UserId,
		})
	}
	return &GetPlannerByOwnerResponse{
		Id:     plannerResponses[0].Id,
		UserId: plannerResponses[0].UserId,
	}, nil
}

// ListPlannersResponse represents a response with a list of planners.
// The Planners field is an array of GetPlannerResponse objects.
type ListPlannersResponse struct {
	Planners []*GetPlannerResponse `json:"planners"`
}

// ListPlanners retrieves a list of all planners.
// It calls the ListPlanners method of the PlannerService to get the list of planners from the store.
// It then converts each planner to a GetPlannerResponse and appends it to the plannerResponses slice.
// Finally, it returns a ListPlannersResponse containing the list of GetPlannerResponse objects.
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

// generatePlannerUUID generates a new UUID (Universally Unique Identifier)
// using the uuid.NewRandom() function.
// It returns the generated UUID as a string and any potential error.
func generatePlannerUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
