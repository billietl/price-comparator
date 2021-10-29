package pages

import (
	"fmt"
	"net/http"
	"price-comparator/dao"
	"price-comparator/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type StoreController struct {
	Dao *dao.Bundle
}

func NewStoreController(dao *dao.Bundle) *StoreController {
	return &StoreController{
		Dao: dao,
	}
}

func (sc StoreController) SetupRouter(router *mux.Router) {
	router.
		Methods(http.MethodGet).
		Path("/").
		Name("Store list page").
		HandlerFunc(sc.GetStoreList)
	router.
		Methods(http.MethodGet).
		Path("/add").
		Name("Add store form").
		HandlerFunc(sc.AddStoreForm)
	router.
		Methods(http.MethodPost).
		Path("/addAction").
		Name("Add store form action").
		HandlerFunc(sc.AddStoreAction)
}

func (sc StoreController) GetStoreList(w http.ResponseWriter, r *http.Request) {
	paginator := &dao.Paginator{
		PageNumber: 0,
		PageSize:   5,
	}
	stores, err := sc.Dao.StoreDAO.List(r.Context(), paginator)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	rnd.HTML(w, http.StatusOK, "store/List", map[string]interface{}{
		"stores":    stores,
		"paginator": paginator,
	})
}

func (sc StoreController) AddStoreForm(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "store/AddForm", nil)
}

func (sc StoreController) AddStoreAction(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	store := model.Store{}
	err = decoder.Decode(&store, r.PostForm)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	store.GenerateID()
	err = sc.Dao.StoreDAO.Upsert(r.Context(), &store)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("location", ".")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusFound)
}
