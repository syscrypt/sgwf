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

func TestStructTypesReturnsNilOnWrongType(t *testing.T) {
	_, err := getStructTypes(10)
	assert.ErrorIs(t, ErrNoStructParam, err)
}

func TestStructTypesInference(t *testing.T) {
	d, err := getStructTypes(*&dest{})
	assert.NoError(t, err)
	assert.EqualValues(t, map[string]string{
		"param1": "int",
		"param2": "bool",
		"param3": "string",
	}, d)
}

func TestStructTypesInferenceReturnsErrOnNoJsonTags(t *testing.T) {
	_, err := getStructTypes(struct {
		Param1 int
	}{})
	assert.ErrorIs(t, err, ErrWrongJsonTag)
}

func TestStructTypesInferenceReturnsErrOnMalformedJsonTag(t *testing.T) {
	_, err := getStructTypes(struct {
		Param1 int `json:"test,omitempty"`
	}{})
	assert.ErrorIs(t, err, ErrWrongJsonTag)
}

func TestPackRequestPathVariables(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test/1234", nil)
	req = mux.SetURLVars(req, map[string]string{
		"param1": "1234",
		"param2": "false",
		"param3": "test",
	})

	d := &dest{}
	err := PackRequestPathVariables(req, d)
	assert.NoError(t, err)
}

func TestPackRequestPathVariablesFailsWithoutPointer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test/1234", nil)
	req = mux.SetURLVars(req, map[string]string{
		"param1": "1234",
		"param2": "false",
		"param3": "test",
	})

	d := dest{}
	err := PackRequestPathVariables(req, d)
	assert.ErrorIs(t, err, ErrNonPointerType)
}
