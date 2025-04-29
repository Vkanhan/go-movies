package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vkanhan/go-movies/internal/data"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year int `json:"year"`
		Showtime int32 `json:"showtime"`
		Genres []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.serverError(w, r, err)
		return 
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return 
	}

	movie := &data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Interstellar",
		Showtime: 100,
		Genres: []string{"crime", "romance", "action"},
		Version: 1,
	}

	err = app.respondWithJSON(w, http.StatusOK, envelop{"movie": movie}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}




