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

var teacherService = services.NewTeacherService()

func ListTeacher(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherService.GetAll()
	if err != nil {
		response.Internal(w)
		return
	}
	response.JSON(w, http.StatusOK, teachers)
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	teacher, err := teacherService.GetById(uint(id))
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, teacher)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var input dto.TeacherInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	out, err := teacherService.Create(&input)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, out)
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	var updated dto.TeacherInputDTO
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(updated); err != nil {
		response.ValidationError(w, err)
		return
	}

	teacher, err := teacherService.Update(uint(id), &updated)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, teacher)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	if err := teacherService.Delete(uint(id)); err != nil {
		response.FromError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}