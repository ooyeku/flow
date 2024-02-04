package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ooyeku/flow/pkg/handle"
	"net/http"
)

// GoalHandler represents a handler for managing goals.
// It has a Control field of type *handle.GoalControl that provides functions for handling various goal operations.
// Example usage for creating a goal:
//
//	func (h *GoalHandler) CreateGoal(w http.ResponseWriter, r *http.Request) {
//			var req handle.CreateGoalRequest
//			err := json.NewDecoder(r.Body).Decode(&req)
//			// Handle error
//			res, err := h.Control.CreateGoal(&req)
//			// Handle error
//			err = json.NewEncoder(w).Encode(res)
//			// Handle error
//	}
//
// Example usage for getting a goal by ID:
//
//	func (h *GoalHandler) GetGoal(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			id := vars["id"]
//			req := handle.GetGoalRequest{Id: id}
//			res, err := h.Control.GetGoal(&req)
//			// Handle error
//			err = json.NewEncoder(w).Encode(res)
//			// Handle error
//	}
//
// Example usage for getting goals by objective:
//
//	func (h *GoalHandler) GetGoalByObjective(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			objective := vars["objective"]
//			req := handle.GetGoalByObjectiveRequest{Objective: objective}
//			res, err := h.Control.GetGoalByObjective(&req)
//			// Handle error
//			err = json.NewEncoder(w).Encode(res)
//			// Handle error
//	}
//
// Example usage for getting goals by planner ID:
//
//	func (h *GoalHandler) GetGoalsByPlannerIdRequest(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			plannerId := vars["planner_id"]
//			req := handle.GetGoalsByPlannerIdRequest{PlannerID: plannerId}
//			res, err := h.Control.GetGoalsByPlannerId(&req)
//			// Handle error
//			err = json.NewEncoder(w).Encode(res)
//			// Handle error
//	}
//
// Example usage for updating a goal:
//
//	func (h *GoalHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			id := vars["id"]
//
//			var req handle.UpdateGoalRequest
//			err := json.NewDecoder(r.Body).Decode(&req)
//			// Handle error
//
//			req.Id = id // Ensure the ID from the URL is used
//
//			err = h.Control.UpdateGoal(&req)
//			// Handle error
//
//			w.WriteHeader(http.StatusOK)
//	}
//
// Example usage for deleting a goal:
//
//	func (h *GoalHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			id := vars["id"]
//			req := handle.DeleteGoalRequest{Id: id}
//			err := h.Control.DeleteGoal(&req)
//			// Handle error
//			w.WriteHeader(http.StatusOK)
//	}
//
// Example usage for listing all goals:
//
//	func (h *GoalHandler) ListGoals(w http.ResponseWriter, r *http.Request) {
//			res, err := h.Control.ListGoals()
//			// Handle error
//			err = json.NewEncoder(w).Encode(res)
//			// Handle error
//	}
type GoalHandler struct {
	Control *handle.GoalControl
}

