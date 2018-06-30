package jsonapi

import (
	"github.com/stretchr/testify/suite"
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
	c := NewClient(":123")
	t.False(c.useSSL)
	c.UseSSL()
	t.True(c.useSSL)
}
