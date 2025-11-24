package cmd

import (
	"Ecommerce/config"
	"Ecommerce/rest"
	"Ecommerce/rest/handlers/product"
	"Ecommerce/rest/handlers/user"
	"Ecommerce/rest/middlewares"
)

func Server() {
	cnf := config.GetConfig()
	middlewares := middlewares.NewMiddleware(cnf)
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
