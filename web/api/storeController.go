package api

import (
	"encoding/json"
	"log"
	"net/http"
	"price-comparator/dao"
	"price-comparator/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
		Path("/{id}").
		Name("Get a single store").
		HandlerFunc(sc.GetStoreController)
	router.
		Methods(http.MethodPut).
		Path("/").
		Name("Create a single store").
		HandlerFunc(sc.CreateStoreController)
	router.
		Methods(http.MethodDelete).
		Path("/{id}").
		Name("Delete a single store").
		HandlerFunc(sc.DeleteStoreController)
	router.
		Methods(http.MethodPatch).
		Path("/{id}").
		Name("Update a single store").
		HandlerFunc(sc.UpdateStoreController)
}

func (sc StoreController) CreateStoreController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	store := model.Store{}

	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		log.Printf("Could not decode request")
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store.GenerateID()
	err = sc.Dao.StoreDAO.Upsert(r.Context(), &store)
	if err != nil {
		log.Printf("Could not upsert store")
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&store)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}
}

func (sc StoreController) GetStoreController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	store, err := sc.Dao.StoreDAO.Load(r.Context(), id)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			log.Printf("Store not found : %s", id)
			log.Print(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Error fetching store %s", id)
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&store)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}
}

func (sc StoreController) UpdateStoreController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	if id == "" {
		log.Printf("No store ID found in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store := model.Store{}

	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		log.Printf("Could not decode request")
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store.ID = id
	err = sc.Dao.StoreDAO.Upsert(r.Context(), &store)
	if err != nil {
		log.Printf("Could not upsert store")
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&store)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}
}

func (sc StoreController) DeleteStoreController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	err := sc.Dao.StoreDAO.Delete(r.Context(), id)
	if err != nil {
		log.Printf("Error deleting store %s", id)
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
