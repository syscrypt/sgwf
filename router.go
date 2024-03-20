package sgwf

import (
	h "net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	responseWrapper ResponseWrapper
}

func NewRouter(wrapper ResponseWrapper) *Router {
	return &Router{
		Router:          mux.NewRouter(),
		responseWrapper: wrapper,
	}
}

func (r *Router) HandleFunc(path string, f func(w h.ResponseWriter, req *h.Request) *HttpResponse) *mux.Route {
	return r.Router.HandleFunc(path, r.responseWrapper.Wrap(f))
}
