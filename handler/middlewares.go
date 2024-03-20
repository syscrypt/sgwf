package handler

import (
	"net/http"

	"github.com/syscrypt/sgwf/log/logger"
)

func (h *Handler) LogMiddleware(hl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.FieldLogger != nil {
			h.FieldLogger.WithFields(logger.Fields{
				"url":    r.URL.String(),
				"method": r.Method,
			}).Infoln("new http request")
		} else {
			h.Logger.Infof("new http %s request %s", r.Method, r.URL.String())
		}

		hl.ServeHTTP(w, r)
	})
}
