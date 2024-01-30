package main

import (
	"bufio"
	"flow/internal/conf"
	"flow/internal/inmemory"
	"flow/pkg/handle"
	"flow/pkg/services"
	"fmt"
	"github.com/asdine/storm"
	"log"
	"os"
	"strings"
)

// main function is the entry point of the CLI application.
// It prompts the user to enter a command, reads the command from standard input,
// and passes it to the runCommand function for execution.
// If there is an error reading or executing the command, it prints the error message to standard error.
// It keeps looping until the user enters the "exit" command to exit the application.
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to your CLI app! 😼")
	for {
		fmt.Println("Enter command: ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
		err = runCommand(cmdString)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
	}
}

// cliSetup initializes the router, service, and in-memory store. It opens the database using the dbPath obtained from conf.GetDBPath(). If there is an error opening the database, it
func cliSetup() (*handle.TaskControl, *handle.GoalControl, *handle.PlanControl, *handle.PlannerControl, *storm.DB) {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("error opening db: %s", err)
	}

	// Intialize router, service and inmemory store
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)

	goalStore := inmemory.NewInMemoryGoalStore(db)
	goalService := services.NewGoalService(goalStore)
	goalRouter := handle.NewGoalControl(goalService)

	planStore := inmemory.NewInMemoryPlanStore(db)
	planService := services.NewPlanService(planStore)
	planRouter := handle.NewPlanControl(planService)

	plannerStore := inmemory.NewInMemoryPlannerStore(db)
	plannerService := services.NewPlannerService(plannerStore)
	plannerRouter := handle.NewPlannerControl(plannerService)

	return taskRouter, goalRouter, planRouter, plannerRouter, db
}

func promptUser(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read from stdin: %s", err)
	}
	return strings.TrimSpace(response), nil
}

// taskCommands is a map that contains various commands related to task operations. The key represents the command name and the value represents the corresponding function to be executed
var taskCommands = map[string]func(*handle.TaskControl){
	"create-task":       createTask,
	"get-task":          getTask,
	"get-task-by-title": getTaskByTitle,
	"get-task-by-owner": getTaskByOwner,
	"update-task":       updateTasks,
	"delete-task":       deleteTask,
	"list-tasks":        listTasks,
}

// goalCommands is a map that maps command names to their corresponding functions in the handle.GoalControl struct.
// The handle.GoalControl struct is responsible for handling goal-related requests and executing the corresponding actions.
// The functions listed below are the available commands that can be executed using goalCommands:
//   - create-goal: This command creates a new goal by reading input from the user and making a request to the CreateGoal function in the GoalControl struct.
//     It prompts the user to enter the goal objective, deadline, and planner ID.
//   - get-goal: This command retrieves a goal by its ID by reading input from the user and making a request to the GetGoal function in the GoalControl struct.
//     It prompts the user to enter the goal ID.
//   - get-goal-by-title: This command retrieves a goal by its objective by reading input from the user and making a request to the GetGoalByObjective function in the GoalControl struct
var goalCommands = map[string]func(*handle.GoalControl){
	"create-goal":       createGoal,
	"get-goal":          getGoal,
	"get-goal-by-title": getGoalByObjective,
	"get-goal-by-owner": getGoalsByPlannerId,
	"update-goal":       updateGoals,
	"delete-goal":       deleteGoal,
	"list-goals":        listGoals,
}

// planCommands is a map that maps command names to their corresponding functions
var planCommands = map[string]func(*handle.PlanControl){
	"create-plan":      createPlan,
	"get-plan":         getPlan,
	"get-plan-by-name": getPlanByName,
	"get-plan-by-goal": getPlanByGoal,
	"update-plan":      updatePlans,
	"delete-plan":      deletePlan,
	"list-plans":       listPlans,
}

// plannerCommands is a map that associates string keys with functions that operate on a *handle.PlannerControl object.
// Each key represents a command, and the associated function performs the corresponding action.
var plannerCommands = map[string]func(*handle.PlannerControl){
	"create-planner":       createPlanner,
	"get-planner":          getPlanner,
	"get-planner-by-title": getPlannerByGoal,
	"get-planner-by-owner": getPlannerByOwner,
	"update-planner":       updatePlanners,
	"delete-planner":       deletePlanner,
	"list-planners":        listPlanners,
}

