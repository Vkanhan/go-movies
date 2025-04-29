package api

import (
	"fmt"
	"net/http"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New movie")
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show movie")
}


