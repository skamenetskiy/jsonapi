package jsonapi

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

// Controller interface
type Controller interface {
	Methods() ControllerMethods // should return controller methods
}

// ControllerMethods is a map linking http method to controller paths
type ControllerMethods map[string]ControllerPaths

// ControllerPaths is a map linking path to handler
type ControllerPaths map[string]ControllerHandler

// BaseController is a helper struct that implements basic
// response methods that can be used by inheriting controller
type BaseController struct {
}

// OK returns 200 OK response
func (c *BaseController) OK(v json.Marshaler) *Result {
	return &Result{Data: v}
}

// Err returns an error response with http code
func (c *BaseController) Err(err error, code int) *Result {
	if err == nil {
		err = errors.New("unknown error")
	}
	return &Result{
		Err: &Error{err.Error(), code},
	}
}

// ErrNotFound returns http 404 error response
func (c *BaseController) ErrNotFound(err error) *Result {
	return c.Err(err, fasthttp.StatusNotFound)
}

// Result is an object returned from controller method
// If Err is not nil, an error will be returned to
// the client
type Result struct {
	Data json.Marshaler
	Err  *Error
}

// HasError returns TRUE if the result contains an error
func (r *Result) HasError() bool {
	return r.Err != nil
}

// Error sets an err to result and return the result
func (r *Result) Error(err *Error) *Result {
	r.Err = err
	return r
}

// ResultList
type ResultList []json.Marshaler

func (l ResultList) MarshalJSON() ([]byte, error) {
	res := [][]byte{[]byte("[")}
	for _, i := range l {
		b, err := i.MarshalJSON()
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}
	res = append(res, []byte("]"))
	return bytes.Join(res, []byte(",")), nil
}
