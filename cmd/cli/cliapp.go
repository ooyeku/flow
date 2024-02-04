package main

import (
	"bufio"
	"fmt"
	"github.com/asdine/storm"
	"github.com/ooyeku/flow/internal/conf"
	"github.com/ooyeku/flow/internal/inmemory"
	"github.com/ooyeku/flow/pkg/handle"
	"github.com/ooyeku/flow/pkg/services"
	"log"
	"os"
	"strings"
)

// main is the entry point function for the CLI application.
// It displays a welcome message and continuously prompts the user for a command.
// The entered command string is passed to the runCommand function for execution.
// Any error that occurs during the process is printed to stderr.
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

// cliSetup is a function that initializes the CLI setup by creating instances of different controls and the database.
// It opens the database using the path obtained from the configuration.
// It then creates in-memory stores, services, and routers for tasks, goals, plans, and planners using the database.
// Finally, it returns the taskRouter, goalRouter, planRouter, plannerRouter, and the database instance.
// The returned taskRouter is of type *handle.TaskControl and has the following methods:
// - CreateTask: Creates a new task with the provided request and returns the response.
// - UpdateTask: Updates an existing task using the provided request.
// - DeleteTask: Deletes a task with the specified ID.
// - GetTask: Retrieves a task with the specified ID and returns the response.
// - GetTaskByTitle: Retrieves a task with the specified title and returns the response.
// - GetTaskByOwner: Retrieves tasks with the specified owner and returns a slice of responses.
// - ListTasks: Retrieves all tasks and returns a slice of responses.
// The returned goalRouter is of type *handle.GoalControl and has the following methods:
// - CreateGoal: Creates a new goal with the provided request and returns the response.
// - UpdateGoal: Updates an existing goal using the provided request.
// - DeleteGoal: Deletes a goal with the specified ID.
// - GetGoal: Retrieves a goal with the specified
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

// taskCommands is a map that contains various commands related to task operations.
// The key represents the command name and the value represents the corresponding function to be executed.
var taskCommands = map[string]func(*handle.TaskControl){
	"create-task":       createTask,
	"ct":                createTask,
	"get-task":          getTask,
	"gt":                getTask,
	"get-task-by-title": getTaskByTitle,
	"gtk":               getTaskByTitle,
	"get-task-by-owner": getTaskByOwner,
	"gto":               getTaskByOwner,
	"update-task":       updateTasks,
	"ut":                updateTasks,
	"delete-task":       deleteTask,
	"dt":                deleteTask,
	"list-tasks":        listTasks,
	"lt":                listTasks,
}

// goalCommands is a map that contains various commands related to goal operations. The key represents the command name and the value represents the corresponding function to be executed.
var goalCommands = map[string]func(*handle.GoalControl){
	"create-goal":         createGoal,
	"cg":                  createGoal,
	"get-goal":            getGoal,
	"gg":                  getGoal,
	"get-goal-by-title":   getGoalByObjective,
	"ggt":                 getGoalByObjective,
	"get-goal-by-planner": getGoalsByPlannerId,
	"ggp":                 getGoalsByPlannerId,
	"update-goal":         updateGoals,
	"ug":                  updateGoals,
	"delete-goal":         deleteGoal,
	"dg":                  deleteGoal,
	"list-goals":          listGoals,
	"lg":                  listGoals,
}

// planCommands is a map that contains various commands related to plan operations. The key represents the command name and the value represents the corresponding function to be executed
var planCommands = map[string]func(*handle.PlanControl){
	"create-plan":      createPlan,
	"cp":               createPlan,
	"get-plan":         getPlan,
	"gp":               getPlan,
	"get-plan-by-name": getPlanByName,
	"gpn":              getPlanByName,
	"get-plan-by-goal": getPlanByGoal,
	"gpg":              getPlanByGoal,
	"update-plan":      updatePlans,
	"up":               updatePlans,
	"delete-plan":      deletePlan,
	"dp":               deletePlan,
	"list-plans":       listPlans,
	"lp":               listPlans,
}

