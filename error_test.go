package jsonapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestError(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}

type ErrorTestSuite struct {
	suite.Suite
}

func (t *ErrorTestSuite) TestNewError() {
	err := errors.New("some error")
	e := NewError(err)
	t.NotNil(e)
	t.IsType(&Error{}, e)
	t.Empty(e.Code)
	t.Equal(err.Error(), e.Err)
	e2 := NewError(err, 123, 321)
	t.NotNil(e2)
	t.IsType(&Error{}, e2)
	t.NotEmpty(e2.Code)
	t.Equal(err.Error(), e.Err)
	t.Equal(123, e2.Code)
}

func (t *ErrorTestSuite) TestError() {
	e := NewError(errors.New("great error"), 456)
	t.Equal("456: great error", e.Error())
}

func (t *ErrorTestSuite) TestMarshalJSON() {
	e := NewError(errors.New("some error"), 678)
	j, err := e.MarshalJSON()
	t.NoError(err)
	t.Nil(err)
	t.Equal([]byte(`{"error":"some error","code":678}`), j)
}

func (t *ErrorTestSuite) TestUnmarshalJSON() {
	e := &Error{}
	err := e.UnmarshalJSON([]byte(`{"error":"big error","code":678}`))
	t.NoError(err)
	t.Nil(err)
	t.Equal(e.Err, "big error")
	t.Equal(e.Code, 678)
}
