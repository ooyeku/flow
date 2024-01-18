package models

import "time"

type Plan struct {
	Id              string    `json:"id" storm:"id,unique"`
	PlanName        string    `json:"plan_name"`
	PlanDescription string    `json:"plan_description"`
	PlanDate        time.Time `json:"plan_date"`
	PlanTime        time.Time `json:"plan_time"`
	PlanStatus      string    `json:"plan_status"`
	Tasks           []Task    `json:"tasks"`
	PlanCreatedAt   time.Time `json:"plan_created_at"`
	PlanUpdatedAt   time.Time `json:"plan_updated_at"`
	GoalId          string    `json:"goal_id"`
}

func (p *Plan) GeneratePlanInstance(id, planName, planDescription string, planDate, planTime time.Time) *Plan {
	return &Plan{
		Id:              id,
		PlanName:        planName,
		PlanDescription: planDescription,
		PlanDate:        planDate,
		PlanTime:        planTime,
		PlanStatus:      NotStarted,
		Tasks:           []Task{},
		PlanCreatedAt:   time.Now(),
		PlanUpdatedAt:   time.Now(),
		GoalId:          "",
	}
}

func (p *Plan) ConvertPlanDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func (p *Plan) ConvertPlanTime(planTime string) (time.Time, error) {
	return time.Parse("15:04", planTime)
}
