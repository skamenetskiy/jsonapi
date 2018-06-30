package jsonapi

import (
	"errors"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	DefaultAddr = ":80"

	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

// Handler defines the handler func
type Handler func(*Ctx)

// ControllerHandler defines controller handler func
type ControllerHandler func(*Ctx) *Result

// NewServer creates a new jsonapi Server
func NewServer(addr ...string) *Server {
	s := &Server{
		router:   fasthttprouter.New(),
		authFunc: func(*Ctx) bool { return true },
	}
	if len(addr) == 1 {
		s.addr = addr[0]
	} else {
		s.addr = DefaultAddr
	}
	return s
}

// Server is an http server wrapper
type Server struct {
	addr     string
	authFunc ServerAuthFunc
	router   *fasthttprouter.Router
}

// Listen starts http server and listens on defined addr
func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(s.addr, s.router.Handler)
}

// ListenTLS starts http server and listens on defined addr with TLS
// Reads TLS certificate from certFile and key from keyFile
func (s *Server) ListenTLS(certFile string, keyFile string) error {
	return fasthttp.ListenAndServeTLS(s.addr, certFile, keyFile, s.router.Handler)
}

// ListenTLSEmbed starts http server and listens on defined addr with TLS
// Accepts TLS certificate in cert and key in key
func (s *Server) ListenTLSEmbed(cert []byte, key []byte) error {
	return fasthttp.ListenAndServeTLSEmbed(s.addr, cert, key, s.router.Handler)
}

// ListerUNIX starts http server and listens on UNIX socket
// Accepts mode as file mode
func (s *Server) ListenUNIX(mode os.FileMode) error {
	return fasthttp.ListenAndServeUNIX(s.addr, mode, s.router.Handler)
}

// SetAddr sets listen address
func (s *Server) SetAddr(addr string) *Server {
	s.addr = addr
	return s
}

// SetAuthFunc sets authentication func that will be triggered
// on every request to the server
func (s *Server) SetAuthFunc(authFunc ServerAuthFunc) {
	s.authFunc = authFunc
}

// Route adds a new route handler to router
func (s *Server) Route(method string, path string, handler Handler) *Server {
	s.router.Handle(method, path, func(ctx *fasthttp.RequestCtx) {
		c := &Ctx{ctx}
		c.SetHeader("Content-Type", "application/json")
		c.SetHeader("Server", "jsonapi @ fasthttp")
		// check auth
		if !s.authFunc(c) {
			// return unauthorized error
			c.ErrUnauthorized(ErrUnauthorized)
			return
		}
		// execute handler
		handler(c)
	})
	return s
}

// Get is a shortcut to Route("GET"...)
func (s *Server) Get(path string, handler Handler) *Server {
	return s.Route(MethodGet, path, handler)
}

// Head is a shortcut to Route("HEAD"...)
func (s *Server) Head(path string, handler Handler) *Server {
	return s.Route(MethodHead, path, handler)
}

// Post is a shortcut to Route("POST"...)
func (s *Server) Post(path string, handler Handler) *Server {
	return s.Route(MethodPost, path, handler)
}

// Put is a shortcut to Route("PUT"...)
func (s *Server) Put(path string, handler Handler) *Server {
	return s.Route(MethodPut, path, handler)
}

// Patch is a shortcut to Route("PATCH"...)
func (s *Server) Patch(path string, handler Handler) *Server {
	return s.Route(MethodPatch, path, handler)
}

// Delete is a shortcut to Route("DELETE"...)
func (s *Server) Delete(path string, handler Handler) *Server {
	return s.Route(MethodDelete, path, handler)
}

// Connect is a shortcut to Route("CONNECT"...)
func (s *Server) Connect(path string, handler Handler) *Server {
	return s.Route(MethodConnect, path, handler)
}

// Options is a shortcut to Route("OPTIONS"...)
func (s *Server) Options(path string, handler Handler) *Server {
	return s.Route(MethodOptions, path, handler)
}

// Trace is a shortcut to Route("TRACE"...)
func (s *Server) Trace(path string, handler Handler) *Server {
	return s.Route(MethodTrace, path, handler)
}

// ControllerMethod registers a controller method handler by http method and path
// This method can be called directly without a controller.
func (s *Server) ControllerMethod(method string, path string, handler ControllerHandler) *Server {
	s.Route(method, path, func(ctx *Ctx) {
		res := handler(ctx)
		if res.Err != nil {
			ctx.Err(res.Err, res.Err.Code)
			return
		}
		ctx.OK(res.Data)
	})
	return s
}

// Controller registers a controller
func (s *Server) Controller(ctrl Controller) *Server {
	for method, paths := range ctrl.Methods() {
		for path, handler := range paths {
			s.ControllerMethod(method, path, handler)
		}
	}
	return s
}

// CRUDController assigns a crud controller to a path
func (s *Server) CRUDController(path string, ctrl CRUDController) *Server {
	// handle POST/Create
	s.ControllerMethod(MethodPost, path, ctrl.Create)
	// handler GET/Get
	s.ControllerMethod(MethodGet, path, ctrl.Get)
	// Handle GET/GetByID
	s.ControllerMethod(MethodGet, getCrudPath(path), ctrl.GetByID)
	// Handle PUT/Update
	s.ControllerMethod(MethodPut, getCrudPath(path), ctrl.Update)
	// Handle DELETE/Delete
	s.ControllerMethod(MethodDelete, getCrudPath(path), ctrl.Delete)
	return s
}
