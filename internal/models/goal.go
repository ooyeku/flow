package models

import "time"

// Goal represents a specific objective that a user wants to achieve.
//
//	type Goal struct {
//	    Id            string    `json:"id" storm:"id,unique"`
//	    Objective     string    `json:"objective"`
//	    Plans         []Plan    `json:"plans"`
//	    GoalStatus    string    `json:"goal_status"`
//	    GoalCreatedAt time.Time `json:"goal_created_at"`
//	    GoalUpdatedAt time.Time `json:"goal_updated_at"`
//	    Deadline      time.Time `json:"deadline"`
//	    PlannerID     string    `json:"planner_id"`
//	}
//
// The Goal struct has the following fields:
//
// - Id: Identifier of the goal. It must be unique.
// - Objective: Description of the goal's objective.
// - Plans: List of plans associated with the goal.
// - GoalStatus: Current status of the goal.
// - GoalCreatedAt: The date and time when the goal was created.
// - GoalUpdatedAt: The last time the goal was updated.
// - Deadline: The date and time when the goal should be achieved.
// - PlannerId: The identifier of the planner associated with the goal.
//
// The Goal struct is used in conjunction with the Plan struct, which represents an action plan for achieving the goal.
// Each goal can have one or more plans associated with it.
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

// GenerateGoalInstance generates a new instance of the Goal struct with the provided id, objective, and deadline. It sets the GoalStatus to "Not Started", GoalCreatedAt and GoalUpdatedAt
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

// ConvertDeadline converts a date string to a time.Time value using the format "2006-01-02". It returns the converted time value and an error if the conversion fails.
func (g *Goal) ConvertDeadline(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
