package jsonapi

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// NewClient generates a new jsonapi client
func NewClient(addr string) *Client {
	c := &Client{
		addr:     addr,
		authFunc: func(*Request) {},
		useSSL:   false,
	}
	return c
}

// Client describes jsonapi client
type Client struct {
	addr     string
	authFunc ClientAuthFunc
	useSSL   bool
}

// SetAuthFunc sets authentication modifier function
func (c *Client) SetAuthFunc(authFunc ClientAuthFunc) *Client {
	c.authFunc = authFunc
	return c
}

// UseSSL forces the client to use https protocol
func (c *Client) UseSSL() *Client {
	c.useSSL = true
	return c
}

// Request returns a new *Request object
func (c *Client) Request() *Request {
	r := new(Request)
	if c.useSSL {
		r.addr = "https://" + c.addr
	} else {
		r.addr = "http://" + c.addr
	}
	// call auth func
	c.authFunc(r)
	return r
}

func (c *Client) Get(uri string) (*Response, error) {
	return c.Request().
		SetMethod(MethodGet).
		SetURI(uri).
		Do()
}

func (c *Client) Post(uri string, body json.Marshaler) (*Response, error) {
	b, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return c.Request().
		SetMethod(MethodPost).
		SetURI(uri).
		SetBody(b).
		Do()
}

func (c *Client) Put(uri string, body json.Marshaler) (*Response, error) {
	b, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return c.Request().
		SetMethod(MethodPut).
		SetURI(uri).
		SetBody(b).
		Do()
}

func (c *Client) Delete(uri string) (*Response, error) {
	return c.Request().
		SetMethod(MethodDelete).
		SetURI(uri).
		Do()
}

// Request is the request object
type Request struct {
	addr    string            // request address, including protocol, host and port
	method  string            // request http method
	uri     string            // request uri or path
	body    []byte            // request body
	headers map[string]string // request headers
}

func (r *Request) SetMethod(method string) *Request {
	r.method = method
	return r
}

func (r *Request) SetURI(uri string) *Request {
	r.uri = uri
	return r
}

func (r *Request) SetBody(body []byte) *Request {
	r.body = body
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.headers = headers
	return r
}

func (r *Request) SetHeader(k string, v string) *Request {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[k] = v
	return r
}

func (r *Request) makeURL() string {
	return r.addr + r.uri
}

// Do executes the http request and returns *Response
// or error if the request failed
func (r *Request) Do() (*Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.method)
	req.SetRequestURI(r.makeURL())
	if r.body != nil {
		req.SetBody(r.body)
	}
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}
	return &Response{res}, nil
}

// Response is a wrapper for *fasthttp.Response
// it adds some useful methods for working with json
type Response struct {
	*fasthttp.Response
}

// ReadJSON reads json into v from request body
func (r *Response) ReadJSON(v json.Unmarshaler) error {
	return v.UnmarshalJSON(r.Body())
}
