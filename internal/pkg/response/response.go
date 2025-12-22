package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Details map[string]string `json:"details,omitempty"`
	} `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ValidationError(w http.ResponseWriter, err error) {
	details := make(map[string]string)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			details[fe.Field()] = fe.Tag()
		}
	} else {
		details["error"] = err.Error()
	}

	res := ErrorResponse{}
	res.Error.Code = "VALIDATION_ERROR"
	res.Error.Message = "invalid input"
	res.Error.Details = details

	JSON(w, http.StatusBadRequest, res)
}

func BadRequest(w http.ResponseWriter, msg string) {
	res := ErrorResponse{}
	res.Error.Code = "BAD_REQUEST"
	res.Error.Message = msg
	JSON(w, http.StatusBadRequest, res)
}

func NotFound(w http.ResponseWriter, msg string) {
	res := ErrorResponse{}
	res.Error.Code = "NOT_FOUND"
	res.Error.Message = msg
	JSON(w, http.StatusNotFound, res)
}

func Forbidden(w http.ResponseWriter, msg string) {
	res := ErrorResponse{}
	res.Error.Code = "FORBIDDEN"
	res.Error.Message = msg
	JSON(w, http.StatusForbidden, res)
}

func Internal(w http.ResponseWriter) {
	res := ErrorResponse{}
	res.Error.Code = "INTERNAL_ERROR"
	res.Error.Message = "internal server error"
	JSON(w, http.StatusInternalServerError, res)
}
