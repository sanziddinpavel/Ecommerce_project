package rest

import (
	"Ecommerce/config"

	"Ecommerce/rest/handlers/product"
	"Ecommerce/rest/handlers/user"
	"Ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (Server *Server) Start() {

	manager := middlewares.NewManager()
	// manager.Use()
	manager.Use(
		middlewares.PreflightReq,
		middlewares.CorsMiddleware,
		middlewares.Logger)

	mux := http.NewServeMux()

	WrappedMux := manager.WrapMux(mux)

	Server.productHandler.RegisterRoutes(mux, manager)
	Server.userHandler.RegisterRoutes(mux, manager)
	// InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(Server.cnf.HttpPort)
	fmt.Println("server running on port", addr)
	err := http.ListenAndServe(addr, WrappedMux)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}
