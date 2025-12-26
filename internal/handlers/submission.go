package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"edutrack/internal/dto"
	"edutrack/internal/pkg/response"
	"edutrack/internal/pkg/validator"
	"edutrack/internal/services"

	"github.com/go-chi/chi"
)

var submissionService = services.NewSubmissionService()

func ListSubmissions(w http.ResponseWriter, r *http.Request) {
	submissions, err := submissionService.GetAll()
	if err != nil {
		response.Internal(w)
		return
	}
	response.JSON(w, http.StatusOK, submissions)
}

func GetSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	submission, err := submissionService.GetById(uint(id))
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, submission)
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	var input dto.SubmissionInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	out, err := submissionService.Create(&input)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, out)
}

func UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	var updated dto.SubmissionInputDTO
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(updated); err != nil {
		response.ValidationError(w, err)
		return
	}

	submission, err := submissionService.Update(uint(id), &updated)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, submission)
}

func DeleteSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	if err := submissionService.Delete(uint(id)); err != nil {
		response.FromError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
