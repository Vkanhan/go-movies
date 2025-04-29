package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(_ *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelop{"error": message}

	err := app.respondWithJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := fmt.Sprintf("server error")
	app.errResponse(w, r, http.StatusNotFound, message)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("resource not found")
	app.errResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("method not allowed")
	app.errResponse(w, r, http.StatusMethodNotAllowed, message)
}
