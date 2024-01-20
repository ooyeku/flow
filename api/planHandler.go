package api

import (
	"encoding/json"
	"flow/pkg/handle"
	"github.com/gorilla/mux"
	"net/http"
)

type PlanHandler struct {
	Control *handle.PlanControl
}

func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var req handle.CreatePlanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreatePlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlanHandler) GetPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetPlanRequest{Id: id}
	res, err := h.Control.GetPlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlanHandler) GetPlanByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planName := vars["plan_name"]
	req := handle.GetPlanByNameRequest{PlanName: planName}
	res, err := h.Control.GetPlanByName(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlanHandler) GetPlansByGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId := vars["goal_id"]
	req := handle.GetPlansByGoalRequest{GoalId: goalId}
	res, err := h.Control.GetPlansByGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlanHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req handle.UpdatePlanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)

	req.Id = id // Ensure the ID from the URL is used

	err = h.Control.UpdatePlan(&req)
	handleError(w, err, http.StatusInternalServerError)

	w.WriteHeader(http.StatusOK)
}

func (h *PlanHandler) DeletePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeletePlanRequest{Id: id}
	err := h.Control.DeletePlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

func (h *PlanHandler) ListPlans(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListPlans()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
