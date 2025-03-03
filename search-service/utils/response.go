package utils

import (
	"encoding/json"
	"net/http"
	"search-service/models"
)

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, pagination *models.Pagination) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.SuccessResponse{
		Status:     "success",
		Data:       data,
		Pagination: pagination,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, details interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    statusCode,
		Details: details,
	}

	json.NewEncoder(w).Encode(response)
}
