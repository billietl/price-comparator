package web

import "github.com/gorilla/mux"

type Controller interface {
	SetupRouter(router *mux.Router)
}
