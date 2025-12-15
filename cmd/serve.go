package cmd

import (
	"Ecommerce/config"
	"Ecommerce/infra/db"
	"Ecommerce/repo"
	"Ecommerce/rest"
	"Ecommerce/rest/handlers/product"
	"Ecommerce/rest/handlers/user"
	"Ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Server() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddleware(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	productHandler := product.NewHandler(middlewares, productRepo)

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
