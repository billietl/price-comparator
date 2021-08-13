package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TestController struct{}

func (tc TestController) SetupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/test").
		Name("test controller").
		HandlerFunc(tc.FooHandler)
}

func (tc TestController) FooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
