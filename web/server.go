package web

import (
	"fmt"
	"log"
	"net/http"
	"price-comparator/dao"
	"price-comparator/web/api"
	"price-comparator/web/pages"
)

type Server struct {
	port   int
	router *Router
}

func MakeServer(port int, dao *dao.Bundle) (server *Server) {
	router := NewRouter()

	router.RegisterController(
		api.NewProductController(dao),
		"/api/v1/product",
	)

	router.RegisterController(
		api.NewStoreController(dao),
		"/api/v1/store",
	)

	router.RegisterController(
		pages.NewStoreController(dao),
		"/stores/",
	)

	router.RegisterController(
		pages.NewIndexController(),
		"/",
	)

	return &Server{
		port:   port,
		router: router,
	}
}

func (serv Server) Run() error {

	log.Print(fmt.Sprintf("Starting server on port %d", serv.port))
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", serv.port),
			serv.router,
		),
	)
	log.Print("Shuting server")

	return nil
}
