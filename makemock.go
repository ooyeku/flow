package main

import (
	"github.com/asdine/storm"
	"goworkflow/handle"
	"goworkflow/services"
	"goworkflow/store/inmemory"
	"log"
)

func main() {
	// Open the Bolt database
	db, err := storm.Open("store/inmemory/goworkflow.db", storm.BoltOptions(0600, nil))
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
	plannerControl := handle.NewPlannerControl(plannerService)
	goalControl := handle.NewGoalControl(goalService)
	planControl := handle.NewPlanControl(planService)
	taskControl := handle.NewTaskControl(taskService)

	// Create a new planner
	plannerReq := &handle.CreatePlannerRequest{
		UserId: "456", // Replace with a valid user ID
	}
	plannerRes, err := plannerControl.CreatePlanner(plannerReq)
	if err != nil {
		log.Fatalf("Failed to create planner: %v", err)
	}

	// Create a new goal
	goalReq := &handle.CreateGoalRequest{
		PlannerId: plannerRes.ID,
	}
	goalRes, err := goalControl.CreateGoal(goalReq)
	if err != nil {
		log.Fatalf("Failed to create goal: %v", err)
	}

	// Create a new plan
	planReq := &handle.CreatePlanRequest{
		GoalId: goalRes.ID,
	}
	planRes, err := planControl.CreatePlan(planReq)
	if err != nil {
		log.Fatalf("Failed to create plan: %v", err)
	}

	// Create a new task
	taskReq := &handle.CreateTaskRequest{
		Title:       "Complete the project",
		Description: "Complete the project within the deadline",
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
