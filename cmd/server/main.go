package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joaopedrocampos/go-simple-api/cmd/server/handlers"
)

func main() {
	// Hello world web server

	initializeRouter()

	port := os.Getenv("PORT")

	if port != "" {
		port = ":" + port
	} else {
		port = ":8080"
	}

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/ping", handlers.HealthcheckHandler).Methods("GET")

	r.NotFoundHandler = handlers.NotFoundHandler()

	http.Handle("/", r)
}
