package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupProductRouter(router *mux.Router) {
	router.HandleFunc("/{id}", getProductHandler)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
