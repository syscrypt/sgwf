package sgwf

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syscrypt/sgwf/log/logger_test"
)

var (
	nilRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return nil
	}

	acceptedRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return &HttpResponse{
			StatusCode: http.StatusAccepted,
			Body:       nil,
			Error:      nil,
		}
	}

	bodyRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return &HttpResponse{
			StatusCode: http.StatusOK,
			Body: map[string]interface{}{
				"test1": "abc",
				"test2": 1234,
			},
			Error: nil,
		}
	}

	errNilBodyRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return &HttpResponse{
			StatusCode: http.StatusOK,
			Body:       nil,
			Error:      errors.New("error test"),
		}
	}

	errorBodyRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return &HttpResponse{
			StatusCode: http.StatusOK,
			Body: map[string]interface{}{
				"message": "this error is propagated to the client",
			},
			Error: errors.New("hidden error"),
		}
	}

	byteBodyRespFunc = func(w http.ResponseWriter, r *http.Request) *HttpResponse {
		return &HttpResponse{
			StatusCode: http.StatusOK,
			Body:       []byte{0xca, 0xfe, 0x13, 0x37},
		}
	}
)

func TestNilResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(nilRespFunc)
	f(rr, req)

	assert.Equal(t, 0, len(lg.Buf))
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}

func TestAcceptedResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(acceptedRespFunc)
	f(rr, req)

	assert.Equal(t, 0, len(lg.Buf))
	assert.Equal(t, rr.Result().StatusCode, http.StatusAccepted)
}

func TestOkBodyResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(bodyRespFunc)
	f(rr, req)

	defer rr.Result().Body.Close()
	body, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(lg.Buf))
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.EqualValues(t, "{\"test1\":\"abc\",\"test2\":1234}", string(body))
}

func TestErrorNilBodyResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(errNilBodyRespFunc)
	f(rr, req)

	defer rr.Result().Body.Close()
	body, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(lg.Buf))
	assert.Equal(t, "error: error test", lg.Buf[0])
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, string(body), "")
}

func TestErrorBodyResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(errorBodyRespFunc)
	f(rr, req)

	defer rr.Result().Body.Close()
	body, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(lg.Buf))
	assert.Equal(t, "error: hidden error", lg.Buf[0])
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, string(body), "{\"message\":\"this error is propagated to the client\"}")
}

func TestByteBodyResponse(t *testing.T) {
	responseWrapper := NewDefaultResponseWrapper()
	lg := &logger_test.LoggerMock{}
	responseWrapper.SetLogger(lg)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	f := responseWrapper.Wrap(byteBodyRespFunc)
	f(rr, req)

	defer rr.Result().Body.Close()
	body, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(lg.Buf))
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.EqualValues(t, body, []byte{0xca, 0xfe, 0x13, 0x37})
}
