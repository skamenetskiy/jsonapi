package jsonapi

// ClientAuthFunc is modifying the client request
// to implement authentication in client
type ClientAuthFunc func(*Request)

// ServerAuthFunc is checking the server request for
// authentication and returns a bool
type ServerAuthFunc func(*Ctx) bool
