package main

import (
	"errors"
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"github.com/jumagaliev1/internal/validator"
	"net/http"
)

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID    int    `json:"user_id"`
		ProductID int    `json:"product_id"`
		Message   string `json:"message"`
		Rating    uint8  `json:"rating"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user, err := app.models.Users.GetByID(input.UserID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	product, err := app.models.Products.Get(int64(input.ProductID))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	comment := &data.Comment{
		User:    *user,
		Product: *product,
		Message: input.Message,
		Rating:  input.Rating,
	}
	v := validator.New()

	if data.ValidateComment(v, comment); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Comments.Insert(comment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	product.AllRating += int(comment.Rating)
	product.CountRating++
	product.Rating = float32(product.AllRating) / float32(product.CountRating)

	err = app.models.Products.Update(product)
	if err != nil {
		// to-do
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/products/%d", product.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"comment": comment}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
