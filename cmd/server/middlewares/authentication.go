package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type unauthorizedError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Middleware responsible for authentication
func AuthenticationMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Token ")

			tokenValid, errorMessage := validateToken(authHeader)

			if tokenValid == false {
				writeUnauthorizedResponse(w, errorMessage)
			} else {
				next.ServeHTTP(w, r)
			}

		})
	}
}

func validateToken(authHeader []string) (bool, string) {
	if len(authHeader) != 2 {
		malformedTokenMessage := "Malformed token"
		log.Println(malformedTokenMessage)

		return false, malformedTokenMessage
	}

	if authHeader[1] != "abcd1234" {
		invalidTokenMessage := "Invalid access token"
		log.Println(invalidTokenMessage)

		return false, invalidTokenMessage
	}

	return true, ""
}

func writeUnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusUnauthorized)
	data, _ := json.Marshal(
		unauthorizedError{
			Code:    http.StatusUnauthorized,
			Error:   http.StatusText(http.StatusUnauthorized),
			Message: message,
		},
	)
	w.Write(data)
}
