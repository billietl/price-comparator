package web

import "github.com/gorilla/mux"

func makeRouter() (router *mux.Router) {
	router = mux.NewRouter()
	setupProductRouter(router.PathPrefix("/product").Subrouter())
	return
}