// createTask is a function that prompts the user to enter the details of a task,
// creates the task using the provided TaskControl instance, and prints the ID of the created task.
//
// Parameters:
// - t: *handle.TaskControl - An instance of TaskControl that provides the CreateTask method for task creation.
//
// Example usage:
//
//	taskRouter, db := cliSetup()
//	defer db.Close()
//	createTask(taskRouter)
//
// This function utilizes the TaskControl instance to create a task, by:
// 1. Printing a prompt for the task title and reading the user input.
// 2. Printing a prompt for the task description and reading the user input.
// 3. Printing a prompt for the task owner and reading the user input.
// 4. Creating a CreateTaskRequest instance with the title, description, and owner obtained from user input.
// 5. Invoking the CreateTask method of the TaskControl instance with the CreateTaskRequest instance.
// 6. If an error occurs during task creation, printing the error message.
// 7. Printing the ID of the created task.
func createTask(t *handle.TaskControl) {
	fmt.Println("Creating task...")
	reader := bufio.NewReader(os.Stdin)
	title, err := promptUser(reader, "Enter task title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	description, err := promptUser(reader, "Enter task description: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}

	owner, err := promptUser(reader, "Enter task owner: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.CreateTaskRequest{
		Title:       title,
		Description: description,
		Owner:       owner,
	}
	res, err := t.CreateTask(req)
	if err != nil {
		fmt.Println("Error creating task: ", err)
	}
	fmt.Println("Created task with id: ", res.ID)
}

