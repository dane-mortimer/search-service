package handlers

import (
	"net/http"

	"search-service/services"
	"search-service/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	fields := []string{"title"}

	// Perform the search
	searchResult, totalItems, err := services.Search(query, pageStr, sizeStr, fields)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error searching documents", nil)
		return
	}

	// Create pagination details
	pagination := utils.NewPagination(pageStr, sizeStr, totalItems)

	utils.WriteSuccessResponse(w, http.StatusOK, searchResult, &pagination)
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	suggestions, err := services.Suggest(query)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error fetching suggestions", nil)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, suggestions, nil)
}
