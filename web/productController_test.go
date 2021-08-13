package web

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"price-comparator/dao"
	"price-comparator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductHandler(t *testing.T) {
	// init controller
	dao, err := dao.GetDAOBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewProductController(dao)
	// create a mock product
	reqProduct := model.NewProduct("foobar", true)
	var resProduct model.Product
	// create mock http stuff
	body, _ := json.Marshal(reqProduct)
	req := httptest.NewRequest(http.MethodPut, "/product/", bytes.NewReader(body))
	res := httptest.NewRecorder()
	// run the handler
	ctrl.CreateProductHandler(res, req)
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
