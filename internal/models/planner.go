package models

// enumerate goal status
const (
	NotStarted = "Not Started"
	InProgress = "In Progress"
	Completed  = "Completed"
	Fail       = "Fail"
)

type Planner struct {
	Id     string `json:"id" storm:"id,unique"`
	UserId string `json:"user_id"`
	Goals  []Goal `json:"goals"`
}

func (p *Planner) GeneratePlannerInstance(id, userId string) *Planner {
	return &Planner{
		Id:     id,
		UserId: userId,
	}
}
