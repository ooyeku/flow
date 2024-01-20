package api

import (
	"encoding/json"
	"flow/pkg/handle"
	"github.com/gorilla/mux"
	"net/http"
)

type PlannerHandler struct {
	Control *handle.PlannerControl
}

func (h *PlannerHandler) CreatePlanner(w http.ResponseWriter, r *http.Request) {
	var req handle.CreatePlannerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreatePlanner(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) GetPlanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetPlannerRequest{Id: id}
	res, err := h.Control.GetPlanner(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) UpdatePlanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req handle.UpdatePlannerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)

	req.Id = id // Ensure the ID from the URL is used

	err = h.Control.UpdatePlanner(&req)
	handleError(w, err, http.StatusInternalServerError)

	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) DeletePlanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeletePlannerRequest{Id: id}
	err := h.Control.DeletePlanner(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) ListPlanners(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListPlanners()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
