package sgwf

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrValueNotFound = errors.New("couldn't find any path variables")
)

func CheckVars(r *http.Request, vars ...string) error {
	values := mux.Vars(r)

	if values == nil {
		return ErrValueNotFound
	}

	for _, v := range vars {
		if values[v] == "" {
			return fmt.Errorf("couldn't find path variable %s in request path", v)
		}
	}

	return nil
}
