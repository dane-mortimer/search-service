package main

import (
	"log"
	"net/http"
	"os"

	"search-service/clients"
	"search-service/handlers"
	"search-service/middleware"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	env := os.Getenv("ENV")
	awsRegion := os.Getenv("AWS_REGION")
	opensearchEndpoint := os.Getenv("OPENSEARCH_ENDPOINT")

	// Initialize OpenSearch client
	clients.NewOpenSearchClient(env, opensearchEndpoint, awsRegion)

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.CORSMiddleware)
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