// plannerCommands is a map that contains various commands related to planner operations. The key represents the command name and the value represents the corresponding function to be executed
var plannerCommands = map[string]func(*handle.PlannerControl){
	"create-planner":       createPlanner,
	"cpl":                  createPlanner,
	"get-planner":          getPlanner,
	"gpl":                  getPlanner,
	"get-planner-by-title": getPlannerByGoal,
	"gpt":                  getPlannerByGoal,
	"get-planner-by-owner": getPlannerByOwner,
	"gpo":                  getPlannerByOwner,
	"update-planner":       updatePlanners,
	"upl":                  updatePlanners,
	"delete-planner":       deletePlanner,
	"dpl":                  deletePlanner,
	"list-planners":        listPlanners,
	"lpl":                  listPlanners,
}

// createTask is a function that prompts the user to enter information for a new task,
// creates a CreateTaskRequest using the entered information, and calls the CreateTask method
// on the provided *handle.TaskControl instance to create the task.
// If there is an error creating the task, it prints the error message to the console.
// Finally, it prints the ID of the created task to the console.
// The function makes use of the promptUser function to read user input.
// The promptUser function takes a *bufio.Reader instance and a prompt string as arguments,
// and returns the trimmed user input or an error.
// The promptUser function is defined as follows:
//
//	func promptUser(reader *bufio.Reader, prompt string) (string, error) {
//		fmt.Println(prompt)
//		response, err := reader.ReadString('\n')
//		if err != nil {
//			return "", fmt.Errorf("could not read from stdin: %s", err)
//		}
//		return strings.TrimSpace(response), nil
//	}
//
// The createTask function relies on the following types and methods defined in other parts of the codebase:
// - handle.TaskControl: A user-defined type representing task control. It has a CreateTask method that takes a *handle.CreateTaskRequest and returns a response and an error.
// - handle.CreateTaskRequest: A user-defined type representing a request to create a task. It has fields for title, description, and owner.
// - services.TaskService: A user-defined type representing a task service. It exposes methods for creating, updating, deleting, and retrieving tasks.
// - models.Task: A user-defined type representing a task with fields for ID, title, description, owner, started, completed, createdAt, and updatedAt.
// - store.TaskStore: A user-defined type representing a task store. It has methods for creating, updating, deleting, and retrieving tasks.
// Here is an example of how to use the createTask function:
// t := &handle.TaskControl{service: &services.TaskService{Store: &store.TaskStore{}}} // Instantiate TaskControl
// createTask(t) // Call createTask function with the TaskControl instance as argument
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
		return
	}
	fmt.Println("Created task with id: ", res.ID)
}

// getTask retrieves a task by its ID.
// It prompts the user to enter the task ID and retrieves the task from TaskControl service.
// If there is an error while reading from standard input or retrieving the task, it prints an error message.
// It prints the task's title and description if the task is retrieved successfully.
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

// getTaskByTitle retrieves a task by its title from a TaskControl instance.
// It prompts the user to enter the task title, and then calls the GetTaskByTitle function of the TaskControl instance.
// If the task is found, it prints the task title and description.
// If there is an error retrieving the task or prompting the user, it logs the error message.
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
		fmt.Printf("Error getting task with title %s: %s\n", title, err)
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
}

// getTaskByOwner is a function that retrieves all tasks owned by a given owner.
// It prompts the user to enter the task owner's name, then sends a request to the TaskControl service
// to get the tasks owned by that owner. The function then prints the details of each task.
// If there is an error reading the user's input or retrieving the tasks, an error message is logged.
// Parameters:
// - t: a pointer to a handle.TaskControl struct representing the task control service
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

