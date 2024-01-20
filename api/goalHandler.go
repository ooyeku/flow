package api

import (
	"encoding/json"
	"flow/pkg/handle"
	"github.com/gorilla/mux"
	"net/http"
)

type GoalHandler struct {
	Control *handle.GoalControl
}

func (h *GoalHandler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var req handle.CreateGoalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreateGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *GoalHandler) GetGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetGoalRequest{Id: id}
	res, err := h.Control.GetGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *GoalHandler) GetGoalByObjective(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objective := vars["objective"]
	req := handle.GetGoalByObjectiveRequest{Objective: objective}
	res, err := h.Control.GetGoalByObjective(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *GoalHandler) GetGoalsByPlannerIdRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plannerId := vars["planner_id"]
	req := handle.GetGoalsByPlannerIdRequest{PlannerId: plannerId}
	res, err := h.Control.GetGoalsByPlannerId(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *GoalHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req handle.UpdateGoalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)

	req.Id = id // Ensure the ID from the URL is used

	err = h.Control.UpdateGoal(&req)
	handleError(w, err, http.StatusInternalServerError)

	w.WriteHeader(http.StatusOK)
}

func (h *GoalHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeleteGoalRequest{Id: id}
	err := h.Control.DeleteGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

func (h *GoalHandler) ListGoals(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListGoals()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
