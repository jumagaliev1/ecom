package main

import (
	"encoding/json"
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"net/http"
	"time"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
		Price    int64  `json:"price"`
		Category int32  `json:"category"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	item := data.Item{
		ID:          id,
		CreatedAt:   time.Now(),
		Title:       "New Apartment above Mega",
		Price:       1_500_000,
		IsPurchased: false,
		Category:    1,
		Rating:      5,
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
