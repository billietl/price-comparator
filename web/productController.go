package web

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/mux"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
	"price-comparator/dao"
	"price-comparator/model"
)

type ProductHandler struct {
	Dao *dao.DAOBundle
}

func NewProductController(dao *dao.DAOBundle) *ProductHandler{
	return &ProductHandler{
		Dao: dao,
	}
}

func (ph ProductHandler) SetupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/{id}").
		Name("Get a single product").
		HandlerFunc(ph.GetProductHandler)
	router.
		Methods("PUT").
		Path("/").
		Name("Create a single product").
		HandlerFunc(ph.CreateProductHandler)
}

func (ph ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	product := model.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Could not decode request")
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.GenerateID()
	err = ph.Dao.ProductDAO.Upsert(r.Context(), &product)
	if err != nil {
		log.Printf("Could not upsert product")
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

func (ph ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	product, err := ph.Dao.ProductDAO.Load(r.Context(), id)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			log.Printf(fmt.Sprintf("Product not found : %s", id))
			log.Printf(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf(fmt.Sprintf("Error fetching product %s", id))
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}
