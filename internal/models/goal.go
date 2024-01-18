package models

import "time"

type Goal struct {
	Id            string    `json:"id" storm:"id,unique"`
	Objective     string    `json:"objective"`
	Plans         []Plan    `json:"plans"`
	GoalStatus    string    `json:"goal_status"`
	GoalCreatedAt time.Time `json:"goal_created_at"`
	GoalUpdatedAt time.Time `json:"goal_updated_at"`
	Deadline      time.Time `json:"deadline"`
	PlannerId     string    `json:"planner_id"`
}

func (g *Goal) GenerateGoalInstance(id, objective string, deadline time.Time) *Goal {
	return &Goal{
		Id:            id,
		Objective:     objective,
		GoalStatus:    NotStarted,
		GoalCreatedAt: time.Now(),
		GoalUpdatedAt: time.Now(),
		Deadline:      deadline,
	}
}

func (g *Goal) ConvertDeadline(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
