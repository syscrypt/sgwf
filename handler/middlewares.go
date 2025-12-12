package handler

import (
	"encoding/json"
	"net/http"

	"github.com/syscrypt/sgwf"
	"github.com/syscrypt/sgwf/extra"
	"github.com/syscrypt/sgwf/log/logger"
)

func (h *Handler) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if extra.IsNil(h.FieldLogger) {
			h.Logger.Infof("new http %s request %s from remote addr %s", r.Method, r.URL.String(), r.RemoteAddr)
		} else {
			h.FieldLogger.WithFields(logger.Fields{
				"url":    r.URL.String(),
				"method": r.Method,
				"remote": r.RemoteAddr,
			}).Infoln("new http request")
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) StatusRecorder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := NewStatusRecorder(w)
		next.ServeHTTP(recorder, r)

		if !recorder.WroteHeader {
			return
		}

		if isValidJSON(recorder.Body) {
			w.WriteHeader(recorder.Status)
			if recorder.Body != nil {
				w.Write(recorder.Body)
			}
			return
		}

		var resp []byte
		if recorder.Status >= 400 && recorder.Status < 600 {
			resp, _ = json.Marshal(&sgwf.Message{
				Error: string(recorder.Body),
			})
		} else {
			resp, _ = json.Marshal(&sgwf.Message{
				Message: string(recorder.Body),
			})
		}

		w.WriteHeader(recorder.Status)
		w.Write(resp)
	})
}

func isValidJSON(body []byte) bool {
	if body == nil {
		return true
	}

	var js json.RawMessage
	return json.Unmarshal(body, &js) == nil
}
