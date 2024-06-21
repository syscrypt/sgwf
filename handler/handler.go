package handler

import (
	"net/http"
	"net/url"

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
		h.Router.PathPrefix(path).Handler(http.StripPrefix(path, h.HandleEmptyPath(controller)))
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}

func (h *Handler) HandleEmptyPath(ho http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		if r.URL.Path == "" {
			r2.URL.Path = "/"
		}
		if r.URL.RawPath == "" {
			r2.URL.RawPath = "/"
		}
		ho.ServeHTTP(w, r2)
	})
}
