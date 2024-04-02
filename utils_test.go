package sgwf

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type dest struct {
	Param1 int    `json:"param1"`
	Param2 bool   `json:"param2"`
	Param3 string `json:"param3"`
}

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

func TestRequestVarsArrayList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test/1234", nil)
	req = mux.SetURLVars(req, map[string]string{
		"param1": "1234",
	})

	assert.EqualValues(t, map[string][]string{
		"param1": []string{"1234"},
	}, RequestVarsArrayList(req))
}
