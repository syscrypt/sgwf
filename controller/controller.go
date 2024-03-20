package controller

import (
	"net/http"

	"github.com/syscrypt/sgwf"
	"github.com/syscrypt/sgwf/log/basic"
)

type Controller struct {
	*basic.LogImpl
	Router *sgwf.Router
}

func NewBaseController() *Controller {
	r := sgwf.NewRouter(sgwf.NewDefaultResponseWrapper())

	return (&Controller{
		Router:  r,
		LogImpl: basic.NewLog(),
	})
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Router.ServeHTTP(w, r)
}
