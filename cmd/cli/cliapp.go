package main

import (
	"bufio"
	"fmt"
	"github.com/asdine/storm"
	"goworkflow/internal/conf"
	"goworkflow/internal/inmemory"
	"goworkflow/pkg/handle"
	"goworkflow/pkg/services"
	"log"
	"os"
	"strings"
)

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
func cliSetup() (*handle.TaskControl, *storm.DB) {
	dbPath := conf.GetDBPath()
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("error opening db: %s", err)
	}

	// Intialize router, service and inmemory store
	taskStore := inmemory.NewInMemoryTaskStore(db)
	taskService := services.NewTaskService(taskStore)
	taskRouter := handle.NewTaskControl(taskService)

	return taskRouter, db
}

func promptUser(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read from stdin: %s", err)
	}
	return strings.TrimSpace(response), nil
}

var taskCommands = map[string]func(*handle.TaskControl){
	"create-task":       createTask,
	"get-task":          getTask,
	"get-task-by-title": getTaskByTitle,
	"get-task-by-owner": getTaskByOwner,
	"update-task":       updateTasks,
	"delete-task":       deleteTask,
	"list-tasks":        listTasks,
}

var goalCommands = map[string]func(*handle.GoalControl){
	"create-goal":       createGoal,
	"get-goal":          getGoal,
	"get-goal-by-title": getGoalByTitle,
	"get-goal-by-owner": getGoalByOwner,
	"update-goal":       updateGoals,
	"delete-goal":       deleteGoal,
	"list-goals":        listGoals,
}

var planCommands = map[string]func(*handle.PlanControl){
	"create-plan":       createPlan,
	"get-plan":          getPlan,
	"get-plan-by-title": getPlanByTitle,
	"get-plan-by-owner": getPlanByOwner,
	"update-plan":       updatePlans,
	"delete-plan":       deletePlan,
	"list-plans":        listPlans,
}

var plannerCommands = map[string]func(*handle.PlannerControl){
	"create-planner":       createPlanner,
	"get-planner":          getPlanner,
	"get-planner-by-title": getPlannerByTitle,
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

// getTask retrieves a task from the task router based on the provided task ID.
// It prompts the user to input the task ID, sends a GetTaskRequest to the task router,
// and prints the task details if it exists.
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
	}
	req := handle.GetTaskByOwnerRequest{
		Owner: owner,
	}
	task, err := t.GetTaskByOwner(&req)
	if err != nil {
		return
	}
	fmt.Println("Got task: ", task.Title)
	fmt.Println("Description: ", task.Description)
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

func createGoal(g *handle.GoalControl)     {}
func getGoal(g *handle.GoalControl)        {}
func getGoalByTitle(g *handle.GoalControl) {}
func getGoalByOwner(g *handle.GoalControl) {}
func updateGoals(g *handle.GoalControl)    {}
func deleteGoal(g *handle.GoalControl)     {}
func listGoals(g *handle.GoalControl)      {}

func createPlan(p *handle.PlanControl)     {}
func getPlan(p *handle.PlanControl)        {}
func getPlanByTitle(p *handle.PlanControl) {}
func getPlanByOwner(p *handle.PlanControl) {}
func updatePlans(p *handle.PlanControl)    {}
func deletePlan(p *handle.PlanControl)     {}
func listPlans(p *handle.PlanControl)      {}

func createPlanner(p *handle.PlannerControl)     {}
func getPlanner(p *handle.PlannerControl)        {}
func getPlannerByTitle(p *handle.PlannerControl) {}
func getPlannerByOwner(p *handle.PlannerControl) {}
func updatePlanners(p *handle.PlannerControl)    {}
func deletePlanner(p *handle.PlannerControl)     {}
func listPlanners(p *handle.PlannerControl)      {}

func runCommand(commandStr string) error {
	taskRouter, db := cliSetup()
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

	// Check if the handler exists for the given command name
	if handler, exists := taskCommands[commandName]; exists {
		handler(taskRouter)
	} else {
		return fmt.Errorf("unknown command: %s", commandStr)
	}
	return nil
}