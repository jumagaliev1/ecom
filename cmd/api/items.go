package main

import (
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"github.com/jumagaliev1/internal/validator"
	"net/http"
	"time"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string     `json:"title"`
		Price    data.Price `json:"price"`
		Category int32      `json:"category"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	item := &data.Item{
		Title:    input.Title,
		Price:    input.Price,
		Category: input.Category,
	}

	v := validator.New()

	if data.ValidateItem(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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
