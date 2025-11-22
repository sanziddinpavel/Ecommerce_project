package rest

import (
	"Ecommerce/rest/handlers"
	"Ecommerce/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /hellohandeler",
		http.HandlerFunc(
			handlers.Hellohandler),
	)

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middlewares.AuthenticationJWT,
		),
	)
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middlewares.AuthenticationJWT,
		),
	)

	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middlewares.AuthenticationJWT,
		),
	)

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)

}
