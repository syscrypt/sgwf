package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/syscrypt/sgwf/log/basic"
)

type Controller interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	*basic.LogImpl
	Router *mux.Router
}

func New() *Handler {
	r := mux.NewRouter()

	h := &Handler{
		Router:  r,
		LogImpl: basic.NewLog(),
	}

	h.Router.Use(h.LogMiddleware)

	return h
}

func (h *Handler) SetupRoutes(c map[string]Controller) {
	for path, controller := range c {
		h.Router.PathPrefix(path).Handler(http.StripPrefix(path, controller))
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}
