package main

import (
	"encoding/json"
	"fmt"
	"goworkflow/internal/conf"
	handle2 "goworkflow/pkg/handle"
	"log"
)

func getAllTasks(r *handle2.TaskControl) string {
	tasks, err := r.ListTasks()
	conf.LogAndExitOnError(err, "Failed to list tasks: %v")
	taskJson, err := json.Marshal(tasks)
	conf.LogAndExitOnError(err, "Failed to marshal tasks: %v")
	log.Printf("Tasks: %s", string(taskJson))
	return string(taskJson)
}

func getAllPlanners(r *handle2.PlannerControl) string {
	planners, err := r.ListPlanners()
	conf.LogAndExitOnError(err, "Failed to list planners: %v")
	plannerJson, err := json.Marshal(planners)
	conf.LogAndExitOnError(err, "Failed to marshal planners: %v")
	log.Printf("Planners: %s", string(plannerJson))
	return string(plannerJson)
}

func getAllGoals(r *handle2.GoalControl) string {
	goals, err := r.ListGoals()
	conf.LogAndExitOnError(err, "Failed to list goals: %v")
	goalJson, err := json.Marshal(goals)
	conf.LogAndExitOnError(err, "Failed to marshal goals: %v")
	log.Printf("Goals: %s", string(goalJson))
	return string(goalJson)
}

func getAllPlans(r *handle2.PlanControl) string {
	plans, err := r.ListPlans()
	conf.LogAndExitOnError(err, "Failed to list plans: %v")
	planJson, err := json.Marshal(plans)
	conf.LogAndExitOnError(err, "Failed to marshal plans: %v")
	log.Printf("Plans: %s", string(planJson))
	return string(planJson)
}

func main() {
	taskRouter, plannerRouter, goalRouter, planRouter := conf.Setup()

	fmt.Println("Tasks:")
	getAllTasks(taskRouter)

	fmt.Println("Planners:")
	getAllPlanners(plannerRouter)

	fmt.Println("Goals:")
	getAllGoals(goalRouter)

	fmt.Println("Plans:")
	getAllPlans(planRouter)
}
