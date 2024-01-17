package main

import (
	"flow/internal/inmemory"
	handle2 "flow/pkg/handle"
	"flow/pkg/services"
	"github.com/asdine/storm"
	"log"
	"time"
)

func main() {
	// Open the Bolt database
	db, err := storm.Open("internal/inmemory/goworkflow.db", storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	// Create stores
	plannerStore := inmemory.NewInMemoryPlannerStore(db)
	goalStore := inmemory.NewInMemoryGoalStore(db)
	planStore := inmemory.NewInMemoryPlanStore(db)
	taskStore := inmemory.NewInMemoryTaskStore(db)

	// Create services
	plannerService := services.NewPlannerService(plannerStore)
	goalService := services.NewGoalService(goalStore)
	planService := services.NewPlanService(planStore)
	taskService := services.NewTaskService(taskStore)

	// Create controllers
	plannerControl := handle2.NewPlannerControl(plannerService)
	goalControl := handle2.NewGoalControl(goalService)
	planControl := handle2.NewPlanControl(planService)
	taskControl := handle2.NewTaskControl(taskService)

	// Create a new planner
	plannerReq := &handle2.CreatePlannerRequest{
		UserId: "456", // Replace with a valid user ID
	}
	plannerRes, err := plannerControl.CreatePlanner(plannerReq)
	if err != nil {
		log.Fatalf("Failed to create planner: %v", err)
	}

	// Create a new goal
	goalReq := &handle2.CreateGoalRequest{
		Objective: "This is an automated makemock goal",
		Deadline:  time.Now().AddDate(0, 0, 7),
		PlannerId: "000",
	}
	goalRes, err := goalControl.CreateGoal(goalReq)
	if err != nil {
		log.Fatalf("Failed to create goal: %v", err)
	}

	// Create a new plan
	planReq := &handle2.CreatePlanRequest{
		GoalId: goalRes.ID,
	}
	planRes, err := planControl.CreatePlan(planReq)
	if err != nil {
		log.Fatalf("Failed to create plan: %v", err)
	}

	// Create a new task
	taskReq := &handle2.CreateTaskRequest{
		Title:       "This is an automated makemock task",
		Description: "Useful for testing/debugging",
		Owner:       "456", // Replace with a valid user ID
	}
	taskRes, err := taskControl.CreateTask(*taskReq)
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	log.Println("Workflow created successfully")
	log.Printf("Planner ID: %s\n", plannerRes.ID)
	log.Printf("Goal ID: %s\n", goalRes.ID)
	log.Printf("Plan ID: %s\n", planRes.ID)
	log.Printf("Task ID: %s\n", taskRes.ID)
}
