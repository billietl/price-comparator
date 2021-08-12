package web

import (
	"fmt"
	"log"
	"net/http"
	"price-comparator/dao"
)

type Server struct {
	port int
	dao  *dao.DAOBundle
}

func MakeServer(port int, dao *dao.DAOBundle) (server *Server) {
	return &Server{
		port: port,
		dao:  dao,
	}
}

func (serv Server) Run() error {

	router := makeRouter()

	log.Print(fmt.Sprintf("Starting server on port %d", serv.port))
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", serv.port),
			router,
		),
	)
	log.Print("Shuting server")

	return nil
}
