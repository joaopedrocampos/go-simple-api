package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joaopedrocampos/go-simple-api/cmd/server/middlewares"
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
		Header("Authorization", "Token abcd1234").
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
		Header("Authorization", "Token abcd1234").
		Expect(t).
		Body(`{"ping": "pong"}`).
		Status(http.StatusOK).
		End()
}

func TestNotFoundHandler(t *testing.T) {
	r := mux.NewRouter()
	r.NotFoundHandler = NotFoundHandler()

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(r).
		Get("/notfound").
		Expect(t).
		Body(`{"code":404,"error":"Not Found","message":"Fallback reached"}`).
		Status(404).
		End()
}

func TestMissingAuthorizationHeaders(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HealthcheckHandler)
	r.Use(middlewares.AuthenticationMiddleware())

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(r).
		Get("/ping").
		Expect(t).
		Body(`{"code":401,"error":"Unauthorized","message":"Malformed token"}`).
		Status(http.StatusUnauthorized).
		End()
}

func TestInvalidToken(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HealthcheckHandler)
	r.Use(middlewares.AuthenticationMiddleware())

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(r).
		Get("/ping").
		Header("Authorization", "Token batata").
		Expect(t).
		Body(`{"code":401,"error":"Unauthorized","message":"Invalid access token"}`).
		Status(http.StatusUnauthorized).
		End()
}
