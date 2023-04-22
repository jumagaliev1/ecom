package main

import (
	"github.com/jumagaliev1/internal/data"
	"net/http"
)

type CartReq struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

//	@Summary		Create Cart
//	@Description	Creat Cart for Shop
//	@Security		ApiKeyAuth
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			inout	body		CartReq	true	"input"
//	@Success		200		{object}	data.Cart
//	@Failure		422		{object}	Error
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/cart [post]
func (app *application) CreateCart(w http.ResponseWriter, r *http.Request) {
	input := &data.CartReq{}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)
	product, err := app.models.Products.Get(int64(input.ProductID))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	cart := &data.Cart{
		User:     *user,
		Product:  *product,
		Quantity: input.Quantity,
	}

	err = app.models.Carts.Insert(cart)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"cart": cart}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
