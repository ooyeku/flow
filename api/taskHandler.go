package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goworkflow/pkg/handle"
	"log"
	"net/http"
)

type TaskHandler struct {
	Control *handle.TaskControl
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		log.Printf("Error due to: %s", err)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req handle.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreateTask(req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetTaskRequest{ID: id}
	res, err := h.Control.GetTask(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListTasks()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req handle.UpdateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)

	req.ID = id // Ensure the ID from the URL is used

	err = h.Control.UpdateTask(&req)
	handleError(w, err, http.StatusInternalServerError)

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeleteTaskRequest{ID: id}
	err := h.Control.DeleteTask(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}