// updateTasks is a function that allows the user to update a task in the task control system.
// It takes a pointer to a TaskControl object as a parameter, which provides access to the task control system.
// The function prompts the user to enter the ID of the task to be updated.
// It then retrieves the task from the task control system using the provided ID.
// If an error occurs during the retrieval, the function prints an error message and returns.
// The function displays the retrieved task's title and description to the user.
// The user is prompted to enter a new task title, description, and owner.
// The function then updates the task with the new values and prints a message indicating that the task is being updated.
// If an error occurs during the update, the function prints an error message and returns.
// Finally, the function prints a message indicating that the task has been updated with its ID.
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

// deleteTask deletes a task based on user input.
// It prompts the user to enter the task ID to be deleted.
// It then retrieves the task using the GetTask method of the TaskControl service to show the user which task is being deleted.
// The user is then asked to confirm the deletion by entering "y" or "n".
// If the user confirms the deletion, the task is deleted using the DeleteTask method of the TaskControl service.
// If the deletion is successful, a message is printed to confirm the deletion.
// If the user cancels the deletion or provides invalid input, appropriate messages are printed.
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

// listTasks is a function that lists all tasks using a TaskControl object.
// It calls the ListTasks function of the TaskControl object to get a list of tasks,
// and then prints the ID, title, and description of each task.
//
// Parameters:
// - t: A pointer to a TaskControl object.
// Returns: None.
// Example usage:
//
//	t := &handle.TaskControl{}
//	listTasks(t)
func listTasks(t *handle.TaskControl) {
	fmt.Println("Listing tasks...")
	tasks, err := t.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks: ", err)
		return
	}
	// get task id and title of each task
	for _, task := range tasks {
		fmt.Printf("|Task id: %s, Title: %s, Description %s|\n", task.ID, task.Title, task.Description)
	}
}

