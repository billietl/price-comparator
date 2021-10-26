package pages

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexControllerGetIndex(t *testing.T) {
	ctrl := NewIndexController()
	// create mock http stuff
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	// run the handler
	ctrl.GetIndex(res, req)
	// check
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, "text/html; charset=utf-8", res.Result().Header.Get("content-type"))
}
