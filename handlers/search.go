package handlers

import (
	"encoding/json"
	"net/http"

	"search-service/services"
	"search-service/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	fields := []string{"title", "content", "description"}

	// Perform the search
	searchResult, totalItems, err := services.Search(query, pageStr, sizeStr, fields)
	if err != nil {
		http.Error(w, "Error searching documents", http.StatusInternalServerError)
		return
	}

	// Create pagination details
	pagination := utils.NewPagination(pageStr, sizeStr, totalItems)

	// Prepare the response
	response := map[string]interface{}{
		"data":       searchResult,
		"pagination": pagination,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
