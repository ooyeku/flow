package main

import (
	"goworkflow/cmd"
)

func main() {
	err := cmd.Execute()
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
