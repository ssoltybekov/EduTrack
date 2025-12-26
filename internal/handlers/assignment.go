package handlers

import (
	"edutrack/internal/dto"
	"edutrack/internal/pkg/response"
	"edutrack/internal/pkg/validator"
	"edutrack/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var assignmentService = services.NewAssignmentService()

func ListAssignments(w http.ResponseWriter, r *http.Request) {
	assignments, err := assignmentService.GetAll()
	if err != nil {
		response.Internal(w)
		return
	}

	response.JSON(w, http.StatusOK, assignments)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}

	assignment, err := assignmentService.GetById(uint(id))
	if err != nil {
		response.FromError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, assignment)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	var input dto.AssignmentInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}

	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	out, err := assignmentService.Create(&input)

	if err != nil {
		response.FromError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, out)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}

	var updated dto.AssignmentInputDTO

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}

	if err := validator.Validator.Struct(updated); err != nil {
		response.ValidationError(w, err)
		return
	}

	assignment, err := assignmentService.Update(uint(id), &updated)
	if err != nil {
		response.FromError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, assignment)

	// if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// if err := assignmentService.Update(&assignment); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(assignment)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}

	if err := assignmentService.Delete(uint(id)); err != nil {
		response.FromError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
