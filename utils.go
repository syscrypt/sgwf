package sgwf

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrValueNotFound  = errors.New("couldn't find any path variables")
	ErrNoStructParam  = errors.New("provided parameter is not of type struct")
	ErrWrongJsonTag   = errors.New("provided struct properties must have single name json tags")
	ErrNonPointerType = errors.New("parameter must be a pointer")
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

func RequestVarsArrayList(r *http.Request) map[string][]string {
	vars := mux.Vars(r)
	ret := make(map[string][]string, len(vars))
	for k, v := range vars {
		ret[k] = []string{v}
	}
	return ret
}
