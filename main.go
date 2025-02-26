package main

import (
	"log"
	"net/http"
	"os"

	"search-service/handlers"
	"search-service/middleware"
	"search-service/opensearch"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize OpenSearch client
	opensearch.Init()

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.CSRFMiddleware)
	r.Use(middleware.XSSMiddleware)
	r.Use(middleware.LoggingMiddleware)

	// Define routes
	r.HandleFunc("/search", handlers.SearchHandler).Methods("GET")
	r.HandleFunc("/suggest", handlers.SuggestHandler).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
