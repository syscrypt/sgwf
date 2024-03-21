package sgwf

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCheckVarsValidParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test/1234", nil)
	req = mux.SetURLVars(req, map[string]string{
		"param": "1234",
	})

	assert.NoError(t, CheckVars(req, "param"))
}

func TestCheckVarsInvalidParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test/1234", nil)
	req = mux.SetURLVars(req, map[string]string{
		"param1": "1234",
	})

	err := CheckVars(req, "param1", "param2")
	assert.Equal(t, "couldn't find path variable param2 in request path", err.Error())
}
