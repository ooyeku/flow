package main

import (
	"goworkflow/cmd"
	"goworkflow/internal/conf"
	"goworkflow/pkg/handle"
)

var (
	taskRouter    *handle.TaskControl
	plannerRouter *handle.PlannerControl
	planRouter    *handle.PlanControl
	goalRouter    *handle.GoalControl
)

func main() {
	// initialize routers with configuration
	var err error
	taskRouter, plannerRouter, goalRouter, planRouter = conf.Setup()
	if err != nil {
		return
	}

	err = cmd.Execute()
	if err != nil {
		return
	}

	//taskRouter, plannerRouter, goalRouter, planRouter := conf.Setup()
	//
	//fmt.Println("Tasks:")
	//internal.GetAllTasks(taskRouter)
	//
	//fmt.Println("Planners:")
	//internal.GetAllPlanners(plannerRouter)
	//
	//fmt.Println("Goals:")
	//internal.GetAllGoals(goalRouter)
	//
	//fmt.Println("Plans:")
	//internal.GetAllPlans(planRouter)
}
