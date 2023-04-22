package main

import (
	"errors"
	"github.com/jumagaliev1/internal/data"
	"github.com/jumagaliev1/internal/validator"
	"net/http"
)

//	@Summary		Register User
//	@Description	Registaration user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			input	body		data.UserRegisterInput	true	"Input for remove user"
//	@Success		201		{object}	data.User
//	@Failure		400		{object}	Error
//	@Failure		422		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/users [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input data.UserRegisterInput
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Role:      input.Role,
	}
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return

	}
	//token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
