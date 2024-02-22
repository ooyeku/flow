package services

import (
	"github.com/ooyeku/flow/pkg/models"
)

type CrossService struct {
	goalService    *GoalService
	planService    *PlanService
	taskService    *TaskService
	plannerService *PlannerService
}

func NewCrossService(gs *GoalService, ps *PlanService, ts *TaskService, pls *PlannerService) *CrossService {
	return &CrossService{
		goalService:    gs,
		planService:    ps,
		taskService:    ts,
		plannerService: pls,
	}
}

func (cs *CrossService) CreateVersionImage(v *models.Version) error {
	// Fetch the goal
	g, err := cs.goalService.GetGoal(string(v.GoalID))
	if err != nil {
		return err
	}

	// Fetch the plans
	plans, err := cs.planService.ListPlans()
	if err != nil {
		return err
	}

	// Fetch the tasks
	tasks, err := cs.taskService.ListTasks()
	if err != nil {
		return err
	}

	// Create the snapshot
	snapshot := &models.Snapshot{
		Goal:  g,
		Plans: plans,
		Tasks: tasks,
	}

	v.Image = *snapshot

	return nil
}
