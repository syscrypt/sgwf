package handler

import (
	"bytes"
	"net/http"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status      int
	WroteHeader bool
	Body        []byte
}

func NewStatusRecorder(w http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: w,
	}
}

func (r *StatusRecorder) WriteHeader(code int) {
	if r.WroteHeader {
		return
	}
	r.Status = code
	r.WroteHeader = true
}

func (r *StatusRecorder) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(r.Body)
	return buf.Write(b)
}
