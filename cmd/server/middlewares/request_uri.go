package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Authentication middleware
func RequestUriMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Client request URI:", r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
