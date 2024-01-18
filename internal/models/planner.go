package models

// Constants representing the different states of a task
const (
	NotStarted = "Not Started"
	InProgress = "In Progress"
	Completed  = "Completed"
	Fail       = "Fail"
)

// Planner represents a planner object with its attributes.
type Planner struct {
	Id     string `json:"id" storm:"id,unique"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
	Goals  []Goal `json:"goals"`
}

// GeneratePlannerInstance generates a new instance of Planner with the given id, title, and userId.
// It returns a pointer to the newly created Planner instance.
// Example usage:
//
//	m := &models.Planner{}
//	planner := m.GeneratePlannerInstance(id, req.Title, req.UserId)
//	err = c.Service.CreatePlanner(planner)
//	if err != nil {
//	  return nil, err
//	}
//	return &CreatePlannerResponse{
//	  ID: planner.Id,
//	}, nil
func (p *Planner) GeneratePlannerInstance(id, title string, userId string) *Planner {
	return &Planner{
		Id:     id,
		Title:  title,
		UserId: userId,
	}
}
