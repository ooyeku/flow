package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ooyeku/flow/pkg/handle"
	"net/http"
)

// PlannerHandler is a type that handles planner-related operations.
// It contains a reference to the PlannerControl that provides the implementation for these operations.
// It has methods to handle the following HTTP requests:
// - CreatePlanner: handles the creation of a new planner.
// - GetPlanner: handles the retrieval of a planner by ID.
// - GetPlannerByTitle: handles the retrieval of a planner by title.
// - GetPlannerByOwner: handles the retrieval of planners by owner.
// - UpdatePlanner: handles the updating of an existing planner.
// - DeletePlanner: handles the deletion of a planner.
type PlannerHandler struct {
	Control *handle.PlannerControl
}

// CreatePlanner takes an HTTP response writer and request as input.
// It decodes the request body into a CreatePlannerRequest object.
// If any error occurs during the decoding, it is handled by the handleError function.
// It then sends the CreatePlannerRequest to the PlannerControl's CreatePlanner method to create a new planner.
// If any error occurs during the creation process, it is handled by the handleError function.
// Finally, it encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
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

// GetPlanner takes an HTTP response writer and request as input.
// It sets the "Content-Type" header of the response to "application/json".
// It extracts the "id" variable from the request's route parameters.
// It creates a GetPlannerRequest with the extracted id.
// It sends the GetPlannerRequest to the PlannerControl's GetPlanner method to retrieve the planner.
// If any error occurs during the retrieval process, it is handled by the handleError function.
// It encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
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

// GetPlannerByTitle takes an HTTP response writer and request as input.
// It extracts the "title" variable from the request URL path and assigns it to the "title" variable.
// It then creates a GetPlannerByTitleRequest object with the extracted title.
// It sends the GetPlannerByTitleRequest to the PlannerControl's GetPlannerByTitle method to retrieve the planner with the given title.
// If any error occurs during the retrieval process, it is handled by the handleError function.
// Finally, it encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
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

// GetPlannerByOwner takes an HTTP response writer and request as input.
// It sets the "Content-Type" header to "application/json".
// It retrieves the "owner" variable from the request URL parameters.
// It creates a GetPlannerByOwnerRequest object with the UserId field set to the value of "owner".
// It sends the GetPlannerByOwnerRequest to the PlannerControl's GetPlannerByOwner method to retrieve the planner by owner.
// If any error occurs during the retrieval process, it is handled by the handleError function.
// Finally, it encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
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

// UpdatePlanner takes an HTTP response writer and request as input.
// It sets the content type header of the response writer to application/json.
// It gets the "id" parameter from the request URL using mux.Vars.
// It creates an UpdatePlannerRequest object with the "id" parameter set.
// It decodes the request body into the UpdatePlannerRequest object.
// If any error occurs during the decoding, it is handled by the handleError function.
// It sends the UpdatePlannerRequest to the PlannerControl's UpdatePlanner method to update the planner.
// If any error occurs during the update process, it is handled by the handleError function.
// It sets the HTTP status code to 200 (OK).
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

// DeletePlanner takes an HTTP response writer and request as input.
// It sets the "Content-Type" header of the response to "application/json".
// It gets the "id" parameter from the request URL using the mux.Vars() function.
// It creates a DeletePlannerRequest object with the extracted id and passes it to h.Control.DeletePlanner().
// If any error occurs during the deletion process, it is handled by the handleError function.
// It sets the HTTP response writer status code to 200 (OK).
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

// ListPlanners takes an HTTP response writer and request as input.
// It sets the "Content-Type" header of the response to "application/json".
// It then calls the ListPlanners method of the PlannerControl to retrieve a list of planners.
// If any error occurs during the retrieval process, it is handled by the handleError function.
// Finally, it encodes the response using JSON and writes it to the HTTP response writer.
// If any error occurs during the encoding process, it is handled by the handleError function.
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
