package main

import (
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Handler(http.MethodGet, "/v1/healthcheck", app.authenticate(http.HandlerFunc(app.healthcheckHandler)))
	router.Handler(http.MethodGet, "/v1/products", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.listProductsHandler))))
	router.Handler(http.MethodPost, "/v1/products", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createProductHandler))))
	router.Handler(http.MethodGet, "/v1/products/:id", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.showProductHandler))))
	router.Handler(http.MethodPatch, "/v1/products/:id", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.updateProductHandler))))
	router.Handler(http.MethodDelete, "/v1/products/:id", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.deleteProductHandler))))

	router.Handler(http.MethodPost, "/v1/comment", app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createCommentHandler))))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	//return app.recoverPanic(app.authenticate(router))
	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4000/static/swagger.json")))

	router.ServeFiles("/static/*filepath", http.Dir("docs"))
	return app.recoverPanic(router)
}
