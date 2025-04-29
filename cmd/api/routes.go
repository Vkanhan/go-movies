package api

import (
	"net/http"


)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/v1/healthcheck", app.healthcheck)
	router.HandleFunc("/v1/movies", app.createMovie)
	router.HandleFunc("/v1/movies/{id}", app.showMovie)

	return router
}
