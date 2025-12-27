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

var assignmentService = services.NewAssignmentService()

func ListAssignments(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := chi.URLParam(r, "lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		response.BadRequest(w, "Invalid lesson ID")
		return
	}
	assignments, err := assignmentService.GetAllByLesson(uint(lessonID))
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
	lessonIDStr := chi.URLParam(r, "lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		response.BadRequest(w, "Invalid lesson ID")
		return
	}
	var req struct {
		dto.AssignmentInputDTO
		TeacherID uint `json:"teacher_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(req); err != nil {
		response.ValidationError(w, err)
		return
	}
	out, err := assignmentService.Create(&req.AssignmentInputDTO, uint(lessonID), req.TeacherID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, out)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	var req struct {
		dto.AssignmentInputDTO
		TeacherID uint `json:"teacher_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(req); err != nil {
		response.ValidationError(w, err)
		return
	}
	assignment, err := assignmentService.Update(uint(id), &req.AssignmentInputDTO, req.TeacherID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, assignment)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	var req struct {
		TeacherID uint `json:"teacher_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := assignmentService.Delete(uint(id), req.TeacherID); err != nil {
		response.FromError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
