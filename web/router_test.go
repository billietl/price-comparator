package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
)

func TestRegisterController(t *testing.T) {
	router := NewRouter()

	path := randomstring.HumanFriendlyString(10)
	ctrl := &TestController{}

	router.RegisterController(
		ctrl,
		fmt.Sprintf("/%s", path),
	)

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/%s/test", path),
		nil,
	)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "ok", rr.Body.String())
}