// Prompt user to enter goal objective
func createGoal(g *handle.GoalControl) {
	fmt.Println("Creating goal...")
	reader := bufio.NewReader(os.Stdin)
	objective, err := promptUser(reader, "Enter goal objective: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	deadline, err := promptUser(reader, "Enter goal deadline date in YYYY-MM-DD format: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	deadlineTime, err := promptUser(reader, "Enter goal deadline time in HH:MM format: ")
	plannerid, err := promptUser(reader, "Enter goal plannerid: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}

	// Convert deadline to time.Time
	deadline = deadline + "T" + deadlineTime + ":00"

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

// getGoal is a function that retrieves a goal from GoalControl by prompting the user for the goal ID and
// calling the GetGoal method on GoalControl. It prints the goal's objective, deadline, and planner ID if
// the goal is successfully retrieved. If there is an error, it displays the corresponding error message.
//
// Parameters:
// - g: A pointer to an instance of the GoalControl struct which provides access to goal-related operations.
//
// Example Usage:
//
//	g := &handle.GoalControl{Service: &services.GoalService{}}
//	getGoal(g)
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

// getGoalByObjective retrieves a goal by its objective from the GoalControl service.
// It prompts the user to enter the goal objective, creates a GetGoalByObjectiveRequest,
// and calls the GetGoalByObjective method of the GoalControl service to fetch the goal.
// If there is an error fetching the goal, the error message is logged in the GetGoalByObjective method.
// If the goal is fetched successfully, it prints the goal's objective, deadline, and plannerID to the console.
func getGoalByObjective(g *handle.GoalControl) {
	fmt.Println("Getting goal...")
	reader := bufio.NewReader(os.Stdin)
	objective, err := promptUser(reader, "Enter goal title(objective): ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	req := handle.GetGoalByObjectiveRequest{
		Objective: objective,
	}
	goal, err := g.GetGoalByObjective(&req)
	if err != nil {
		// error message is logged in GetGoalByObjective
		fmt.Println("Error getting goal: ", err)
		return
	}
	fmt.Println("Got goal: ", goal.Goal.Objective)
	fmt.Println("Deadline: ", goal.Goal.Deadline)
	fmt.Println("PlannerID: ", goal.Goal.PlannerId)
}

// getGoalsByPlannerId is a function that retrieves goals by a given planner ID.
// It prompts the user to enter the goal planner ID and sends a request to the GoalControl service to get the goals.
// If there is an error retrieving the goals, the error message is logged in the GetGoalByObjective function.
// The retrieved goals are then printed to the console.
//
// Example usage:
//
//	goalControl := &handle.GoalControl{
//	    Service: &services.GoalService{},
//	}
//	getGoalsByPlannerId(goalControl)
//
// Parameters:
// - g: a handle.GoalControl instance that provides access to the GoalControl service methods.
//
// Dependencies:
// - bufio.NewReader: used for reading user input from standard input.
// - promptUser: a function that prompts the user to enter a value and returns the entered value as a string.
// - handle.GetGoalsByPlannerIdRequest: a struct representing the request to retrieve goals by planner ID.
//
// Returns: None
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
		fmt.Println("Error getting goals: ", err)
		return
	}
	for _, goal := range goals.Goals {
		fmt.Println("Got goal: ", goal.Objective)
		fmt.Println("Deadline: ", goal.Deadline)
		fmt.Println("PlannerID: ", goal.PlannerId)
	}
}

// updateGoals updates a goal by prompting the user for new goal objective, deadline, and planner ID.
// It first prompts the user to enter the goal ID of the goal to be updated.
// Then it retrieves the goal from the GoalControl service using the given ID.
// If there is an error retrieving the goal, it prints an error message and returns.
// Otherwise, it prompts the user to enter the new goal objective, deadline, and planner ID.
// After collecting the new values, it creates an UpdateGoalRequest with the collected values.
// It then calls the UpdateGoal method of the GoalControl service to update the goal.
// If there is an error updating the goal, it prints an error message and returns.
// Otherwise, it prints a success message with the updated goal ID.
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

// deleteGoal deletes a goal based on user input.
// It prompts the user to enter the ID of the goal to be deleted and then confirms the deletion with the user.
// It first gets the goal based on the provided ID to show the user what goal is being deleted.
// If the goal is found, it asks the user for confirmation.
// If the user confirms the deletion, it calls the DeleteGoal method on the GoalControl service to delete the goal.
// If there is an error during any step, it prints an error message accordingly.
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

// listGoals is a function that lists all the goals.
//
// It calls the ListGoals method of the GoalControl service to retrieve the list of goals.
// If there is an error retrieving the goals, it prints an error message and returns.
// It then iterates over the list of goals and prints the goal ID, objective, deadline, and planner ID for each goal.
// The goal information is printed to standard output.
//
// Example usage:
//
//	gCtrl := &handle.GoalControl{Service: &services.GoalService{}}
//	listGoals(gCtrl)
//
// Output:
//
//	Listing goals...
//	Goal id: 1, Objective: Finish project, Deadline: 2022-12-31T23:59:00, PlannerID: 12345
//	Goal id: 2, Objective: Exercise daily, Deadline: 2023-01-01T08:00:00, PlannerID: 67890
func listGoals(g *handle.GoalControl) {
	fmt.Println("Listing goals...")
	goals, err := g.ListGoals()
	if err != nil {
		fmt.Println("Error listing goals: ", err)
		return
	}
	// get goal id and objective of each goal
	for _, goal := range goals.Goals {
		fmt.Printf("|Goal id: %s, Objective: %s, Deadline %s, PlannerID: %s|\n", goal.Id, goal.Objective, goal.Deadline, goal.PlannerId)
	}
}

// createPlan is a function that creates a plan by prompting the user to enter various details such as plan name, goal ID, description, date, and time.
// It takes a *handle.PlanControl parameter p, which is an instance of the PlanControl struct.
// It uses the promptUser function to prompt the user for input and validate it.
// It then creates a handle.CreatePlanRequest object with the provided input, and calls the CreatePlan method of p.Service with the request.
// If there is an error during plan creation, it prints the error message.
// Finally, it prints the ID of the created plan.
// Note that this function runs in a Goroutine, allowing for concurrent plan creation.
// Parameters:
// - p: A pointer to an instance of the handle.PlanControl struct which contains a PlanService instance for plan CRUD operations.
func createPlan(p *handle.PlanControl) {
	fmt.Println("Creating plan...")
	reader := bufio.NewReader(os.Stdin)
	name, err := promptUser(reader, "Enter plan name: ")
	if err != nil {
		log.Fatalf("Could not read from stdin: %s", err)
	}
	goalid, err := promptUser(reader, "Enter plan goalid: ")
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
		GoalId:          goalid,
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

// getPlan function is used to retrieve a specific plan by its ID.
// It prompts the user to enter the plan ID and passes it to the PlanControl.GetPlan method.
// If the plan is found, it prints the plan details to the console.
// If there is any error in retrieving the plan, it prints an error message.
// The function does not return anything.
//
// Example usage:
//
//	p := &handle.PlanControl{
//	  Service: <PlanServiceInstance>,
//	}
//	getPlan(p)
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
	fmt.Println("GoalID: ", plan.Plan.GoalId)
}

// getPlanByName is a function that retrieves a plan by its name from PlanControl.
// It prompts the user to enter the plan name and uses the PlanControl service to get the plan.
// If there is an error retrieving the plan, the error message is logged in GetPlanByName.
// The function prints the plan name and description if the plan is found.
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
		fmt.Printf("Error getting plan with name %s: %s\n", name, err)
		return
	}
	fmt.Println("Got plan: ", plan.Plan.PlanName)
	fmt.Println("Description: ", plan.Plan.PlanDescription)
}

