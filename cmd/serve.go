package cmd

import (
	"Ecommerce/config"
	"Ecommerce/repo"
	"Ecommerce/rest"
	"Ecommerce/rest/handlers/product"
	"Ecommerce/rest/handlers/user"
	"Ecommerce/rest/middlewares"
)

func Server() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddleware(cnf)

	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(middlewares, productRepo)

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
