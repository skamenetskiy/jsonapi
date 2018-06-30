package jsonapi

import (
	"encoding/json"
)

// CRUDController interface defines a basic CRUDController controller
type CRUDController interface {
	Create(*Ctx) *Result  // POST   /path
	Get(*Ctx) *Result     // GET    /path
	GetByID(*Ctx) *Result // GET    /path/:id
	Update(*Ctx) *Result  // PUT    /path/:id
	Delete(*Ctx) *Result  // DELETE /path/:id
}

func NewCRUDClient(addr string, path string) *CRUDClient {
	c := &CRUDClient{
		client: NewClient(addr),
		path:   path,
	}
	return c
}

type CRUDClient struct {
	client *Client
	path   string
}

func (c *CRUDClient) Create(v CRUDModel) error {
	panic("implement me")
}

func (c *CRUDClient) Get(v CRUDModel) error {
	panic("implement me")
}

func (c *CRUDClient) GetByID(id ID, v CRUDModel) {
	panic("implement me")
}

func (c *CRUDClient) Update(id ID, v CRUDModel) {
	panic("implement me")
}

func (c *CRUDClient) Delete(id ID) {
	panic("implement me")
}

type ID interface{}

type CRUDModel interface {
	json.Marshaler
	json.Unmarshaler
	GetID() ID
	SetID(id ID)
}

func getCrudPath(path string) string {
	return path + "/:id"
}
