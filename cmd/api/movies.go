package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"greenlight.tarala.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// fields must be exported (i.e. start with Capital Letters)

	var input struct {
		Title   string   `json:"title"`
		Runtime data.Runtime    `json:"runtime"`
		Genres  []string `json:"genres"`
		Year    int32    `json:"year"`
	}

	// &input doesn't have to be initialized - it's already initialized above with empty values
	// decoder check struct tag names above. If they don't match - it tries to match key names case-insensitive
	// json values, which can't be mapped, will be silently ignored 
	// important - it MUST BE non-nil pointer
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

//this is exactly the same as json.Decode(). But more  verbose and less efficient 80% more memory 
// thus we are going to use createMovieHandler method later
func (app *application) createMovieHandlerWithUnmarshall(w http.ResponseWriter, r *http.Request) {
	// fields must be exported (i.e. start with Capital Letters)

	var input struct {
		Title   string   `json:"title"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
		Year    int32    `json:"year"`
	}

	body, err:= io.ReadAll(r.Body)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//same here - target must be a struct pointer. What actually makes sense, as this method doesn't return anything  
	err = json.Unmarshal(body, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}


	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := &data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
