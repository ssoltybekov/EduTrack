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

var submissionService = services.NewSubmissionService()

func ListSubmissions(w http.ResponseWriter, r *http.Request) {
	submissions, err := submissionService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submissions)
}

func GetSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	submission, err := submissionService.GetById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submission)
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	var input dto.SubmissionInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	submission, err := submissionService.Create(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(submission)
}

func UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var input dto.SubmissionInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	submission, err := submissionService.Update(uint(id), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submission)
}

func DeleteSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := submissionService.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
