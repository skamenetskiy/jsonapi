package jsonapi

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// Ctx is a wrapper for fasthttp.RequestCtx
type Ctx struct {
	*fasthttp.RequestCtx
}

// ReadJSON will try to read request body into v
func (c *Ctx) ReadJSON(v json.Unmarshaler) error {
	return v.UnmarshalJSON(c.PostBody())
}

// WriteJSON will try to write v to response body
// or will output default error if marshal fails
func (c *Ctx) WriteJSON(v json.Marshaler) {
	b, err := v.MarshalJSON()
	if err != nil {
		// output default marshal error
		c.SetStatusCode(fasthttp.StatusInternalServerError)
		c.SetBody([]byte(`{"error":"failed to marshal json"}`))
	} else {
		// output json
		c.SetBody(b)
	}
}

// SetHeader sets a response header
func (c *Ctx) SetHeader(k string, v string) {
	c.Response.Header.Set(k, v)
}

// GetHeader returns a request header
func (c *Ctx) GetHeader(k string) string {
	return string(c.Request.Header.Peek(k))
}

func (c *Ctx) OK(v json.Marshaler) {
	c.SetStatusCode(fasthttp.StatusOK)
	c.WriteJSON(v)
}

func (c *Ctx) ErrNotFound(err error) {
	c.Err(err, fasthttp.StatusNotFound)
}

func (c *Ctx) ErrMethodNotAllowed(err error) {
	c.Err(err, fasthttp.StatusMethodNotAllowed)
}

func (c *Ctx) ErrInternalServerError(err error) {
	c.Err(err, fasthttp.StatusInternalServerError)
}

func (c *Ctx) ErrUnauthorized(err error) {
	c.Err(err, fasthttp.StatusUnauthorized)
}

func (c *Ctx) Err(err error, code int) {
	c.SetStatusCode(code)
	c.WriteJSON(Error{err.Error(), code})
}
