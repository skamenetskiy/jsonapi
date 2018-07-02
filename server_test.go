package jsonapi

import (
	"net"
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
}

func (t *ServerTestSuite) TestNewServer() {
	s := NewServer()
	t.NotNil(s)
	t.IsType(&Server{}, s)
	t.Nil(s.ln)
	t.IsType(&fasthttprouter.Router{}, s.router)
	t.NotNil(s.authFunc)
	t.NotEmpty(s.authFunc)
	t.True(s.authFunc(new(Ctx)))
	s2 := NewServer("asd:asd")
	t.NotNil(s2)
	t.IsType(&Server{}, s2)
	t.Nil(s.ln)
	t.IsType(&fasthttprouter.Router{}, s2.router)
	t.NotNil(s2.authFunc)
	t.NotEmpty(s2.authFunc)
}

func (t *ServerTestSuite) TestServerListen() {
	err := NewServer("asd:asd").Listen()
	t.Error(err)
}

func (t *ServerTestSuite) TestServerNewListener() {
	err := NewServer("127.0.0.1:").newListener()
	t.NoError(err)
	err = NewServer("asd:asd").newListener()
	t.Error(err)
}

func (t *ServerTestSuite) TestServerSetAddr() {
	s := NewServer()
	t.Equal(DefaultAddr, s.addr)
	s.SetAddr("addr:877")
	t.Equal("addr:877", s.addr)
}

func (t *ServerTestSuite) TestServerSetAuthFunc() {
	s := NewServer()
	t.NotEmpty(s.authFunc)
	t.True(s.authFunc(new(Ctx)))
	s.SetAuthFunc(func(*Ctx) bool { return false })
	t.False(s.authFunc(new(Ctx)))
}

func (t *ServerTestSuite) TestServerRoute() {
	ln := fasthttputil.NewInmemoryListener()
	s := NewServer().SetListener(ln)
	s.Route(MethodGet, "/a1", func(ctx *Ctx) {
		ctx.Write([]byte(`{"path":"` + string(ctx.Path()) + `"}`))
	})
	go s.Listen()
	defer s.ln.Close()
	cl := fasthttp.Client{
		Dial: fasthttp.DialFunc(func(string) (net.Conn, error) {
			return ln.Dial()
		}),
	}
	rq := fasthttp.AcquireRequest()
	rq.SetRequestURI("http://" + ln.Addr().String() + "/a1")
	rq.Header.SetMethod(MethodGet)
	rs := fasthttp.AcquireResponse()
	err := cl.Do(rq, rs)
	t.NoError(err)
	t.NotNil(rs)
	t.Equal([]byte(`{"path":"/a1"}`), rs.Body())
}

func (t *ServerTestSuite) TestServerGet() {
	ln, s := t.getServer()
	s.Get("/a1", func(ctx *Ctx) {
		ctx.Write([]byte(`{"path":"` + string(ctx.Path()) + `"}`))
	})
	go s.Listen()
	defer s.ln.Close()
	rq, rs, err := t.request(ln, MethodGet)
	t.NoError(err)
	t.NotNil(rq)
	t.NotNil(rs)
	t.Equal([]byte(`{"path":"/a1"}`), rs.Body())
}

func (t *ServerTestSuite) TestServerPost() {
	ln, s := t.getServer()
	s.Post("/a1", func(ctx *Ctx) {
		ctx.Write([]byte(`{"path":"` + string(ctx.Path()) + `"}`))
	})
	go s.Listen()
	defer s.ln.Close()
	rq, rs, err := t.request(ln, MethodPost)
	t.NoError(err)
	t.NotNil(rq)
	t.NotNil(rs)
	t.Equal([]byte(`{"path":"/a1"}`), rs.Body())
}

func (t *ServerTestSuite) TestServerPut() {
	ln, s := t.getServer()
	s.Put("/a1", func(ctx *Ctx) {
		ctx.Write([]byte(`{"path":"` + string(ctx.Path()) + `"}`))
	})
	go s.Listen()
	defer s.ln.Close()
	rq, rs, err := t.request(ln, MethodPut)
	t.NoError(err)
	t.NotNil(rq)
	t.NotNil(rs)
	t.Equal([]byte(`{"path":"/a1"}`), rs.Body())
}

func (t *ServerTestSuite) TestServerDelete() {
	ln, s := t.getServer()
	s.Delete("/a1", func(ctx *Ctx) {
		ctx.Write([]byte(`{"path":"` + string(ctx.Path()) + `"}`))
	})
	go s.Listen()
	defer s.ln.Close()
	rq, rs, err := t.request(ln, MethodDelete)
	t.NoError(err)
	t.NotNil(rq)
	t.NotNil(rs)
	t.Equal([]byte(`{"path":"/a1"}`), rs.Body())
}

func (t *ServerTestSuite) TestServerController() {

}

func (t *ServerTestSuite) getServer() (*fasthttputil.InmemoryListener, *Server) {
	ln := fasthttputil.NewInmemoryListener()
	return ln, NewServer().SetListener(ln)
}

func (t *ServerTestSuite) getClient(ln *fasthttputil.InmemoryListener) fasthttp.Client {
	return fasthttp.Client{
		Dial: fasthttp.DialFunc(func(string) (net.Conn, error) {
			return ln.Dial()
		}),
	}
}

func (t *ServerTestSuite) request(ln *fasthttputil.InmemoryListener, method string) (*fasthttp.Request, *fasthttp.Response, error) {
	rq := fasthttp.AcquireRequest()
	rq.SetRequestURI("http://" + ln.Addr().String() + "/a1")
	rq.Header.SetMethod(method)
	rs := fasthttp.AcquireResponse()
	cl := t.getClient(ln)
	return rq, rs, cl.Do(rq, rs)
}

type c1 struct {
	BaseController
}

func (c *c1) Methods() ControllerMethods {
	return ControllerMethods{
		MethodGet: {

		},
	}
}