// getTask retrieves a task with a given ID by calling the GetTask method of the TaskControl struct.
// It prompts the user to enter the task ID, sends a GetTaskRequest to the TaskControl, and displays the task details if found.
// Otherwise, it prints an error message if the task is not found or an error occurs.
// The function uses the promptUser function to read user input from stdin.
func getTask(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter task id: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

// getTaskByTitle prompts the user to enter a task title and retrieves the task with that title using the provided TaskControl instance.
// It prints the task title and description if the task is found.
// If there is an error reading from stdin or retrieving the task, an error message is logged and the function returns.
func getTaskByTitle(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	title, err := promptUser(reader, "Enter task title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetTaskByTitleRequest{
		Title: title,
	}
	task, err := t.GetTaskByTitle(&req)
	if err != nil {
		// error message is logged in GetTaskByTitle
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

func getTaskByOwner(t *handle.TaskControl) {
	fmt.Println("Getting task...")
	reader := bufio.NewReader(os.Stdin)
	owner, err := promptUser(reader, "Enter task owner: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
		return
	}
	req := handle.GetTaskByOwnerRequest{
		Owner: owner,
	}
	tasks, err := t.GetTaskByOwner(&req)
	if err != nil {
		// error message is logged in GetTaskByOwner
		return
	}
	for _, task := range tasks {
		fmt.Println("Got task: ", task.Title)
		fmt.Println("Description: ", task.Description)
	}
}

// updateTasks prompts the user to enter a task ID, then retrieves the task with the specified ID using the GetTask function from the provided TaskControl object.
// It then prompts the user to enter a new task title, description, and owner.
// After that, it constructs an UpdateTaskRequest object with the provided ID, new title, description, owner, and default values for started and completed.
// Finally, it calls the UpdateTask function from the TaskControl object to update the task with the new values.
// If any error occurs during the process, it prints an error message.
// Example usage:
// t := &handle.TaskControl{service: taskService}
// updateTasks(t)
func updateTasks(t *handle.TaskControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter task id of task to be updated: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
	fmt.Println("Enter New task title: ")

	title, err := promptUser(reader, "Enter task title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	description, err := promptUser(reader, "Enter task description: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	owner, err := promptUser(reader, "Enter task owner: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	update := handle.UpdateTaskRequest{
		ID:          id,
		Title:       title,
		Description: description,
		Owner:       owner,
	}
	fmt.Println("Updating task...")
	err = t.UpdateTask(&update)
	if err != nil {
		fmt.Printf("Error updating task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Updated task with id: ", task.ID)
}

// deleteTask prompts the user to enter the ID of the task to be deleted. It retrieves the task with the given ID using the TaskControl.GetTask function.
// It then displays the details of the task to the user and asks for confirmation to delete the task. If the user confirms, it deletes the task using the TaskControl.DeleteTask function
func deleteTask(t *handle.TaskControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter task id to be deleted: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	// Get task first to show user what task is being deleted
	req := handle.GetTaskRequest{
		ID: id,
	}
	task, err := t.GetTask(&req)
	if err != nil {
		fmt.Printf("Error getting task with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("are you sure you want to delete this task? (y/n)")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	confirm = strings.TrimSpace(confirm)
	if confirm == "n" {
		fmt.Println("Task not deleted")
	} else if confirm == "y" {
		fmt.Println("Deleting task...")
		req := handle.DeleteTaskRequest{
			ID: id,
		}
		err = t.DeleteTask(&req)
		if err != nil {
			fmt.Printf("Error deleting task with id %s: %s\n", id, err)
			return
		}
		fmt.Println("Deleted task with id: ", task.ID)
	} else {
		fmt.Println("Invalid input")
	}
}

// listTasks lists all tasks using the given TaskControl instance.
// It fetches all tasks using the ListTasks() method and prints the task ID, title, and description for each task.
// If there is an error in fetching the tasks, it prints the error message.
func listTasks(t *handle.TaskControl) {
	fmt.Println("Listing tasks...")
	tasks, err := t.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks: ", err)
		return
	}
	// get task id and title of each task
	for _, task := range tasks {
		fmt.Printf("Task id: %s, Title: %s, Description %s\n", task.ID, task.Title, task.Description)
	}
}

// createGoal prints a prompt to the user, reads input from stdin, and creates a new goal using the provided GoalControl.
// promptUser reads input from the user with the given prompt message and returns the user input as a string.
// handle.CreateGoalRequest is a struct that represents the data needed to create a goal. It contains the objective, deadline, and planner ID.
// handle.GoalControl is a struct that provides methods for creating, updating, deleting, and retrieving goals.
// services.GoalService is a struct that implements the logic for goal creation, update, deletion, and retrieval.
// CreateGoalResponse is a struct that contains the ID of the newly created goal.
// generateGoalUUID generates a new UUID for the goal ID.
// models.Goal is a struct that represents a goal and provides helper methods for goal instance creation and deadline conversion.
// Usage example:
//
//	g := &handle.GoalControl{
//	  Service: &services.GoalService{store: store2.GoalStore{}},
//	}
//	createGoal(g)
func createGoal(g *handle.GoalControl) {
	fmt.Println("Creating goal...")
	reader := bufio.NewReader(os.Stdin)
	objective, err := promptUser(reader, "Enter goal objective: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	deadline, err := promptUser(reader, "Enter goal deadline in YYYY-MM-DD format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	plannerid, err := promptUser(reader, "Enter goal plannerid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}

	req := handle.CreateGoalRequest{
		Objective: objective,
		Deadline:  deadline,
		PlannerId: plannerid,
	}
	res, err := g.CreateGoal(&req)
	if err != nil {
		fmt.Println("Error creating goal: ", err)
	}
	fmt.Println("Created goal with id: ", res.ID)
}

// getGoal takes a pointer to a GoalControl and prompts the user to enter a goal id. It then calls the GetGoal method on the GoalControl to retrieve the goal information associated with
func getGoal(g *handle.GoalControl) {
	fmt.Println("Getting goal...")
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter goal id: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetGoalRequest{
		Id: id,
	}
	goal, err := g.GetGoal(&req)
	if err != nil {
		fmt.Printf("Error getting goal with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got goal: ", goal.Goal.Objective)
	fmt.Println("Deadline: ", goal.Goal.Deadline)
	fmt.Println("PlannerID: ", goal.Goal.PlannerId)
}
func getGoalByObjective(g *handle.GoalControl) {
	fmt.Println("Getting goal...")
	reader := bufio.NewReader(os.Stdin)
	objective, err := promptUser(reader, "Enter goal objective: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetGoalByObjectiveRequest{
		Objective: objective,
	}
	goal, err := g.GetGoalByObjective(&req)
	if err != nil {
		// error message is logged in GetGoalByObjective
		return
	}
	fmt.Println("Got goal: ", goal.Goal.Objective)
	fmt.Println("Deadline: ", goal.Goal.Deadline)
	fmt.Println("PlannerID: ", goal.Goal.PlannerId)
}
func getGoalsByPlannerId(g *handle.GoalControl) {
	fmt.Println("Getting goal...")
	reader := bufio.NewReader(os.Stdin)
	plannerid, err := promptUser(reader, "Enter goal plannerid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetGoalsByPlannerIdRequest{
		PlannerId: plannerid,
	}
	goals, err := g.GetGoalsByPlannerId(&req)
	if err != nil {
		// error message is logged in GetGoalByObjective
		return
	}
	for _, goal := range goals.Goals {
		fmt.Println("Got goal: ", goal.Objective)
		fmt.Println("Deadline: ", goal.Deadline)
		fmt.Println("PlannerID: ", goal.PlannerId)
	}
}

// updateGoals updates a goal by taking input from the user. It prompts the user for the goal ID, retrieves the goal from the GoalControl using the GetGoal method, and then prompts the
func updateGoals(g *handle.GoalControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter goal id of goal to be updated: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	req := handle.GetGoalRequest{
		Id: id,
	}
	goal, err := g.GetGoal(&req)
	if err != nil {
		fmt.Printf("Error getting goal with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got goal: ", goal.Goal.Objective)
	fmt.Println("Deadline: ", goal.Goal.Deadline)
	fmt.Println("PlannerID: ", goal.Goal.PlannerId)
	fmt.Println("Enter New goal objective: ")

	objective, err := promptUser(reader, "Enter goal objective: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	deadline, err := promptUser(reader, "Enter goal deadline in YYYY-MM-DD format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	plannerid, err := promptUser(reader, "Enter goal plannerid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	update := handle.UpdateGoalRequest{
		Id:        id,
		Objective: objective,
		Deadline:  deadline,
		PlannerId: plannerid,
	}
	fmt.Println("Updating goal...")
	err = g.UpdateGoal(&update)
	if err != nil {
		fmt.Printf("Error updating goal with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Updated goal with id: ", goal.Goal.Id)
}

// deleteGoal prompts the user for a goal id to delete and then proceeds to delete the goal.
// It first gets the goal from the provided GoalControl using the GetGoal method.
// It then asks the user for confirmation to delete the goal and proceeds with deleting it.
// If the user confirms, it calls the DeleteGoal method of GoalControl to delete the goal.
// It prints the id of the deleted goal if successful.
// Parameters:
// - g: the handle.GoalControl instance used to interact with the goal service
func deleteGoal(g *handle.GoalControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter goal id to be deleted: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	// Get goal first to show user what goal is being deleted
	req := handle.GetGoalRequest{
		Id: id,
	}
	goal, err := g.GetGoal(&req)
	if err != nil {
		fmt.Printf("Error getting goal with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got goal: ", goal.Goal.Objective)
	fmt.Println("are you sure you want to delete this goal? (y/n)")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	confirm = strings.TrimSpace(confirm)
	if confirm == "n" {
		fmt.Println("Goal not deleted")
	} else if confirm == "y" {
		fmt.Println("Deleting goal...")
		req := handle.DeleteGoalRequest{
			Id: id,
		}
		err = g.DeleteGoal(&req)
		if err != nil {
			fmt.Printf("Error deleting goal with id %s: %s\n", id, err)
			return
		}
		fmt.Println("Deleted goal with id: ", goal.Goal.Id)
	} else {
		fmt.Println("Invalid input")
	}
}

// Example usage:
// g := &handle.GoalControl{}
// listGoals(g)
// Output:
// Listing goals...
// Goal id: {goalId}, Objective: {objective}, Deadline {deadline}
// Goal id: {goalId}, Objective: {objective}, Deadline {deadline}
// ...
//
// Goal id represents the unique identifier of the goal
// Objective represents the objective of the goal
// Deadline represents the deadline of the goal
func listGoals(g *handle.GoalControl) {
	fmt.Println("Listing goals...")
	goals, err := g.ListGoals()
	if err != nil {
		fmt.Println("Error listing goals: ", err)
		return
	}
	// get goal id and objective of each goal
	for _, goal := range goals.Goals {
		fmt.Printf("Goal id: %s, Objective: %s, Deadline %s\nPlannerID: %s\nPlans: %v\n", goal.Id, goal.Objective, goal.Deadline, goal.PlannerId, goal.Plans)
	}
}

func createPlan(p *handle.PlanControl) {
	fmt.Println("Creating plan...")
	reader := bufio.NewReader(os.Stdin)
	name, err := promptUser(reader, "Enter plan name: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	description, err := promptUser(reader, "Enter plan description: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	date, err := promptUser(reader, "Enter plan date in YYYY-MM-DD format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	time, err := promptUser(reader, "Enter plan time in HH:MM format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}

	req := handle.CreatePlanRequest{
		PlanName:        name,
		PlanDescription: description,
		PlanDate:        date,
		PlanTime:        time,
	}
	res, err := p.CreatePlan(&req)
	if err != nil {
		fmt.Println("Error creating plan: ", err)
	}
	fmt.Println("Created plan with id: ", res.ID)
}

func getPlan(p *handle.PlanControl) {
	fmt.Println("Getting plan...")
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter plan id: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlanRequest{
		Id: id,
	}
	plan, err := p.GetPlan(&req)
	if err != nil {
		fmt.Printf("Error getting plan with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got plan: ", plan.Plan.PlanName)
	fmt.Println("Description: ", plan.Plan.PlanDescription)
}

func getPlanByName(p *handle.PlanControl) {
	fmt.Println("Getting plan...")
	reader := bufio.NewReader(os.Stdin)
	name, err := promptUser(reader, "Enter plan name: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlanByNameRequest{
		PlanName: name,
	}
	plan, err := p.GetPlanByName(&req)
	if err != nil {
		// error message is logged in GetPlanByTitle
		return
	}
	fmt.Println("Got plan: ", plan.Plan.PlanName)
	fmt.Println("Description: ", plan.Plan.PlanDescription)
}

func getPlanByGoal(p *handle.PlanControl) {
	fmt.Println("Getting plan...")
	reader := bufio.NewReader(os.Stdin)
	goalid, err := promptUser(reader, "Enter plan goalid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlansByGoalRequest{
		GoalId: goalid,
	}
	plans, err := p.GetPlansByGoal(&req)
	if err != nil {
		// error message is logged in GetPlanByTitle
		log.Fatalf("Error getting plans by goal: %s", err)
	}
	for _, plan := range plans.Plans {
		fmt.Println("Got plan: ", plan.PlanName)
		fmt.Println("Description: ", plan.PlanDescription)
	}
}

func updatePlans(p *handle.PlanControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter plan id of plan to be updated: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	req := handle.GetPlanRequest{
		Id: id,
	}
	plan, err := p.GetPlan(&req)
	if err != nil {
		fmt.Printf("Error getting plan with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got plan: ", plan.Plan.PlanName)
	fmt.Println("Description: ", plan.Plan.PlanDescription)
	fmt.Println("Enter New plan name: ")

	name, err := promptUser(reader, "Enter plan name: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	description, err := promptUser(reader, "Enter plan description: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	date, err := promptUser(reader, "Enter plan date in YYYY-MM-DD format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	time, err := promptUser(reader, "Enter plan time in HH:MM format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	update := handle.UpdatePlanRequest{
		Id:              id,
		PlanName:        name,
		PlanDescription: description,
		PlanDate:        date,
		PlanTime:        time,
	}
	fmt.Println("Updating plan...")
	err = p.UpdatePlan(&update)
	if err != nil {
		fmt.Printf("Error updating plan with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Updated plan with id: ", plan.Plan.Id)
}

func deletePlan(p *handle.PlanControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter plan id to be deleted: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	// Get plan first to show user what plan is being deleted
	req := handle.GetPlanRequest{
		Id: id,
	}
	plan, err := p.GetPlan(&req)
	if err != nil {
		fmt.Printf("Error getting plan with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got plan: ", plan.Plan.PlanName)
	fmt.Println("are you sure you want to delete this plan? (y/n)")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	confirm = strings.TrimSpace(confirm)
	if confirm == "n" {
		fmt.Println("Plan not deleted")
	} else if confirm == "y" {
		fmt.Println("Deleting plan...")
		req := handle.DeletePlanRequest{
			Id: id,
		}
		err = p.DeletePlan(&req)
		if err != nil {
			fmt.Printf("Error deleting plan with id %s: %s\n", id, err)
			return
		}
		fmt.Println("Deleted plan with id: ", plan.Plan.Id)
	} else {
		fmt.Println("Invalid input")
	}
}

// listPlans retrieves a list of plans from the PlanControl and prints them to the console.
// It first prints a message indicating that the plans are being listed, then calls the ListPlans method of the PlanControl
// to get the list of plans. It then iterates over the list and prints the ID, name, and description of each plan.
func listPlans(p *handle.PlanControl) {
	fmt.Println("Listing plans...")
	plans, err := p.ListPlans()
	if err != nil {
		fmt.Println("Error listing plans: ", err)
		return
	}
	// get plan id and name of each plan
	for _, plan := range plans.Plans {
		fmt.Printf("Plan id: %s, Name: %s, Description %s\n", plan.Id, plan.PlanName, plan.PlanDescription)
	}
}

// createPlanner prompts the user to enter a planner title and userid, and then uses the PlannerControl service to create and store the planner. It prints the ID of the created planner
func createPlanner(p *handle.PlannerControl) {
	fmt.Println("Creating planner...")
	reader := bufio.NewReader(os.Stdin)
	title, err := promptUser(reader, "Enter planner title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	userid, err := promptUser(reader, "Enter planner userid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}

	req := handle.CreatePlannerRequest{
		Title:  title,
		UserId: userid,
	}
	res, err := p.CreatePlanner(&req)
	if err != nil {
		fmt.Println("Error creating planner: ", err)
	}
	fmt.Println("Created planner with id: ", res.Id)
}

// getPlanner prompts the user to enter a planner ID and retrieves the planner from PlannerControl if it exists.
//
// The function takes a *handle.PlannerControl as input.
//
// It first prompts the user to enter a planner ID using the promptUser function.
// If there is an error reading from stdin, it logs a fatal error and exits.
//
// Then, it creates a handle.GetPlannerRequest with the entered ID and calls p.GetPlanner to get the planner.
// If there is an error getting the planner, it prints an error message and returns.
//
// Finally, it prints details about the retrieved planner, including the title and user ID.
//
// Example usage:
//
//	p := &handle.PlannerControl{...}
//	getPlanner(p)
//
//	Output:
//	  Getting planner...
//	  Enter planner id: 1
//	  Got planner: My Planner
//	  User ID: 123
func getPlanner(p *handle.PlannerControl) {
	fmt.Println("Getting planner...")
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter planner id: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlannerRequest{
		Id: id,
	}
	planner, err := p.GetPlanner(&req)
	if err != nil {
		fmt.Printf("Error getting planner with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got planner: ", planner.Title)
	fmt.Println("User ID: ", planner.UserId)
}

func getPlannerByGoal(p *handle.PlannerControl) {
	fmt.Println("Getting planner...")
	reader := bufio.NewReader(os.Stdin)
	title, err := promptUser(reader, "Enter planner title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlannerByTitleRequest{
		Title: title,
	}
	planner, err := p.GetPlannerByTitle(&req)
	if err != nil {
		// error message is logged in GetPlannerByTitle
		return
	}
	fmt.Println("Got planner: ", planner.Title)
	fmt.Println("User ID: ", planner.UserId)
}

func getPlannerByOwner(p *handle.PlannerControl) {
	fmt.Println("Getting planner...")
	reader := bufio.NewReader(os.Stdin)
	userid, err := promptUser(reader, "Enter planner userid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetPlannerByOwnerRequest{
		UserId: userid,
	}
	planners, err := p.GetPlannerByOwner(&req)
	if err != nil {
		// error message is logged in GetPlannerByTitle
		return
	}
	for _, planner := range planners {
		fmt.Println("Got planner: ", planner.Title)
		fmt.Println("User ID: ", planner.UserId)
	}
}

// updatePlanners takes a PlannerControl pointer as input. It prompts the user to enter the ID of the planner to be updated and retrieves the planner using the GetPlanner method from
func updatePlanners(p *handle.PlannerControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter planner id of planner to be updated: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	req := handle.GetPlannerRequest{
		Id: id,
	}
	planner, err := p.GetPlanner(&req)
	if err != nil {
		fmt.Printf("Error getting planner with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got planner: ", planner.Title)
	fmt.Println("User ID: ", planner.UserId)
	fmt.Println("Enter New planner title: ")

	title, err := promptUser(reader, "Enter planner title: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	userid, err := promptUser(reader, "Enter planner userid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	update := handle.UpdatePlannerRequest{
		Id:     id,
		Title:  title,
		UserId: userid,
	}
	fmt.Println("Updating planner...")
	err = p.UpdatePlanner(&update)
	if err != nil {
		fmt.Printf("Error updating planner with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Updated planner with id: ", planner.Id)
}

// deletePlanner deletes a planner based on user input.
// It prompts the user to enter a planner ID to be deleted.
// It then retrieves the planner from the PlannerControl and displays its details to the user.
// The user is then asked to confirm the deletion.
// If the user confirms with 'y', the planner is deleted using the DeletePlanner method of the PlannerControl.
// If the deletion is successful, a success message is printed.
// If the user enters 'n', the deletion is canceled.
// If the user enters any other value, an invalid input message is printed.
// If there is an error in any step, appropriate error messages are printed.
func deletePlanner(p *handle.PlannerControl) {
	reader := bufio.NewReader(os.Stdin)
	id, err := promptUser(reader, "Enter planner id to be deleted: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	// Get planner first to show user what planner is being deleted
	req := handle.GetPlannerRequest{
		Id: id,
	}
	planner, err := p.GetPlanner(&req)
	if err != nil {
		fmt.Printf("Error getting planner with id %s: %s\n", id, err)
		return
	}
	fmt.Println("Got planner: ", planner.Title)
	fmt.Println("are you sure you want to delete this planner? (y/n)")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read from stdin: %s\n", err)
	}
	confirm = strings.TrimSpace(confirm)
	if confirm == "n" {
		fmt.Println("Planner not deleted")
	} else if confirm == "y" {
		fmt.Println("Deleting planner...")
		req := handle.DeletePlannerRequest{
			Id: id,
		}
		err = p.DeletePlanner(&req)
		if err != nil {
			fmt.Printf("Error deleting planner with id %s: %s\n", id, err)
			return
		}
		fmt.Println("Deleted planner with id: ", planner.Id)
	} else {
		fmt.Println("Invalid input")
	}
}

func listPlanners(p *handle.PlannerControl) {
	fmt.Println("Listing planners...")
	planners, err := p.ListPlanners()
	if err != nil {
		fmt.Println("Error listing planners: ", err)
		return
	}
	// get planner id and title of each planner
	for _, planner := range planners.Planners {
		fmt.Printf("Planner id: %s, Title: %s, User ID %s\n", planner.Id, planner.Title, planner.UserId)
	}
}

func runCommand(commandStr string) error {
	taskRouter, goalRouter, planRouter, plannerRouter, db := cliSetup()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}()

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	commandName := arrCommandStr[0]

	if commandName == "exit" {
		os.Exit(0)
	}

	if command, ok := taskCommands[commandName]; ok {
		command(taskRouter)
	} else if command, ok := goalCommands[commandName]; ok {
		command(goalRouter)
	} else if command, ok := planCommands[commandName]; ok {
		command(planRouter)
	} else if command, ok := plannerCommands[commandName]; ok {
		command(plannerRouter)
	} else {
		fmt.Println("Command not found")
	}
	return nil
}
