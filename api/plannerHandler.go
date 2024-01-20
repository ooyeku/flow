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
	w.Header().Set("Content-Type", "application/json")
	var req handle.CreatePlannerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreatePlanner(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) GetPlanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	req := &handle.GetPlannerRequest{
		Id: id,
	}
	res, err := h.Control.GetPlanner(req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) GetPlannerByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	title := vars["title"]
	req := &handle.GetPlannerByTitleRequest{
		Title: title,
	}
	res, err := h.Control.GetPlannerByTitle(req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) GetPlannerByOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	owner := vars["owner"]
	req := &handle.GetPlannerByOwnerRequest{
		UserId: owner,
	}
	res, err := h.Control.GetPlannerByOwner(req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *PlannerHandler) UpdatePlanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	req := &handle.UpdatePlannerRequest{
		Id: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	err = h.Control.UpdatePlanner(req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) DeletePlanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.Control.DeletePlanner(&handle.DeletePlannerRequest{
		Id: id,
	})
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) ListPlanners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := h.Control.ListPlanners()
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
