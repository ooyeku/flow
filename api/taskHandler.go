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
