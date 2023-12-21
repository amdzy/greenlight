package main

import (
	"net/http"
	"time"

	"github.com/Soul-Remix/greenlight/internal/data"
	"github.com/Soul-Remix/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateMovie(v, &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	})

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	app.writeJSON(w, r, http.StatusCreated, envelope{"input": input}, nil)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		Id:        id,
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
		CreatedAt: time.Now(),
	}

	app.writeJSON(w, r, http.StatusOK, envelope{"movie": movie}, nil)
}
