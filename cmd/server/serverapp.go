package main

import (
	"flow/api"
	"flow/internal/conf"
	"flow/internal/inmemory"
	"flow/pkg/handle"
	"flow/pkg/services"
	"fmt"
	"github.com/asdine/storm"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func cliSetup() (*handle.TaskControl, *storm.DB, error) {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		return nil, nil, fmt.Errorf("error opening db: %s", err)
	}
	// Intialize router, service and inmemory store
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)
	return taskRouter, db, nil
}

// loggingMiddleware logs the request method, URL, and the time it took to process the request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()
	taskRouter, db, err := cliSetup()
	if err != nil {
		log.Fatalf("error setting up cli: %s", err)
	}

	defer func(db *storm.DB) {
		_ = db.Close()
	}(db)

	// Initialize handlers
	taskHandler := &api.TaskHandler{
		Control: taskRouter,
	}
	goalHandler := &api.GoalHandler{
		Control: handle.NewGoalControl(services.NewGoalService(inmemory.NewInMemoryGoalStore(db))),
	}
	planHandler := &api.PlanHandler{
		Control: handle.NewPlanControl(services.NewPlanService(inmemory.NewInMemoryPlanStore(db))),
	}
	plannerHandler := &api.PlannerHandler{
		Control: handle.NewPlannerControl(services.NewPlannerService(inmemory.NewInMemoryPlannerStore(db))),
	}
	// Register handlers and routes
	r.HandleFunc("/listtasks", taskHandler.ListTasks).Methods("GET")
	r.HandleFunc("/task/new", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/task/title/{title}", taskHandler.GetTaskByTitle).Methods("GET")
	r.HandleFunc("/task/owner/{owner}", taskHandler.GetTaskByOwner).Methods("GET")
	r.HandleFunc("/task/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", taskHandler.DeleteTask).Methods("DELETE")

	r.HandleFunc("/listgoals", goalHandler.ListGoals).Methods("GET")
	r.HandleFunc("/goal/new", goalHandler.CreateGoal).Methods("POST")
	r.HandleFunc("/goal/{id}", goalHandler.GetGoal).Methods("GET")
	r.HandleFunc("/goal/{id}", goalHandler.UpdateGoal).Methods("PUT")
	r.HandleFunc("/goal/{id}", goalHandler.DeleteGoal).Methods("DELETE")

	r.HandleFunc("/listplans", planHandler.ListPlans).Methods("GET")
	r.HandleFunc("/plan/new", planHandler.CreatePlan).Methods("POST")
	r.HandleFunc("/plan/name/{plan_name}", planHandler.GetPlanByName).Methods("GET")
	r.HandleFunc("/plan/goal/{goal_id}", planHandler.GetPlansByGoal).Methods("GET")
	r.HandleFunc("/plan/{id}", planHandler.GetPlan).Methods("GET")
	r.HandleFunc("/plan/{id}", planHandler.UpdatePlan).Methods("PUT")
	r.HandleFunc("/plan/{id}", planHandler.DeletePlan).Methods("DELETE")

	r.HandleFunc("/listplanners", plannerHandler.ListPlanners).Methods("GET")
	r.HandleFunc("/planner/new", plannerHandler.CreatePlanner).Methods("POST")
	r.HandleFunc("/planner/{id}", plannerHandler.GetPlanner).Methods("GET")
	r.HandleFunc("/planner/title/{title}", plannerHandler.GetPlannerByTitle).Methods("GET")
	r.HandleFunc("/planner/owner/{owner}", plannerHandler.GetPlannerByOwner).Methods("GET")
	r.HandleFunc("/planner/{id}", plannerHandler.UpdatePlanner).Methods("PUT")
	r.HandleFunc("/planner/{id}", plannerHandler.DeletePlanner).Methods("DELETE")
	// Apply the middleware to the router
	r.Use(loggingMiddleware)

	log.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("error serving: %s", err)
	}
}
