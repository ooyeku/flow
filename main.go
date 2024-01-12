package main

import (
	"goworkflow/cmd"
	"goworkflow/pkg/handle"
)

var (
	TaskRouter    *handle.TaskControl
	PlannerRouter *handle.PlannerControl
	PlanRouter    *handle.PlanControl
	GoalRouter    *handle.GoalControl
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
