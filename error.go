package jsonapi

import (
	"fmt"
)

// Error is a custom error object
//go:generate easyjson
//easyjson:json
type Error struct {
	Err  string `json:"error"`
	Code int    `json:"code"`
}

// Error implements error interface
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Err)
}

// NewError returns a new *Error object
func NewError(err error, code ...int) *Error {
	e := new(Error)
	e.Err = err.Error()
	if len(code) > 0 {
		e.Code = code[0]
	}
	return e
}
