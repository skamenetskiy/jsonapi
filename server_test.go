package jsonapi

import (
	"fmt"
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
	t.IsType(&fasthttprouter.Router{}, s.router)
	t.NotNil(s.authFunc)
	t.NotEmpty(s.authFunc)
	t.True(s.authFunc(new(Ctx)))
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
		fmt.Println("Tes")
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
