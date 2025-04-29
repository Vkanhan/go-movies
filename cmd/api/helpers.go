package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"maps"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, errors.New("id parameter missing")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil 
}

type envelop map[string]any

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload any, headers http.Header) error {
	data, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err 
	}

	data = append(data, '\n')

	//range over add header to header map
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

	return nil 
}
