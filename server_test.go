package jsonapi

import (
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/stretchr/testify/suite"
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
	t.True(s.authFunc(new(Ctx)))
}
