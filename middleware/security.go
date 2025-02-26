package middleware

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/microcosm-cc/bluemonday"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	return csrf.Protect([]byte("32-byte-long-auth-key"))(next)
}

func XSSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sanitizer := bluemonday.UGCPolicy()
		for key, values := range r.URL.Query() {
			for i, value := range values {
				r.URL.Query()[key][i] = sanitizer.Sanitize(value)
			}
		}
		next.ServeHTTP(w, r)
	})
}
