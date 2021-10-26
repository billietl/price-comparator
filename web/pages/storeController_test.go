package pages

import (
	"context"
	"net/http"
	"net/http/httptest"
	"price-comparator/dao"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreControllerGetStoreList(t *testing.T) {
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	// run the handler
	ctrl.GetStoreList(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "text/html; charset=utf-8", res.Result().Header.Get("content-type"))
}

func TestStoreControllerAddStoreForm(t *testing.T) {
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	// run the handler
	ctrl.AddStoreForm(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "text/html; charset=utf-8", res.Result().Header.Get("content-type"))
}

func TestStoreControllerAddStoreAction(t *testing.T) {
	dao, err := dao.GetBundle(context.Background(), "firestore")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ctrl := NewStoreController(dao)
	// create mock http stuff
	req := httptest.NewRequest(http.MethodPost, "/addAction", strings.NewReader("name=foo&city=bar&zipcode=00000"))
	res := httptest.NewRecorder()
	// run the handler
	ctrl.AddStoreAction(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusFound)
	assert.Equal(t, ".", res.Result().Header.Get("location"))
	assert.Equal(t, "text/html; charset=utf-8", res.Result().Header.Get("content-type"))
}
