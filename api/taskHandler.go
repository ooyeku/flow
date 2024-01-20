package api

import (
	"encoding/json"
	"flow/pkg/handle"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// TaskHandler handles HTTP requests related to tasks. It contains a control object of type *handle.TaskControl,
// which is used to interact with the task service. It has methods to handle the following HTTP requests:
// - CreateTask: handles the creation of a new task.
// - UpdateTask: handles the updating of an existing task.
// - DeleteTask: handles the deletion of a task.
// - GetTask: handles the retrieval of a task by ID.
// - GetTaskByTitle: handles the retrieval of a task by title.
// - GetTaskByOwner: handles the retrieval of tasks by owner.
// - ListTasks: handles the listing of all tasks.
type TaskHandler struct {
	Control *handle.TaskControl
}

// handleError checks if there is an error and if so, it writes the error message to the response writer
// with the specified status code and logs the error message.
func handleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		log.Printf("Error due to: %s", err)
	}
}

// CreateTask creates a new task based on the request data.
// It decodes the JSON request body to a CreateTaskRequest struct.
// Then, it generates a unique ID for the task, creates a task instance using the provided data, and calls the CreateTask method of the TaskControl.
// If the task creation is successful, it returns a CreateTaskResponse with the ID of the created task.
// If any error occurs during the process, it handles the error by writing an HTTP error response with the corresponding status code and logging the error message.
// Finally, it encodes the response data into JSON and writes it to the HTTP response.
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req handle.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreateTask(req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetTask retrieves a task by its ID.
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetTaskRequest{ID: id}
	res, err := h.Control.GetTask(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetTaskByTitle is a method of TaskHandler that handles the GET request to retrieve a task by its title.
// It expects the request URL to contain a "title" parameter.
// It creates a GetTaskByTitleRequest with the title parameter and passes it to the TaskControl's GetTaskByTitle method.
// If there is an error during the process, it returns a JSON error response with the corresponding status code.
// It encodes the response into JSON format and writes it to the ResponseWriter.
// If there is an error encoding the response, it returns a JSON error response with the corresponding status code.
func (h *TaskHandler) GetTaskByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	req := handle.GetTaskByTitleRequest{Title: title}
	res, err := h.Control.GetTaskByTitle(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetTaskByOwner is a method of TaskHandler that handles the GET request to retrieve tasks by owner.
// It expects the request to include an "owner" path variable.
// It retrieves tasks based on the owner from the request and returns them as a JSON response.
// If there is an error during the process, it returns a JSON error response with the corresponding status code.
func (h *TaskHandler) GetTaskByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	req := handle.GetTaskByOwnerRequest{Owner: owner}
	res, err := h.Control.GetTaskByOwner(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// ListTasks retrieves a list of tasks.
// It calls the ListTasks method of the TaskControl to retrieve the tasks.
// If there is any error during the retrieval, it returns an Internal Server Error status.
// Finally, it encodes the retrieved tasks into JSON and writes it to the response writer.
// If there is any error during encoding, it returns an Internal Server Error status.
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListTasks()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// UpdateTask is a method of TaskHandler that handles the PUT request to update a specific task.
// It expects the request to include an ID path variable and a request body containing the updated task information.
// The ID from the URL is used to ensure the correct task is updated.
// If there is an error during the process, it returns a JSON error response with the corresponding status code.
// If the update is successful, it returns a 200 OK status code.
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

// DeleteTask deletes a task with the given ID.
// It extracts the task ID from the URL parameters and creates a DeleteTaskRequest with that ID.
// Then, it calls the DeleteTask method of TaskControl to delete the task.
// If there is any error during task deletion, it returns an Internal Server Error status.
// Finally, it sets the response writer's status code to OK.
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeleteTaskRequest{ID: id}
	err := h.Control.DeleteTask(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}
