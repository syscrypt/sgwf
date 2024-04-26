package sgwf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syscrypt/sgwf/log/basic"
)

type ResponseWrapper interface {
	Wrap(f func(w http.ResponseWriter, r *http.Request) *HttpResponse) func(w http.ResponseWriter, r *http.Request)
}

type DefaultResponseWrapper struct {
	*basic.LogImpl
}

func NewDefaultResponseWrapper() *DefaultResponseWrapper {
	return &DefaultResponseWrapper{
		LogImpl: basic.NewLog(),
	}
}

func (rw *DefaultResponseWrapper) logErr(err error) {
	if err == nil {
		return
	}

	if rw.FieldLogger != nil {
		rw.FieldLogger.Error(err)
		return
	}
	rw.Logger.Errorf(err.Error())
}

func (rw *DefaultResponseWrapper) writeJsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	if body == nil {
		w.WriteHeader(statusCode)
		return
	}

	var content []byte
	var err error

	if b, ok := body.([]byte); ok {
		content = b
	} else {
		content, err = json.Marshal(body)
		if err != nil {
			rw.logErr(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(statusCode)

	if ln, err := w.Write(content); err != nil {
		rw.logErr(err)
	} else if ln != len(content) {
		rw.logErr(fmt.Errorf("written body len of %d doesn't equal content len %d", ln, len(content)))
	}
}

func (rw *DefaultResponseWrapper) Wrap(f func(w http.ResponseWriter, r *http.Request) *HttpResponse) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := f(w, r)

		if resp == nil {
			return
		}

		rw.logErr(resp.Error)

		rw.writeJsonResponse(w, resp.StatusCode, resp.Body)
	}
}
