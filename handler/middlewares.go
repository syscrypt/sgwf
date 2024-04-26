package handler

import (
	"net/http"

	"github.com/syscrypt/sgwf/extra"
	"github.com/syscrypt/sgwf/log/logger"
)

func (h *Handler) LogMiddleware(hl http.Handler) http.Handler {
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

		hl.ServeHTTP(w, r)
	})
}
