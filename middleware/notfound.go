package middleware

import (
	"net/http"
	"search-service/utils"

	"github.com/gorilla/mux"
)

// NotFoundMiddleware returns a 404 response if the route is not found
func NotFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the route exists
		var routeMatch mux.RouteMatch
		if !mux.CurrentRoute(r).Match(r, &routeMatch) {
			// If the route does not exist, return a 404 response

			utils.WriteErrorResponse(w, http.StatusNotFound, "404 Not Found", nil)
			return
		}

		// If the route exists, call the next handler
		next.ServeHTTP(w, r)
	})
}
