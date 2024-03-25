package sgwf

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

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

func PackRequestPathVariables(r *http.Request, dest interface{}) error {
	values := mux.Vars(r)

	if values == nil {
		return ErrValueNotFound
	}

	if reflect.TypeOf(dest).Kind() != reflect.Pointer {
		return ErrNonPointerType
	}

	types, err := getStructTypes(reflect.ValueOf(dest).Elem().Interface())
	if err != nil {
		return err
	}

	result := []string{}
	for k, v := range values {
		if types[k] != "string" {
			result = append(result, fmt.Sprintf("\"%s\": %s", k, v))
		} else {
			result = append(result, fmt.Sprintf("\"%s\": \"%s\"", k, v))
		}
	}

	resultStr := strings.Join(result, ",")
	resultStr = fmt.Sprintf("{%s}", resultStr)

	return json.Unmarshal([]byte(resultStr), dest)
}

func getStructTypes(i interface{}) (map[string]string, error) {
	val := reflect.ValueOf(i)
	typ := val.Type()
	ret := make(map[string]string)

	if val.Kind() != reflect.Struct {
		return nil, ErrNoStructParam
	}

	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldJsonName := fieldType.Tag.Get("json")
		if fieldJsonName == "" || len(strings.Split(fieldJsonName, ",")) > 1 {
			return nil, ErrWrongJsonTag
		}

		ret[fieldJsonName] = fieldType.Type.Name()
	}

	return ret, nil
}
