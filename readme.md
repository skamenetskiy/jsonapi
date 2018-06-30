[![Build Status](https://travis-ci.org/skamenetskiy/jsonapi.svg?branch=master)](https://travis-ci.org/skamenetskiy/jsonapi)
[![codecov](https://codecov.io/gh/skamenetskiy/jsonapi/branch/master/graph/badge.svg)](https://codecov.io/gh/skamenetskiy/jsonapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/skamenetskiy/jsonapi)](https://goreportcard.com/report/github.com/skamenetskiy/jsonapi)
[![godoc](https://godoc.org/github.com/skamenetskiy/jsonapi?status.svg)](http://godoc.org/github.com/skamenetskiy/jsonapi)

# jsonapi
A small library, built on top of fasthttp and fasthttprouter

# Quick start
```go
func main() {
	err := jsonapi.
		NewServer().
		Get("/" func(ctx *Ctx) {
		    ctx.WriteJSON()
		}).
		Listen()
	if err != nil {
		log.Fatal(err)
	}
}
```

# Dependencies
- [fasthttp](https://github.com/valyala/fasthttp)
- [fasthttprouter](https://github.com/buaazp/fasthttprouter)
- [easyjson](https://github.com/mailru/easyjson)

# Install
```
go get -u github.com/skamenetskiy/jsonapi
```
