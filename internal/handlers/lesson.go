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

var lessonService = services.NewLessonService()

type lessonRequest struct {
    dto.LessonInputDTO
    TeacherID uint `json:"teacher_id" validate:"required"`
}

func ListLessons(w http.ResponseWriter, r *http.Request) {
    lessons, err := lessonService.GetAll()
    if err != nil {
        response.Internal(w)
        return
    }
    response.JSON(w, http.StatusOK, lessons)
}

func GetLesson(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.BadRequest(w, "Invalid ID")
        return
    }
    lesson, err := lessonService.GetById(uint(id))
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusOK, lesson)
}

func CreateLesson(w http.ResponseWriter, r *http.Request) {
    var input lessonRequest
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        response.BadRequest(w, "Invalid JSON")
        return
    }
    if err := validator.Validator.Struct(input); err != nil {
        response.ValidationError(w, err)
        return
    }

    out, err := lessonService.Create(&input.LessonInputDTO, input.TeacherID)
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusCreated, out)
}

func UpdateLesson(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.BadRequest(w, "Invalid ID")
        return
    }

    var input lessonRequest
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        response.BadRequest(w, "Invalid JSON")
        return
    }
    if err := validator.Validator.Struct(input); err != nil {
        response.ValidationError(w, err)
        return
    }

    lesson, err := lessonService.Update(uint(id), &input.LessonInputDTO, input.TeacherID)
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusOK, lesson)
}

func DeleteLesson(w http.ResponseWriter, r *http.Request) {
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
        response.BadRequest(w, "Invalid JSON: teacher_id required")
        return
    }

    if err := lessonService.Delete(uint(id), req.TeacherID); err != nil {
        response.FromError(w, err)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}