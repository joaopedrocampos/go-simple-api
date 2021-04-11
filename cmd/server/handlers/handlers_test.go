package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func TestHelloWorldHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloWorldHandler).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(r).
		Get("/hello").
		Expect(t).
		Body(`{"hello": "world"}`).
		Status(http.StatusOK).
		End()
}

func TestHealthcheckHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HealthcheckHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(r).
		Get("/ping").
		Expect(t).
		Body(`{"ping": "pong"}`).
		Status(http.StatusOK).
		End()
}
