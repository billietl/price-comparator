package web

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

type ProductController struct {
	Dao *dao.Bundle
}

func NewProductController(dao *dao.Bundle) *ProductController {
	return &ProductController{
		Dao: dao,
	}
}

func (ph ProductController) SetupRouter(router *mux.Router) {
	router.
		Methods(http.MethodGet).
		Path("/{id}").
		Name("Get a single product").
		HandlerFunc(ph.GetProductController)
	router.
		Methods(http.MethodPut).
		Path("/").
		Name("Create a single product").
		HandlerFunc(ph.CreateProductController)
	router.
		Methods(http.MethodDelete).
		Path("/{id}").
		Name("Delete a single product").
		HandlerFunc(ph.DeleteProductController)
	router.
		Methods(http.MethodPatch).
		Path("/{id}").
		Name("Update a single product").
		HandlerFunc(ph.UpdateProductController)
}

func (ph ProductController) CreateProductController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	product := model.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Could not decode request")
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.GenerateID()
	err = ph.Dao.ProductDAO.Upsert(r.Context(), &product)
	if err != nil {
		log.Printf("Could not upsert product")
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}
}

func (ph ProductController) GetProductController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	product, err := ph.Dao.ProductDAO.Load(r.Context(), id)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			log.Printf("Product not found : %s", id)
			log.Print(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Error fetching product %s", id)
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}

}

func (ph ProductController) DeleteProductController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	err := ph.Dao.ProductDAO.Delete(r.Context(), id)
	if err != nil {
		log.Printf("Error fetching product %s", id)
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ph ProductController) UpdateProductController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	if id == "" {
		log.Printf("No product ID found in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := model.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Could not decode request")
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID = id
	err = ph.Dao.ProductDAO.Upsert(r.Context(), &product)
	if err != nil {
		log.Printf("Could not upsert product")
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		log.Printf("Error writing response")
		log.Print(err.Error())
	}
}
