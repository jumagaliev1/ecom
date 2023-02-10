package main

import (
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"github.com/jumagaliev1/internal/validator"
	"net/http"
	"time"
)

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		User        int64      `json:"user"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Price       data.Price `json:"price"`
		Category    int32      `json:"category"`
		Stock       int        `json:"stock"`
		Images      []string   `json:"images"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	product := &data.Product{
		User:        input.User,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Stock:       input.Stock,
		Images:      input.Images,
	}

	v := validator.New()

	if data.ValidateItem(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Products.Insert(product)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/products/%d", product.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"product": product}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	product := data.Product{
		ID:          id,
		Category:    1,
		User:        0,
		Title:       "IPhone 11",
		Description: "Apple phones Memory 64gb",
		Price:       1_500_000,
		Rating:      5,
		Stock:       4,
		Images:      []string{"google.com", "youtube.com"},
		CreatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
