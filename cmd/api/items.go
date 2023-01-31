package main

import (
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"net/http"
	"time"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new item")
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
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