// getPlanByGoal retrieves plans based on a goal ID.
// It prompts the user to enter a plan goal ID and makes a request to the PlanControl to get the plans.
// It then prints the plan names and descriptions.
//
// Parameters:
// - p: A pointer to the PlanControl struct.
//
// Example usage:
//
//	getPlanByGoal(p)
//
// Returns: None
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
		fmt.Printf("Error getting plans with goal id %s: %s\n", goalid, err)
		return
	}
	for _, plan := range plans.Plans {
		fmt.Println("Got plan: ", plan.PlanName)
		fmt.Println("Description: ", plan.PlanDescription)
	}
}

// updatePlans is a function that allows the user to update a specific plan.
// It prompts the user to enter the plan ID of the plan to be updated.
// It then retrieves the plan information using the plan ID.
// It prompts the user to enter a new plan name, description, date, and time.
// After receiving the new values, it creates an UpdatePlanRequest object and sends it to the PlanControl's UpdatePlan method.
// If an error occurs during the process, it prints an error message.
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

// deletePlan is a function that allows the user to delete a plan based on its ID.
// It prompts the user to enter the plan ID to be deleted.
// It then retrieves the plan details using the GetPlan method from the PlanControl service.
// If the plan is found, it displays the plan name and asks for confirmation to delete.
// If the user confirms with 'y', it calls the DeletePlan method from the PlanControl service to delete the plan.
// If the plan is successfully deleted, it displays a success message.
// If the user enters 'n' or any other input, it cancels the deletion process.
// Parameters:
// - p: a pointer to the PlanControl structure that provides access to plan-related operations.
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
		fmt.Printf("|Plan id: %s, Name: %s, Description %s|\n", plan.Id, plan.PlanName, plan.PlanDescription)
	}
}

// createPlanner is a function that creates a planner by prompting the user to enter a title and userid.
// It uses the PlannerControl struct to make a request to create the planner using the PlannerService.
// If the creation is successful, it prints the planner's ID.
// If there is an error during the creation process, it prints the error message.
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

