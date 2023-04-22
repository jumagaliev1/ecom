package main

import (
	"github.com/jumagaliev1/internal/data"
	"net/http"
)

// @Summary		Create Order
// @Description	Creat Order for Shop
// @Security		ApiKeyAuth
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			inout	body		data.OrderReq	true	"input"
// @Success		200		{object}	data.Order
// @Failure		422		{object}	Error
// @Failure		404		{object}	Error
// @Failure		500		{object}	Error
// @Router			/order [post]
func (app *application) CreateOrder(w http.ResponseWriter, r *http.Request) {
	input := &data.OrderReq{}

	err := app.readJSON(w, r, input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	cart, err := app.models.Carts.GetByID(input.CartID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	order := &data.Order{
		OrderStatus: data.OrderStatusCreated,
		Cart:        *cart,
		Quantity:    input.Quantity,
		TotalPrice:  cart.Product.Price * cart.Quantity * input.Quantity,
	}

	err = app.models.Orders.Insert(order)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// @Summary		Approve Order
// @Description	Approve Order for Shop
// @Security		ApiKeyAuth
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	string
// @Failure		422	{object}	Error
// @Failure		404	{object}	Error
// @Failure		500	{object}	Error
// @Router			/order/{id} [patch]
func (app *application) ApproveOrder(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	order, err := app.models.Orders.GetByID(int(id))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if order.OrderStatus != data.OrderStatusCreated {
		err = app.writeJSON(w, http.StatusForbidden, envelope{"message": order.OrderStatus}, nil)
	}

	order.OrderStatus = data.OrderStatusFinish

	err = app.writeJSON(w, http.StatusOK, envelope{"message": order.OrderStatus}, nil)
}

// @Summary		Cancel Order
// @Description	Cancel Order for Shop
// @Security		ApiKeyAuth
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	string
// @Failure		404	{object}	Error
// @Failure		500	{object}	Error
// @Router			/order/{id} [delete]
func (app *application) CancelOrder(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	order, err := app.models.Orders.GetByID(int(id))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if order.OrderStatus != data.OrderStatusCreated {
		err = app.writeJSON(w, http.StatusForbidden, envelope{"message": order.OrderStatus}, nil)
	}

	order.OrderStatus = data.OrderStatusCancel

	err = app.writeJSON(w, http.StatusOK, envelope{"message": order.OrderStatus}, nil)
}
