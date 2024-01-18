package models

import "time"

// Plan represents a plan.
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

// GeneratePlanInstance is a method of the Plan struct that creates a new instance of a plan with the given information.
// It takes in parameters for the id, planName, planDescription, planDate, and planTime of the plan, and returns a pointer to the newly created plan instance.
// The PlanStatus is set initially to "Not Started", and the Tasks slice is initialized as an empty slice.
// The PlanCreatedAt and PlanUpdatedAt fields are set to the current time.
// The GoalId field is left empty.
//
// Example usage:
//
//	 id, err := generatePlanUUID()
//	 if err != nil {
//	   return nil, err
//	 }
//	 m := &models.Plan{}
//	 // convert planDate to time.Time
//	 planDate, err := m.ConvertPlanDate(req.PlanDate)
//	 if err != nil {
//	   return nil, err
//	 }
//	 // convert planTime to time.Time
//	 planTime, err := m.ConvertPlanTime(req.PlanTime)
//	 if err != nil {
//	   return nil, err
//	 }
//	 plan := m.GeneratePlanInstance(id, req.PlanName, req.PlanDescription, planDate, planTime)
//	 err = c.Service.CreatePlan(plan)
//	 if err != nil {
//	   return nil, err
//	 }
//	 return &CreatePlanResponse{
//	   ID: plan.Id,
//	 }, nil
//
//	 plan := m.GeneratePlanInstance(req.Id, req.PlanName, req.PlanDescription, planDate, planTime)
//	 plan.GoalId = req.GoalId
//	 return c.Service.UpdatePlan(plan)
//
//	Parameters:
//	  id (string): The unique identifier for the plan.
//	  planName (string): The name of the plan.
//	  planDescription (string): The description of the plan.
//	  planDate (time.Time): The date of the plan.
//	  planTime (time.Time): The time of the plan.
//
//	Returns:
//	  *Plan: A pointer to the newly created plan instance.
//
//	See also:
//	  - Plan
//	  - ConvertPlanDate
//	  - ConvertPlanTime
//	  - NotStarted
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

// ConvertPlanDate takes a date string and returns a time.Time value representing that date.
// The date string should be in the format "YYYY-MM-DD".
// Example usage:
//
//	planDate, err := m.ConvertPlanDate(req.PlanDate)
//	if err != nil {
//	    return nil, err
//	}
//
// Where:
//   - `m` is an instance of the `Plan` struct.
//   - `req` is an instance of the `CreatePlanRequest` struct containing the `PlanDate` field.
func (p *Plan) ConvertPlanDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// ConvertPlanTime parses a string representing time in the format "15:04" and returns the corresponding time.Time value.
// It uses the Go standard library "time" package's Parse function to perform the parsing.
// If the parsing is successful, it returns the parsed time.Time value and nil error.
// If the parsing fails, it returns a zero time.Time value and a non-nil error indicating the parsing error.
func (p *Plan) ConvertPlanTime(planTime string) (time.Time, error) {
	return time.Parse("15:04", planTime)
}
