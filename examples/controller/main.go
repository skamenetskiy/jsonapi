package main

import (
	"gitlab.com/ngrs/lib-go/jsonapi"
	"log"
	"errors"
	"strconv"
)

func main() {
	err := jsonapi.
		NewServer().
		Controller(new(controller)).
		Listen()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	names = &Names{
		{1, "Peter"},
		{2, "Tom"},
		{3, "John"},
	}
)

//go:generate easyjson
//easyjson:json
type Name struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//go:generate easyjson
//easyjson:json
type Names []*Name

type controller struct {
	jsonapi.BaseController
}

//...
func (c *controller) Methods() jsonapi.ControllerMethods {
	return jsonapi.ControllerMethods{
		jsonapi.MethodGet: {
			"/":    c.getNames,
			"/:id": c.getNameByID,
		},
	}
}

func (c *controller) getNames(ctx *jsonapi.Ctx) *jsonapi.Result {
	return c.OK(names)
}

func (c *controller) getNameByID(ctx *jsonapi.Ctx) *jsonapi.Result {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	for _, name := range *names {
		if name.ID == id {
			return c.OK(name)
		}
	}
	return c.ErrNotFound(errors.New("name not found"))
}
