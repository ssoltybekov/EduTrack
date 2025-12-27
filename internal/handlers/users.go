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

var userService = services.NewUserService()

func ListUsers(w http.ResponseWriter, r *http.Request) {
    users, err := userService.GetAll()
    if err != nil {
        response.Internal(w)
        return
    }
    response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.BadRequest(w, "Invalid ID")
        return
    }
    user, err := userService.GetById(uint(id))
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var input dto.UserInputDTO
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        response.BadRequest(w, "Invalid JSON")
        return
    }
    if err := validator.Validator.Struct(input); err != nil {
        response.ValidationError(w, err)
        return
    }
    out, err := userService.Create(&input)
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusCreated, out)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.BadRequest(w, "Invalid ID")
        return
    }
    var input dto.UserInputDTO
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        response.BadRequest(w, "Invalid JSON")
        return
    }
    if err := validator.Validator.Struct(input); err != nil {
        response.ValidationError(w, err)
        return
    }
    user, err := userService.Update(uint(id), &input)
    if err != nil {
        response.FromError(w, err)
        return
    }
    response.JSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.BadRequest(w, "Invalid ID")
        return
    }
    if err := userService.Delete(uint(id)); err != nil {
        response.FromError(w, err)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}