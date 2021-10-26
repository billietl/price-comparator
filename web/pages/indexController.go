package pages

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (ic IndexController) SetupRouter(router *mux.Router) {
	router.
		Methods(http.MethodGet).
		Path("/").
		Name("Index page").
		HandlerFunc(ic.GetIndex)
}

func (ic IndexController) GetIndex(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "index", nil)
}
