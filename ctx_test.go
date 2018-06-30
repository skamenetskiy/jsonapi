package jsonapi

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
)

func TestCtx(t *testing.T) {
	suite.Run(t, new(CtxTestSuite))
}

type CtxTestSuite struct {
	suite.Suite
}

func (t *CtxTestSuite) TestGetParamString() {
	ctx := &Ctx{&fasthttp.RequestCtx{}}
	ctx.SetUserValue("param", "some string")
	t.Equal("some string", ctx.GetParamString("param"))
}

func (t *CtxTestSuite) TestGetParamInt() {
	ctx := &Ctx{&fasthttp.RequestCtx{}}
	ctx.SetUserValue("param", "123")
	v, err := ctx.GetParamInt("param")
	t.NoError(err)
	t.Nil(err)
	t.Equal(123, v)
	t.IsType(int(0), v)
	ctx.SetUserValue("param2", "abc")
	v2, err := ctx.GetParamInt("param2")
	t.Error(err)
	t.NotNil(err)
	t.Empty(v2)
}

func (t *CtxTestSuite) TestGetParamInt64() {
	ctx := &Ctx{&fasthttp.RequestCtx{}}
	ctx.SetUserValue("param", "123")
	v, err := ctx.GetParamInt64("param")
	t.NoError(err)
	t.Nil(err)
	t.Equal(int64(123), v)
	t.IsType(int64(0), v)
	ctx.SetUserValue("param2", "abc")
	v2, err := ctx.GetParamInt64("param2")
	t.Error(err)
	t.NotNil(err)
	t.Empty(v2)
}

func (t *CtxTestSuite) TestGetParamUint64() {
	ctx := &Ctx{&fasthttp.RequestCtx{}}
	ctx.SetUserValue("param", "123")
	v, err := ctx.GetParamUint64("param")
	t.NoError(err)
	t.Nil(err)
	t.Equal(uint64(123), v)
	t.IsType(uint64(0), v)
	ctx.SetUserValue("param2", "abc")
	v2, err := ctx.GetParamUint64("param2")
	t.Error(err)
	t.NotNil(err)
	t.Empty(v2)
}

func (t *CtxTestSuite) TestGetParamFloat64() {
	ctx := &Ctx{&fasthttp.RequestCtx{}}
	ctx.SetUserValue("param", "123.123")
	v, err := ctx.GetParamFloat64("param")
	t.NoError(err)
	t.Nil(err)
	t.Equal(float64(123.123), v)
	t.IsType(float64(0), v)
	ctx.SetUserValue("param2", "abc")
	v2, err := ctx.GetParamFloat64("param2")
	t.Error(err)
	t.NotNil(err)
	t.Empty(v2)
}
