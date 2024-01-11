package models

import "time"

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

type Goal struct {
	Id            string    `json:"id"`
	Objective     string    `json:"objective"`
	Plans         []Plan    `json:"plans"`
	GoalStatus    string    `json:"goal_status"`
	GoalCreatedAt time.Time `json:"goal_created_at"`
	GoalUpdatedAt time.Time `json:"goal_updated_at"`
	Deadline      time.Time `json:"deadline"`
}

type Plan struct {
	Id              string    `json:"id"`
	PlanName        string    `json:"plan_name"`
	PlanDescription string    `json:"plan_description"`
	PlanDate        time.Time `json:"plan_date"`
	PlanTime        time.Time `json:"plan_time"`
	PlanStatus      string    `json:"plan_status"`
	Tasks           []Task    `json:"tasks"`
	PlanCreatedAt   time.Time `json:"plan_created_at"`
	PlanUpdatedAt   time.Time `json:"plan_updated_at"`
}
