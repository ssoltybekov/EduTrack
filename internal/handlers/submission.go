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

type submissionCreateRequest struct {
	dto.SubmissionInputDTO
	StudentID uint `json:"student_id" validate:"required"`
}

type submissionGradeRequest struct {
	dto.SubmissionGradeInputDTO
	TeacherID uint `json:"teacher_id" validate:"required"`
}

func ListSubmissions(w http.ResponseWriter, r *http.Request) {
	assignmentIDStr := chi.URLParam(r, "assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		response.BadRequest(w, "Invalid assignment ID")
		return
	}
	submissions, err := submissionService.GetAllByAssignment(uint(assignmentID))
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
	assignmentIDStr := chi.URLParam(r, "assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		response.BadRequest(w, "Invalid assignment ID")
		return
	}

	var input submissionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	out, err := submissionService.Create(&input.SubmissionInputDTO, uint(assignmentID), input.StudentID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, out)
}

func GradeSubmission(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}

	var input submissionGradeRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	submission, err := submissionService.Grade(uint(id), input.Grade, input.Feedback, input.TeacherID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, submission)
}
