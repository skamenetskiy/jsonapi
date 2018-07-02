package jsonapi

import (
	"net"
	"os"
	"path"
	"sync"

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

	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
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
		mu:       new(sync.Mutex),
	}
	s.mu.Lock()
	if len(addr) == 1 {
		s.addr = addr[0]
	} else {
		s.addr = DefaultAddr
	}
	s.mu.Unlock()
	return s
}

// Server is an http server wrapper
type Server struct {
	addr     string
	authFunc ServerAuthFunc
	ln       net.Listener
	router   *fasthttprouter.Router
	mu       *sync.Mutex
}

// Listen starts http server and listens on defined addr
func (s *Server) Listen() error {
	if err := s.newListener(); err != nil {
		return err
	}
	return fasthttp.Serve(s.ln, s.router.Handler)
}

// ListenTLS starts http server and listens on defined addr with TLS
// Reads TLS certificate from certFile and key from keyFile
func (s *Server) ListenTLS(certFile string, keyFile string) error {
	if err := s.newListener(); err != nil {
		return err
	}
	return fasthttp.ServeTLS(s.ln, certFile, keyFile, s.router.Handler)
}

// ListenTLSEmbed starts http server and listens on defined addr with TLS
// Accepts TLS certificate in cert and key in key
func (s *Server) ListenTLSEmbed(cert []byte, key []byte) error {
	if err := s.newListener(); err != nil {
		return err
	}
	return fasthttp.ServeTLSEmbed(s.ln, cert, key, s.router.Handler)
}

// ListenUNIX starts http server and listens on UNIX socket
// Accepts mode as file mode
func (s *Server) ListenUNIX(mode os.FileMode) error {
	return fasthttp.ListenAndServeUNIX(s.addr, mode, s.router.Handler)
}

// SetListener sets net.Listener that will be used
func (s *Server) SetListener(ln net.Listener) *Server {
	s.ln = ln
	return s
}

// SetAddr sets listen address
func (s *Server) SetAddr(addr string) *Server {
	s.addr = addr
	return s
}

func (s *Server) GetAddr() string {
	if s.ln != nil {
		return s.ln.Addr().String()
	}
	return s.addr
}

// SetAuthFunc sets authentication func that will be triggered
// on every request to the server
func (s *Server) SetAuthFunc(authFunc ServerAuthFunc) {
	s.authFunc = authFunc
}

// Route adds a new route handler to router
func (s *Server) Route(method string, p string, handler Handler) *Server {
	s.router.Handle(method, p, func(ctx *fasthttp.RequestCtx) {
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
func (s *Server) Get(p string, handler Handler) *Server {
	return s.Route(MethodGet, p, handler)
}

// Head is a shortcut to Route("HEAD"...)
func (s *Server) Head(p string, handler Handler) *Server {
	return s.Route(MethodHead, p, handler)
}

// Post is a shortcut to Route("POST"...)
func (s *Server) Post(p string, handler Handler) *Server {
	return s.Route(MethodPost, p, handler)
}

// Put is a shortcut to Route("PUT"...)
func (s *Server) Put(p string, handler Handler) *Server {
	return s.Route(MethodPut, p, handler)
}

// Patch is a shortcut to Route("PATCH"...)
func (s *Server) Patch(p string, handler Handler) *Server {
	return s.Route(MethodPatch, p, handler)
}

// Delete is a shortcut to Route("DELETE"...)
func (s *Server) Delete(p string, handler Handler) *Server {
	return s.Route(MethodDelete, p, handler)
}

// Connect is a shortcut to Route("CONNECT"...)
func (s *Server) Connect(p string, handler Handler) *Server {
	return s.Route(MethodConnect, p, handler)
}

// Options is a shortcut to Route("OPTIONS"...)
func (s *Server) Options(p string, handler Handler) *Server {
	return s.Route(MethodOptions, p, handler)
}

// Trace is a shortcut to Route("TRACE"...)
func (s *Server) Trace(p string, handler Handler) *Server {
	return s.Route(MethodTrace, p, handler)
}

// ControllerMethod registers a controller method handler by http method and path
// This method can be called directly without a controller.
func (s *Server) ControllerMethod(method string, p string, handler ControllerHandler) *Server {
	s.Route(method, p, func(ctx *Ctx) {
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
func (s *Server) Controller(basePath string, ctrl Controller) *Server {
	for method, paths := range ctrl.Methods() {
		for p, handler := range paths {
			s.ControllerMethod(method, path.Join(basePath, p), handler)
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

func (s *Server) newListener() error {
	s.mu.Lock()
	if s.ln == nil {
		var err error
		s.ln, err = net.Listen("tcp4", s.addr)
		if err != nil {
			s.mu.Unlock()
			return err
		}
	}
	s.mu.Unlock()
	return nil
}
