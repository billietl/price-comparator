package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"price-comparator/dao"
	"price-comparator/model"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestStoreControllerCreateStoreController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create a mock Store
	reqStore := model.NewStore("Auchan", "Paris", "75000")
	var resStore model.Store
	// create mock http stuff
	body, _ := json.Marshal(reqStore)
	req := httptest.NewRequest(http.MethodPut, "/store/", bytes.NewReader(body))
	res := httptest.NewRecorder()
	// run the handler
	ctrl.CreateStoreController(res, req)
	err = json.NewDecoder(res.Body).Decode(&resStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqStore.ValueEquals(&resStore))
}

func TestStoreControllerGetStoreController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create a mock Store
	reqStore := model.NewStore("Carrefour", "Annemasse", "74100")
	var resStore model.Store
	// persist the mock Store
	dao.StoreDAO.Upsert(context.Background(), reqStore)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", reqStore.ID), nil)
	res := httptest.NewRecorder()
	// run the handler
	router := mux.NewRouter()
	router.HandleFunc("/{id}", ctrl.GetStoreController)
	router.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&resStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqStore.ValueEquals(&resStore))
}

func TestStoreControllerUpdateStoreController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create a mock Store
	reqStore := model.NewStore("Leclerc", "Villeneuve d'Ascq", "59491")
	var resStore model.Store
	// persist the mock Store
	dao.StoreDAO.Upsert(context.Background(), reqStore)
	// modify the mock Store
	reqStore.Name = "barbaz"
	reqStore.City = "Croix"
	reqStore.Zipcode = "59170"
	// create mock http stuff
	body, _ := json.Marshal(reqStore)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s", reqStore.ID), bytes.NewReader(body))
	res := httptest.NewRecorder()
	// run the handler
	router := mux.NewRouter()
	router.HandleFunc("/{id}", ctrl.UpdateStoreController)
	router.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&resStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqStore.ValueEquals(&resStore))
}

func TestStoreControllerDeleteStoreController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create a mock Store
	reqStore := model.NewStore("Intermarché", "Brest", "29200")
	// persist the mock Store
	dao.StoreDAO.Upsert(context.Background(), reqStore)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%s", reqStore.ID), nil)
	res := httptest.NewRecorder()
	// run the handler
	router := mux.NewRouter()
	router.HandleFunc("/{id}", ctrl.DeleteStoreController)
	router.ServeHTTP(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusOK)
}
