package web

import (
	"fmt"
	"log"
	"net/http"
	"price-comparator/dao"
)

type Server struct {
	port int
	router *Router
}

func MakeServer(port int, dao *dao.DAOBundle) (server *Server) {
	router := NewRouter()

	router.RegisterController(
		NewProductController(dao),
		"/api/v1/product",
	)

	return &Server{
		port: port,
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
