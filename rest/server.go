package rest

import (
	"Ecommerce/config"

	"Ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middlewares.NewManager()
	// manager.Use()
	manager.Use(
		middlewares.PreflightReq,
		middlewares.CorsMiddleware,
		middlewares.Logger)

	mux := http.NewServeMux()

	WrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server running on port", addr)
	err := http.ListenAndServe(addr, WrappedMux)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}
