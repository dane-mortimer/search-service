package handlers

import (
	"encoding/json"
	"net/http"

	"search-service/controllers"
	"search-service/models"
	"search-service/utils"

	"github.com/gorilla/mux"
)

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course

	// Decode the request body into the course struct
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result, err := controllers.CreateCourseController(course)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error creating course", nil)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, result, nil)
}

func GetCourseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := controllers.GetCourseController(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error getting course", nil)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, result, nil)
}

func SearchCourseHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	fields := []string{"title"}

	searchResult, totalItems, err := controllers.SearchCourseController(query, pageStr, sizeStr, fields)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error searching documents", nil)
		return
	}

	pagination := utils.NewPagination(pageStr, sizeStr, totalItems)

	utils.WriteSuccessResponse(w, http.StatusOK, searchResult, &pagination)
}

func SuggestCourseHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	suggestions, err := controllers.SuggestCourseController(query)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error fetching suggestions", nil)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, suggestions, nil)
}
