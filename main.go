package main

import (
	"encoding/json"
	"goworkflow/conf"
	"goworkflow/handle"
	"log"
)

func getAllTasks(r *handle.TaskControl) string {
	tasks, err := r.ListTasks()
	conf.LogAndExitOnError(err, "Failed to list tasks: %v")
	taskJson, err := json.Marshal(tasks)
	conf.LogAndExitOnError(err, "Failed to marshal tasks: %v")
	log.Printf("Tasks: %s", string(taskJson))
	return string(taskJson)
}

func getAllPlanners(r *handle.PlannerControl) string {
	planners, err := r.ListPlanners()
	conf.LogAndExitOnError(err, "Failed to list planners: %v")
	plannerJson, err := json.Marshal(planners)
	conf.LogAndExitOnError(err, "Failed to marshal planners: %v")
	log.Printf("Planners: %s", string(plannerJson))
	return string(plannerJson)
}

func main() {
	taskRouter, planRouter := conf.Setup()
	getAllTasks(taskRouter)
	getAllPlanners(planRouter)
}
