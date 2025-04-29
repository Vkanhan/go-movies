package api

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status available")
	fmt.Fprintf(w, "env: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %d\n", version)
}
