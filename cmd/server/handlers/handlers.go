package handlers

import (
	"encoding/json"
	"net/http"
)

// Hello world handler
func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(
		helloWorld{
			Hello: "world",
		},
	)

	writeResponse(w, http.StatusOK, data)
}

// Healthchech handler
func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(
		healthcheck{
			Ping: "pong",
		},
	)

	writeResponse(w, http.StatusOK, data)
}

// Not found handler
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data, _ := json.Marshal(
			notFound{
				Code:    http.StatusNotFound,
				Message: "Fallback reached",
				Error:   http.StatusText(http.StatusNotFound),
			},
		)

		writeResponse(w, http.StatusNotFound, data)
	})
}

func writeResponse(w http.ResponseWriter, status int, data []byte) {
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(status)
	w.Write(data)
}
