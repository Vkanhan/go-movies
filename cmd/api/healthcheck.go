package main

import (
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	env := envelop{
		"status":  "available",
		"system_info": map[string]string{
			"env":     app.config.env,
			"version": version,
		},
	}

	err := app.respondWithJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

}
