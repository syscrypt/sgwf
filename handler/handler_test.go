package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syscrypt/sgwf/log/logger_test"
)

type controller struct{}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestLogMiddleware(t *testing.T) {
	handler := New()

	lg := &logger_test.LoggerMock{}
	handler.SetLogger(lg)

	handler.SetupRoutes(map[string]Controller{
		"/test": &controller{},
	})

	rr := httptest.NewRecorder()

	assert.Equal(t, 0, len(lg.Buf))
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 1, len(lg.Buf))
	assert.Equal(t, "info: new http GET request /test", lg.Buf[0])
	lg.Buf = []string{}

	assert.Equal(t, 0, len(lg.Buf))
	req, _ = http.NewRequest(http.MethodPost, "/test?value=10", nil)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 1, len(lg.Buf))
	assert.Equal(t, "info: new http POST request /test?value=10", lg.Buf[0])
	lg.Buf = []string{}
}
