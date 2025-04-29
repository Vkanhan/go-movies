package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheck)
	router.HandleFunc("/v1/movies", app.createMovie)
	router.HandleFunc("/v1/movies/{id}", app.showMovie)

	return router
}
