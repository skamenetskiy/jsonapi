package jsonapi

import (
	"encoding/json"
	"strconv"

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
		c.SetStatusCode(StatusInternalServerError)
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

// GetParamString returns path parameter k as string
func (c *Ctx) GetParamString(k string) string {
	return c.UserValue(k).(string)
}

// GetParamInt returns path parameter k as int or error
func (c *Ctx) GetParamInt(k string) (int, error) {
	return strconv.Atoi(c.GetParamString(k))
}

// GetParamInt64 returns path parameter k as int64 or error
func (c *Ctx) GetParamInt64(k string) (int64, error) {
	return strconv.ParseInt(c.GetParamString(k), 10, 0)
}

// GetParamUint64 returns path parameter k as uint64 or error
func (c *Ctx) GetParamUint64(k string) (uint64, error) {
	return strconv.ParseUint(c.GetParamString(k), 10, 0)
}

// GetParamFloat64 returns path parameter k as float64 or error
func (c *Ctx) GetParamFloat64(k string) (float64, error) {
	return strconv.ParseFloat(c.GetParamString(k), 0)
}

// OK is writing the response to response body with
// http status 200OK
func (c *Ctx) OK(v json.Marshaler) {
	c.SetStatusCode(StatusOK)
	c.WriteJSON(v)
}

// Err is writing an error to response body with code
func (c *Ctx) Err(err error, code int) {
	c.SetStatusCode(code)
	c.WriteJSON(Error{err.Error(), code})
}

// ErrBadRequest writes http error BadRequest to response body
func (c *Ctx) ErrBadRequest(err error)  {
	c.Err(err, StatusBadRequest)
}

// ErrUnauthorized writes http error Unauthorized to response body
func (c *Ctx) ErrUnauthorized(err error)  {
	c.Err(err, StatusUnauthorized)
}

// ErrForbidden writes http error Forbidden to response body
func (c *Ctx) ErrForbidden(err error)  {
	c.Err(err, StatusForbidden)
}

// ErrNotFound writes http error NotFound to response body
func (c *Ctx) ErrNotFound(err error)  {
	c.Err(err, StatusNotFound)
}

// ErrMethodNotAllowed writes http error MethodNotAllowed to response body
func (c *Ctx) ErrMethodNotAllowed(err error)  {
	c.Err(err, StatusMethodNotAllowed)
}

// ErrInternalServerError writes http error InternalServerError to response body
func (c *Ctx) ErrInternalServerError(err error)  {
	c.Err(err, StatusInternalServerError)
}

// ErrBadGateway writes http error BadGateway to response body
func (c *Ctx) ErrBadGateway(err error)  {
	c.Err(err, StatusBadGateway)
}

// ErrServiceUnavailable writes http error ServiceUnavailable to response body
func (c *Ctx) ErrServiceUnavailable(err error)  {
	c.Err(err, StatusServiceUnavailable)
}

// ErrGatewayTimeout writes http error GatewayTimeout to response body
func (c *Ctx) ErrGatewayTimeout(err error)  {
	c.Err(err, StatusGatewayTimeout)
}