// CreateGoal is a method of the GoalHandler struct that handles the creation of a new goal.
// It takes an HTTP response writer and request as parameters.
// The request body is decoded into a handle.CreateGoalRequest struct.
// If there is an error decoding the request body, a Bad Request HTTP response is returned.
// The CreateGoal method of the GoalControl struct is then called with the decoded request as a parameter.
// If there is an error creating the goal, an Internal Server Error HTTP response is returned.
// The response is encoded into the response writer as JSON.
// If there is an error encoding the response, an Internal Server Error HTTP response is returned.
func (h *GoalHandler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var req handle.CreateGoalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreateGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetGoal is a method of GoalHandler that retrieves a specific goal based on the provided ID.
// It takes in a http.ResponseWriter `w` and http.Request `r` as parameters.
// The `w` is used to write the response back to the client, and `r` is used to obtain the ID of the goal from the request URL.
// It begins by extracting the ID from the request URL using mux.Vars(r).
// Then, it creates a GetGoalRequest object with the extracted ID.
// Next, it calls the GetGoal method of the GoalControl instance within the GoalHandler instance, passing the GetGoalRequest object as a parameter.
// The result of the GetGoal method is stored in `res`, and any error that occurred is stored in `err`.
// If an error occurred while executing the GetGoal method, it is handled by calling the handleError function,
// which writes an error response with the corresponding HTTP status code to the client and logs the error.
// If no error occurred, the result `res` is encoded as JSON and written as the response to the client.
// Any error encountered during encoding is also handled by calling the handleError function.
// Example Usage:
//
//	handler := GoalHandler{Control: &GoalControl{Service: &services.GoalService{}}}
//	router.HandleFunc("/goal/{id}", handler.GetGoal).Methods("GET")
//
//	HTTP GET /goal/1 -> calls GetGoal method with ID = "1"
//	Result -> goal object with ID = "1" is returned as response JSON
func (h *GoalHandler) GetGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetGoalRequest{Id: id}
	res, err := h.Control.GetGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetGoalByObjective retrieves a goal by its objective.
// It takes the objective as a URL parameter and sends it to the GoalControl's GetGoalByObjective method.
// If an error occurs while retrieving the goal or encoding the response, it returns a 500 Internal Server Error.
// The retrieved goal is encoded as JSON and written to the http.ResponseWriter.
//
// Example:
//
//	GET /goals?objective=example
//
//	{
//		"goal": {
//			"id": "123",
//			"objective": "example",
//			"deadline": "2022-01-01T00:00:00Z",
//			"created_at": "2021-01-01T00:00:00Z",
//			"updated_at": "2021-01-01T00:00:00Z"
//		}
//	}
func (h *GoalHandler) GetGoalByObjective(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objective := vars["objective"]
	req := handle.GetGoalByObjectiveRequest{Objective: objective}
	res, err := h.Control.GetGoalByObjective(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetGoalsByPlannerIdRequest retrieves a list of goals by planner ID from the goal control.
//
// It takes a HTTP response writer and request as parameters.
//
// It extracts the planner ID from the request URL using Gorilla Mux and creates a GetGoalsByPlannerIdRequest object.
// It then calls the GetGoalsByPlannerId method of the goal control to retrieve the goals.
// Any errors that occur during the process are handled using the handleError function.
// Finally, the retrieved goals are encoded as JSON and written to the response.
//
// Example usage:
//
//	handler := &GoalHandler{
//	    Control: &handle.GoalControl{
//	        Service: &services.GoalService{},
//	    },
//	}
//	http.HandleFunc("/goals/{planner_id}", handler.GetGoalsByPlannerIdRequest)
//	log.Fatal(http.ListenAndServe(":8080", nil))
func (h *GoalHandler) GetGoalsByPlannerIdRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plannerId := vars["planner_id"]
	req := handle.GetGoalsByPlannerIdRequest{PlannerId: plannerId}
	res, err := h.Control.GetGoalsByPlannerId(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// UpdateGoal updates a goal with the provided ID based on the request body.
// It first decodes the request body into an update goal request object.
// Then it ensures that the ID from the URL is used.
// It then calls the UpdateGoal method of the GoalControl service.
// If there is an error during the update, it returns a 500 Internal Server Error response.
// If the update is successful, it returns a 200 OK response.
// If there is an error decoding the request body, it returns a 400 Bad Request response.
//
// Example usage:
//
//	req := handle.UpdateGoalRequest{
//	    Id:        "goalId",
//	    Objective: "New objective",
//	    Deadline:  "2022-01-01",
//	    PlannerID: "plannerId",
//	}
//	h.Control.UpdateGoal(&req)
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

// DeleteGoal deletes a goal based on the ID provided in the request URL.
//
// Example:
//
//	handler := &GoalHandler{Control: &GoalControl{Service: goalService}}
//	router.HandleFunc("/goals/{id}", handler.DeleteGoal).Methods(http.MethodDelete)
//
//	// Request: DELETE /goals/123
//
// This handler function expects a DELETE HTTP request to the "/goals/{id}" URL pattern,
// where "{id}" is the ID of the goal to be deleted.
//
// The function retrieves the ID from the request URL, creates a DeleteGoalRequest with the
// ID, and calls the DeleteGoal method of the GoalControl struct passed in the GoalHandler.
//
// If there is an error while executing the DeleteGoal method, the function calls the handleError
// function to handle the error. If no error occurs, the function writes a 200 OK status code
// to the response writer.
//
// Example response body:
//
//	HTTP/1.1 200 OK
func (h *GoalHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeleteGoalRequest{Id: id}
	err := h.Control.DeleteGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

// ListGoals retrieves a list of goals.
//
// It calls the ListGoals method of GoalControl and encodes the response as JSON.
// If there is an error retrieving the goals, it returns an HTTP 500 Internal Server Error.
//
// Example usage:
//
//	goalHandler := &GoalHandler{Control: goalControl}
//	http.HandleFunc("/goals", goalHandler.ListGoals)
func (h *GoalHandler) ListGoals(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListGoals()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
