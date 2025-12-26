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

var studentService = services.NewStudentService()

func ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := studentService.GetAll()
	if err != nil {
		response.Internal(w)
		return
	}
	response.JSON(w, http.StatusOK, students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	student, err := studentService.GetById(uint(id))
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var input dto.StudentInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(input); err != nil {
		response.ValidationError(w, err)
		return
	}

	out, err := studentService.Create(&input)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, out)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	var updated dto.StudentInputDTO
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		response.BadRequest(w, "Invalid JSON")
		return
	}
	if err := validator.Validator.Struct(updated); err != nil {
		response.ValidationError(w, err)
		return
	}

	student, err := studentService.Update(uint(id), &updated)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(w, "Invalid ID")
		return
	}
	if err := studentService.Delete(uint(id)); err != nil {
		response.FromError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}