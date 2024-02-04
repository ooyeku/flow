package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ooyeku/flow/pkg/handle"
	"net/http"
)

// PlanHandler is a type that handles HTTP requests related to plans.
// It contains a control object of type *handle.PlanControl, which is used to interact with the plan service.
//
// CreatePlan handles the creation of a new plan.
// It decodes the request body into a CreatePlanRequest, calls the CreatePlan method on PlanControl,
// and encodes the response in the HTTP response writer.
//
// UpdatePlan handles the updating of an existing plan.
// It decodes the request body into an UpdatePlanRequest, calls the UpdatePlan method on PlanControl,
// and returns any errors that occur during the process.
//
// DeletePlan handles the deletion of a plan.
// It extracts the plan ID from the request URL, creates a DeletePlanRequest with the ID,
// calls the DeletePlan method on PlanControl, and returns any errors that occur during the process.
//
// GetPlan handles the retrieval of a plan by ID.
// It extracts the plan ID from the request URL, creates a GetPlanRequest with the ID,
// calls the GetPlan method on PlanControl, and returns the retrieved plan or any errors that occur during the process.
//
// GetPlanByName handles the retrieval of a plan by name.
// It extracts the plan name from the request URL, creates a GetPlanByNameRequest with the name,
// calls the GetPlanByName method on PlanControl, and returns the retrieved plan or any errors that occur during the process.
//
// GetPlansByGoal handles the retrieval of all plans associated with a goal.
// It extracts the goal ID from the request URL, creates a GetPlansByGoalRequest with the ID,
// calls the GetPlansByGoal method on PlanControl, and returns the retrieved plans or any errors that occur during the process.
//
// ListPlans handles the listing of all plans.
// It calls the ListPlans method on PlanControl and returns the retrieved plans or any errors that occur during the process.
type PlanHandler struct {
	Control *handle.PlanControl
}

// CreatePlan takes an HTTP response writer and request as input.
// It decodes the request body into a CreatePlanRequest object.
// If any error occurs during the decoding, it is handled by the handleError function.
// It then sends the CreatePlanRequest to the PlanControl's CreatePlan method to create a new plan.
// If any error occurs during the creation process, it is handled by the handleError function.
// Finally, it encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
//
// Example usage:
// var h PlanHandler
// h.CreatePlan(w, r)
func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var req handle.CreatePlanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	handleError(w, err, http.StatusBadRequest)
	res, err := h.Control.CreatePlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetPlan takes an HTTP response writer and request as input.
// It retrieves the "id" parameter from the URL, creates a GetPlanRequest,
// and sends it to the PlanControl's GetPlan method to retrieve the plan information.
// The function then encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the process, it is handled by the handleError function.
func (h *PlanHandler) GetPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.GetPlanRequest{Id: id}
	res, err := h.Control.GetPlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetPlanByName gets a plan by its name
func (h *PlanHandler) GetPlanByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planName := vars["plan_name"]
	req := handle.GetPlanByNameRequest{PlanName: planName}
	res, err := h.Control.GetPlanByName(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// GetPlansByGoal retrieves plans based on a goal ID
// It takes in a response writer, a request object, and extracts the goal ID from the request URL
// It creates a GetPlansByGoalRequest with the extracted goal ID
// It then calls the GetPlansByGoal method on the PlanControl instance, passing in the request object
// Any error encountered during the process is handled using the handleError function
// The response is encoded into JSON format and sent in the response writer
// Any error encountered during the encoding process is handled using the handleError function
func (h *PlanHandler) GetPlansByGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId := vars["goal_id"]
	req := handle.GetPlansByGoalRequest{GoalId: goalId}
	res, err := h.Control.GetPlansByGoal(&req)
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}

// UpdatePlan updates an existing plan
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

// DeletePlan deletes a plan by its ID
func (h *PlanHandler) DeletePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	req := handle.DeletePlanRequest{Id: id}
	err := h.Control.DeletePlan(&req)
	handleError(w, err, http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
}

// ListPlans fetches a list of plans
// and encodes them as JSON before returning the response
// to the HTTP client
func (h *PlanHandler) ListPlans(w http.ResponseWriter, r *http.Request) {
	res, err := h.Control.ListPlans()
	handleError(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(res)
	handleError(w, err, http.StatusInternalServerError)
}
