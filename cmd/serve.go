package cmd

import (
	"Ecommerce/config"
	"Ecommerce/rest"
)

func Server() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
