package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	response := Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ValidationErrorResponse(w http.ResponseWriter, errors map[string]string) {
	response := Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Validation failed",
		Data:       errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
