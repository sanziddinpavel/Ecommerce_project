package cmd

import (
	"Ecommerce/config"
	"Ecommerce/rest"
	"Ecommerce/rest/handlers/product"
	"Ecommerce/rest/handlers/user"
)

func Server() {
	cnf := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
