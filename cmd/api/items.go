package main

import (
	"errors"
	"fmt"
	"github.com/jumagaliev1/internal/data"
	"github.com/jumagaliev1/internal/validator"
	"net/http"
)

// @Summary      Create Product
// @Description  Creat Product for Shop
// @Security	ApiKeyAuth
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        inout body  data.InputCreateProduct  true  "input"
// @Success      200  {object}  data.Product
// @Failure      422  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [post]
func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	input := &data.InputCreateProduct{}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := app.contextGetUser(r)
	product := &data.Product{
		User:        user.ID,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Stock:       input.Stock,
		Images:      input.Images,
	}

	v := validator.New()

	if data.ValidateProduct(v, product); !v.Valid() {
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

// @Summary      Show Product
// @Description  Show Product for Shop
// @Security	ApiKeyAuth
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id path int true  "Product ID"
// @Success      200  {object}  data.Product
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	product, err := app.models.Products.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	comments, err := app.models.Comments.GetByProduct(product)
	if err != nil && !errors.Is(err, data.ErrRecordNotFound) {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"product": product, "comments": comments}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary      Update Product
// @Description  Update Product for Shop
// @Security	ApiKeyAuth
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input body data.InputUpdateProduct true  "Input"
// @Success      200  {object}  data.Product
// @Failure      422  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [patch]
func (app *application) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	product, err := app.models.Products.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	input := &data.InputUpdateProduct{}
	user := app.contextGetUser(r)
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Category != nil {
		product.Category = *input.Category
	}

	product.User = user.ID

	if input.Title != nil {
		product.Title = *input.Title
	}

	if input.Description != nil {
		product.Description = *input.Description
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.Rating != nil {
		product.Rating = *input.Rating
	}

	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	if input.Images != nil {
		product.Images = input.Images
	}

	v := validator.New()

	if data.ValidateProduct(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Products.Update(product)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary      Delete Product
// @Description  Delete Product for Shop
// @Security	ApiKeyAuth
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id path int true  "Product ID"
// @Success      200  {object}  string
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [delete]
func (app *application) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Products.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "product successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary      List of Products
// @Description  All list of Products on Shop
// @Security	ApiKeyAuth
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        title query string false  "title"
// @Param        category query int false  "category"
// @Param        page query int false  "page"
// @Param        page_size query int false  "Page size"
// @Param        sort query string false  "sort"
// @Success      200  {object}  []data.Product
// @Failure      422  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [get]
func (app *application) listProductsHandler(w http.ResponseWriter, r *http.Request) {
	input := &data.InputListProducts{}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Category = app.readInt(qs, "category", 0, v)

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "category", "price", "rating", "-id", "-title", "-category", "-price", "-rating"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	products, metadata, err := app.models.Products.GetAll(input.Title, input.Category, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"products": products, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
