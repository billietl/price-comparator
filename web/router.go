package web

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

func NewRouter() (rtr *Router) {
	rtr = &Router{
		Router: mux.NewRouter(),
	}
	return
}

func (r Router) RegisterController(ctrl Controller, path string) (err error) {
	ctrl.SetupRouter(r.PathPrefix(path).Subrouter())
	return nil
}
