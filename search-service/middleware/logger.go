package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests",
	}, []string{"method", "path", "status"})

	requestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "path", "status"})
)

// LoggingMiddleware logs all incoming requests and outgoing responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		latency := time.Since(start).Seconds()
		status := rw.status

		// Log the outgoing response
		log.Printf("Outgoing response: %s %s %d %fs", r.Method, r.URL.Path, status, latency)

		// Record Prometheus metrics
		requestDuration.WithLabelValues(r.Method, r.URL.Path, http.StatusText(status)).Observe(latency)
		requestCount.WithLabelValues(r.Method, r.URL.Path, http.StatusText(status)).Inc()
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
