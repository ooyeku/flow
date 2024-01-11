package main

import (
	"github.com/asdine/storm"
	"goworkflow/handle"
	"goworkflow/models"
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
	defer db.Close()

	// Create a new instance of the PlannerStore using the Bolt database
	plannerStore := inmemory.NewInMemoryPlannerStore(db)

	// Create a new instance of the PlannerService using the PlannerStore
	plannerService := services.NewPlannerService(plannerStore)

	// Create a new instance of the PlannerControl using the PlannerService
	plannerControl := handle.NewPlannerControl(plannerService)

	// Create a new planner with goals
	planner := &models.Planner{
		Id:     "123", // Replace with a valid UUID
		UserId: "456", // Replace with a valid user ID
		Goals: []models.Goal{
			{
				Id:         "789", // Replace with a valid UUID
				Objective:  "Complete the project",
				Plans:      nil, // Add plans as needed
				GoalStatus: models.NotStarted,
			},
		},
	}

	// Create a request to create a planner
	req := &handle.CreatePlannerRequest{
		UserId: planner.UserId,
	}

	// Add the new planner to the database using the CreatePlanner method of the PlannerControl
	_, err = plannerControl.CreatePlanner(req)
	if err != nil {
		log.Fatalf("Failed to create planner: %v", err)
	}

	log.Println("Planner created successfully")
}
