package jsonapi

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCRUD(t *testing.T) {
	suite.Run(t, new(CRUDTestSuite))
}

type CRUDTestSuite struct {
	suite.Suite
}

func (t *CRUDTestSuite) TestNewCRUDClient() {
	c := NewCRUDClient("321:123", "/a/b/c")
	t.IsType(&CRUDClient{}, c)
	t.IsType(&Client{}, c.client)
	t.Equal("/a/b/c", c.path)
}
