package web

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

	"github.com/stretchr/testify/assert"
)

func TestCreateProductController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewProductController(dao)
	// create a mock product
	reqProduct := model.NewProduct("foobar", true, false)
	var resProduct model.Product
	// create mock http stuff
	body, _ := json.Marshal(reqProduct)
	req := httptest.NewRequest(http.MethodPut, "/product/", bytes.NewReader(body))
	res := httptest.NewRecorder()
	// run the handler
	ctrl.CreateProductController(res, req)
	err = json.NewDecoder(res.Body).Decode(&resProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqProduct.ValueEquals(&resProduct))
}

func TestGetProductController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewProductController(dao)
	// create a mock product
	reqProduct := model.NewProduct("foobar", true, false)
	var resProduct model.Product
	// persist the mock product
	dao.ProductDAO.Upsert(context.Background(), reqProduct)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", reqProduct.ID), nil)
	res := httptest.NewRecorder()
	// run the handler
	router := NewRouter()
	router.HandleFunc("/{id}", ctrl.GetProductController)
	router.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&resProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqProduct.ValueEquals(&resProduct))
}

func TestDeleteProductController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewProductController(dao)
	// create a mock product
	reqProduct := model.NewProduct("foobar", true, false)
	// persist the mock product
	dao.ProductDAO.Upsert(context.Background(), reqProduct)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%s", reqProduct.ID), nil)
	res := httptest.NewRecorder()
	// run the handler
	router := NewRouter()
	router.HandleFunc("/{id}", ctrl.DeleteProductController)
	router.ServeHTTP(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusOK)
}

func TestUpdateProductController(t *testing.T) {
	// init controller
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewProductController(dao)
	// create a mock product
	reqProduct := model.NewProduct("foobar", true, false)
	var resProduct model.Product
	// persist the mock product
	dao.ProductDAO.Upsert(context.Background(), reqProduct)
	// modify the mock product
	reqProduct.Name = "barbaz"
	reqProduct.Bio = false
	// create mock http stuff
	body, _ := json.Marshal(reqProduct)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s", reqProduct.ID), bytes.NewReader(body))
	res := httptest.NewRecorder()
	// run the handler
	router := NewRouter()
	router.HandleFunc("/{id}", ctrl.UpdateProductController)
	router.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&resProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
	assert.Equal(t, true, reqProduct.ValueEquals(&resProduct))
}
