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
	dynamodbEndpoint := os.Getenv("DYNAMODB_ENDPOINT")

	clients.NewOpenSearchClient(env, opensearchEndpoint, awsRegion)
	clients.InitializeDynamoDBClient(awsRegion, dynamodbEndpoint)

	r := mux.NewRouter()

	r.Use(middleware.CORSMiddleware)

	if env != "local" {
		r.Use(middleware.CSRFMiddleware)
		r.Use(middleware.XSSMiddleware)
	}
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/api/v1").Subrouter()

	courses := api.PathPrefix("/course").Subrouter()

	courses.HandleFunc("/search", handlers.SearchCourseHandler).Methods("GET")
	courses.HandleFunc("/suggest", handlers.SuggestCourseHandler).Methods("GET")
	courses.HandleFunc("/{id}", handlers.GetCourseHandler).Methods("GET")
	courses.HandleFunc("", handlers.CreateCourseHandler).Methods("POST", "OPTIONS")

	r.Handle("/metrics", promhttp.Handler())

	// Apply the NotFoundMiddleware
	r.Use(middleware.NotFoundMiddleware)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