// getPlanner retrieves a planner from the PlannerControl service by its id.
// It prompts the user to enter the planner id, and then uses the GetPlanner method of the PlannerControl service to retrieve the planner.
// If there is an error reading from stdin, it will log a fatal error and exit.
// If there is an error retrieving the planner, it will print an error message and return.
// Otherwise, it will print the planner's title and owner's user ID.
// Function signature: func getPlanner(p *handle.PlannerControl)
// Parameter:
// - p: a pointer to a PlannerControl instance
// Example usage:
//
//	reader := bufio.NewReader(os.Stdin)
//	plannerControl := handle.PlannerControl{Service: &services.PlannerService{}}
//	getPlanner(&plannerControl)
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
		fmt.Println("Error getting planner: ", err)
		return
	}
	fmt.Println("Got planner: ", planner.Title)
	fmt.Println("User ID: ", planner.UserId)
}

// getPlannerByOwner retrieves a list of planners owned by a specific user.
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
		fmt.Println("Error getting planner: ", err)
		return
	}
	for _, planner := range planners {
		fmt.Println("Got planner: ", planner.Title)
		fmt.Println("User ID: ", planner.UserId)
	}
}

// updatePlanners is a function that updates a planner in the PlannerControl object.
// It prompts the user to enter the planner ID of the planner to be updated.
// It then retrieves the planner from the PlannerControl object using the GetPlanner method.
// If the planner is found, it prompts the user to enter the new title and user ID for the planner.
// It then creates an UpdatePlannerRequest object with the updated details.
// Finally, it calls the UpdatePlanner method of the PlannerControl object to update the planner.
// If there is an error retrieving or updating the planner, it prints an error message.
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

// deletePlanner is a function that prompts the user to enter a planner ID and then deletes the planner with that ID.
// It first gets the planner information to show the user what planner is being deleted.
// After confirming with the user, it proceeds to delete the planner and prints a success message.
// If there is any error during the process, it prints an error message.
// Parameters:
// - p: a pointer to a PlannerControl struct that contains the planner service.
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

// listPlanners is a function that retrieves a list of planners from the PlannerControl service and prints the details of each planner.
// It takes a pointer to a PlannerControl instance as a parameter.
// First, it prints a message indicating that the planners are being listed.
// Then, it calls the ListPlanners method of the PlannerControl service to get the list of planners.
// If there is an error during the retrieval, it prints an error message and returns.
// Otherwise, it iterates over the list of planners and prints the ID, Title, and User ID of each planner.
func listPlanners(p *handle.PlannerControl) {
	fmt.Println("Listing planners...")
	planners, err := p.ListPlanners()
	if err != nil {
		fmt.Println("Error listing planners: ", err)
		return
	}
	// get planner id and title of each planner
	for _, planner := range planners.Planners {
		fmt.Printf("|Planner id: %s, Title: %s, User ID %s|\n", planner.Id, planner.Title, planner.UserId)
	}
}

// runCommand takes a command string as input and executes the corresponding command.
// It initializes the necessary routers and database connection using cliSetup.
// If the command string is empty, it prints a message and returns.
// It trims any newline characters from the command string.
// It splits the command string into individual words.
// If the first word is "exit", it exits the application.
// If the first word matches a command in the taskCommands map,
// it executes the corresponding command using the taskRouter.
// If the first word matches a command in the goalCommands map,
// it executes the corresponding command using the goalRouter.
// If the first word matches a command in the planCommands map,
// it executes the corresponding command using the planRouter.
// If the first word matches a command in the plannerCommands map,
// it executes the corresponding command using the plannerRouter.
// If no command matches the first word, it prints a message.
// It returns nil to indicate success.
func runCommand(commandStr string) error {
	taskRouter, goalRouter, planRouter, plannerRouter, db := cliSetup()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}()

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) == 0 {
		fmt.Println("No command entered. Please try again.")
		return nil
	}

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
