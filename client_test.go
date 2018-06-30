package jsonapi

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

type ClientTestSuite struct {
	suite.Suite
}

func (t *ClientTestSuite) TestNewClient() {
	c := NewClient(":123")
	t.IsType(&Client{}, c)
	t.Equal(":123", c.addr)
	t.NotNil(c.authFunc)
	t.False(c.useSSL)
}

func (t *ClientTestSuite) TestSetAuthFunc() {
	c := NewClient(":123")
	t.NotNil(c.authFunc)
	c.SetAuthFunc(nil)
	t.Nil(c.authFunc)
	c.SetAuthFunc(func(*Request) {})
	t.NotNil(c.authFunc)
}

func (t *ClientTestSuite) TestUseSSL() {
	c := NewClient("local:123")
	t.False(c.useSSL)
	c.UseSSL()
	t.True(c.useSSL)
}

func (t *ClientTestSuite) TestGet() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"some":"data"}`))
	}))
	r, err := NewClient(s.URL[7:]).Get("")
	t.NoError(err)
	t.NotNil(r)
	t.Equal([]byte(`{"some":"data"}`), r.Body())
}

func (t *ClientTestSuite) TestPost() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(r.Method, MethodPost)
		b,_ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		t.Equal(b, []byte(`{"error":"some error"}`))
		w.Write([]byte(`{"some":"data"}`))
	}))
	r, err := NewClient(s.URL[7:]).Post("", &Error{Err:"some error"})
	t.NoError(err)
	t.NotNil(r)
	t.Equal([]byte(`{"some":"data"}`), r.Body())
}

func (t *ClientTestSuite) TestPut() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(r.Method, MethodPut)
		b,_ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		t.Equal(b, []byte(`{"error":"some error"}`))
		w.Write([]byte(`{"some":"data"}`))
	}))
	r, err := NewClient(s.URL[7:]).Put("", &Error{Err:"some error"})
	t.NoError(err)
	t.NotNil(r)
	t.Equal([]byte(`{"some":"data"}`), r.Body())
}

func (t *ClientTestSuite) TestDelete() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(r.Method, MethodDelete)
		w.Write([]byte(`{"some":"data"}`))
	}))
	r, err := NewClient(s.URL[7:]).Delete("")
	t.NoError(err)
	t.NotNil(r)
	t.Equal([]byte(`{"some":"data"}`), r.Body())
}

func (t *ClientTestSuite) TestRequest() {
	r := NewClient("local:123").Request()
	t.NotNil(r)
	t.IsType(&Request{}, r)
	t.Equal("http://local:123", r.addr)
	r2 := NewClient("local:123").UseSSL().Request()
	t.NotNil(r2)
	t.IsType(&Request{}, r2)
	t.Equal("https://local:123", r2.addr)
}

func (t *ClientTestSuite) TestRequestSetMethod() {
	r := NewClient("local:123").Request()
	t.Empty(r.method)
	r.SetMethod(MethodGet)
	t.Equal(MethodGet, r.method)
	r.SetMethod(MethodPost)
	t.Equal(MethodPost, r.method)
	r.SetMethod(MethodPut)
	t.Equal(MethodPut, r.method)
	r.SetMethod(MethodConnect)
	t.Equal(MethodConnect, r.method)
	r.SetMethod(MethodHead)
	t.Equal(MethodHead, r.method)
	r.SetMethod(MethodOptions)
	t.Equal(MethodOptions, r.method)
	r.SetMethod(MethodPatch)
	t.Equal(MethodPatch, r.method)
	r.SetMethod(MethodTrace)
	t.Equal(MethodTrace, r.method)
}

func (t *ClientTestSuite) TestRequestSetURI() {
	r := NewClient("local:123").Request()
	t.Empty(r.uri)
	r.SetURI("test/uri")
	t.Equal("test/uri", r.uri)
}

func (t *ClientTestSuite) TestRequestSetBody() {
	r := NewClient("local:123").Request()
	t.Empty(r.body)
	r.SetBody([]byte("some body"))
	t.Equal([]byte("some body"), r.body)
}

func (t *ClientTestSuite) TestRequestSetHeaders() {
	r := NewClient("local:123").Request()
	t.Empty(r.headers)
	r.SetHeaders(map[string]string{"x-header": "some header"})
	t.NotEmpty(r.headers)
	t.Equal("some header", r.headers["x-header"])
}

func (t *ClientTestSuite) TestRequestSetHeader() {
	r := NewClient("local:123").Request()
	t.Empty(r.headers)
	r.SetHeader("x-header-1", "some header")
	t.NotEmpty(r.headers)
	t.Equal("some header", r.headers["x-header-1"])
}

func (t *ClientTestSuite) TestRequestDo() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"some":"response"}`))
	}))
	r1 := NewClient(s.URL[7:]).Request()
	rs1, err := r1.Do()
	t.NoError(err)
	t.NotNil(rs1)
	t.Equal([]byte(`{"some":"response"}`), rs1.Body())
}

func (t *ClientTestSuite) TestResponseReadJSON() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"error":"random data"}`))
	}))
	r := NewClient(s.URL[7:]).Request()
	rs, err := r.Do()
	t.NoError(err)
	t.NotNil(rs)
	tt := new(Error)
	err = rs.ReadJSON(tt)
	t.NoError(err)
	t.Equal("random data", tt.Err)
}
