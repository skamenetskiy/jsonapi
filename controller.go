package jsonapi

import (
	"bytes"
	"encoding/json"
	"errors"

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

// Bytes is a method for returning []byte result
// (for example: raw json as []byte)
func (c *BaseController) OKBytes(v []byte) *Result {
	return &Result{Data: BytesResult(v)}
}

// String is a method for returning string result
// (for example: raw json as string)
func (c *BaseController) OKString(v string) *Result {
	return &Result{Data: StringResult(v)}
}

// List is a method for returning list of []json.Marshaler
// without the need to create a custom struct wrapper
func (c *BaseController) OKList(v []json.Marshaler) *Result {
	return &Result{Data: ListResult(v)}
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

// ErrBadRequest return http error BadRequest
func (c *BaseController) ErrBadRequest(err error) *Result {
	return c.Err(err, StatusBadRequest)
}

// ErrUnauthorized return http error Unauthorized
func (c *BaseController) ErrUnauthorized(err error) *Result {
	return c.Err(err, StatusUnauthorized)
}

// ErrForbidden return http error Forbidden
func (c *BaseController) ErrForbidden(err error) *Result {
	return c.Err(err, StatusForbidden)
}

// ErrNotFound returns http error NotFound
func (c *BaseController) ErrNotFound(err error) *Result {
	return c.Err(err, StatusNotFound)
}

// ErrMethodNotAllowed return http error MethodNotAllowed
func (c *BaseController) ErrMethodNotAllowed(err error) *Result {
	return c.Err(err, StatusMethodNotAllowed)
}

// ErrInternalServerError return http error InternalServerError
func (c *BaseController) ErrInternalServerError(err error) *Result {
	return c.Err(err, StatusInternalServerError)
}

// ErrBadGateway return http error BadGateway
func (c *BaseController) ErrBadGateway(err error) *Result {
	return c.Err(err, StatusBadGateway)
}

// ErrServiceUnavailable return http error ServiceUnavailable
func (c *BaseController) ErrServiceUnavailable(err error) *Result {
	return c.Err(err, StatusServiceUnavailable)
}

// ErrGatewayTimeout return http error GatewayTimeout
func (c *BaseController) ErrGatewayTimeout(err error) *Result {
	return c.Err(err, StatusGatewayTimeout)
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

// BytesResult is a wrapper for bytes to return it as json.Marshaler
type BytesResult []byte

// MarshalJSON implements json.Marshaler
func (b BytesResult) MarshalJSON() ([]byte, error) {
	return b, nil
}

// StringResult is a wrapper for string to return it json.Marshaler
type StringResult string

// MarshalJSON implements json.Marshaler
func (s StringResult) MarshalJSON() ([]byte, error) {
	return []byte(s), nil
}

// ListResult is a wrapper for []json.Marshaler to return a list
// without declaring a custom struct for the list
type ListResult []json.Marshaler

// MarshalJSON implements json.Marshaler
func (l ListResult) MarshalJSON() ([]byte, error) {
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
