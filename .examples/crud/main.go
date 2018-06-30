package main

import (
	"errors"

	"github.com/skamenetskiy/jsonapi"
)

func main() {
	jsonapi.
		NewServer().
		CRUDController("/", new(controller)).
		Listen()
}

var (
	data = new(animals)
)

// animal model
//go:generate easyjson
//easyjson:json
type animal struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// animals list
//go:generate easyjson
//easyjson:json
type animals []*animal

func (a *animals) getByID(id string) *animal {
	for _, v := range *a {
		if v.ID == id {
			return v
		}
	}
	return nil
}

type controller struct {
	jsonapi.BaseController
}

func (c *controller) Create(*jsonapi.Ctx) *jsonapi.Result {
	panic("implement me")
}

func (c *controller) Get(*jsonapi.Ctx) *jsonapi.Result {
	return c.OK(data)
}

func (c *controller) GetByID(ctx *jsonapi.Ctx) *jsonapi.Result {
	v := data.getByID(ctx.GetParamString("id"))
	if v == nil {
		return c.ErrNotFound(errors.New("animal not found"))
	}
	return c.OK(v)
}

func (c *controller) Update(*jsonapi.Ctx) *jsonapi.Result {
	panic("implement me")
}

func (c *controller) Delete(*jsonapi.Ctx) *jsonapi.Result {
	panic("implement me")
}
