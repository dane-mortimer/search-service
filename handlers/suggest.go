package handlers

import (
	"encoding/json"
	"net/http"
	"search-service/services"
)

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	suggestions, err := services.Suggest(query)
	if err != nil {
		http.Error(w, "Error fetching suggestions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
