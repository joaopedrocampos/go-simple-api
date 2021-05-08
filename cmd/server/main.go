package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joaopedrocampos/go-simple-api/cmd/server/handlers"
	"github.com/joaopedrocampos/go-simple-api/cmd/server/middlewares"
)

func main() {
	// Hello world web server

	initializeRouter()

	port := portHandler()

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initializeRouter() {
	r := mux.NewRouter()

	// register routes
	r.HandleFunc("/hello", handlers.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/ping", handlers.HealthcheckHandler).Methods("GET")

	// register middlewares
	r.Use(middlewares.RequestUriMiddleware())

	// not found handler
	r.NotFoundHandler = handlers.NotFoundHandler()

	http.Handle("/", r)
}

// port should have the pattern :{PORT}, e.g: :8080
func portHandler() string {
	port := os.Getenv("PORT")

	if port != "" {
		return ":" + port
	}

	return ":8080"
}
