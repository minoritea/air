// Package air provides a http router which wraps `github.com/julienschmidt/httprouter`
// to support `http.Handler` with URL params via `Context` package.
package air

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type contextKey struct{ name string }

// ParamsKey is the key to get URL params from a Context.
var ParamsKey = &contextKey{"key of params"}

// Router is a simple URL router which acts as a http.Handler.
type Router struct{ *httprouter.Router }

// New returns a new Router.
func New() *Router {
	return &Router{httprouter.New()}
}

type handlerWrap struct{ http.Handler }

func (hw *handlerWrap) handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r = r.WithContext(context.WithValue(r.Context(), "", params))
	hw.Handler.ServeHTTP(w, r)
}

// Handle registers a new http.Handler to the Router.
func (r *Router) Handle(method, path string, h http.Handler) {
	w := &handlerWrap{h}
	r.Router.Handle(method, path, w.handle)
}

// DELETE is a shortcut for Handle with "DELETE" method.
func (r *Router) DELETE(path string, h http.Handler) { r.Handle("DELETE", path, h) }

// GET is a shortcut for Handle with "GET" method.
func (r *Router) GET(path string, h http.Handler) { r.Handle("GET", path, h) }

// HEAD is a shortcut for Handle with "HEAD" method.
func (r *Router) HEAD(path string, h http.Handler) { r.Handle("HEAD", path, h) }

// OPTIONS is a shortcut for Handle with "OPTIONS" method.
func (r *Router) OPTIONS(path string, h http.Handler) { r.Handle("OPTIONS", path, h) }

// PATCH is a shortcut for Handle with "PATCH" method.
func (r *Router) PATCH(path string, h http.Handler) { r.Handle("PATCH", path, h) }

// POST is a shortcut for Handle with "POST" method.
func (r *Router) POST(path string, h http.Handler) { r.Handle("POST", path, h) }

// PUT is a shortcut for Handle with "PUT" method.
func (r *Router) PUT(path string, h http.Handler) { r.Handle("PUT", path, h) }

// Param returns the URL parameter from a http.Request object.
func Param(r *http.Request, key string) string {
	params, ok := r.Context().Value(ParamsKey).(httprouter.Params)
	if !ok {
		return ""
	}
	return params.ByName(key)
}

// Middleware is an alias of `func(http.Handler) http.Handler`.
// These functions wrap a http.Handler with some additonal features.
//  func middleware(next http.Handler) http.Handler {
//      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//          // some additional features...
//          next(w, r)
//      })
//  }
type Middleware func(http.Handler) http.Handler

// Compose assembles middlewares into a http.Handler.
// Note that middlewares are applied from right to left, then handler is called.
func Compose(h http.Handler, mws ...Middleware) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}

// Composer returns a new middleware which is composed by passed middlewares.
// It is just a partially applied function of `Compose` .
func Composer(mws ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		return Compose(h, mws...)
	}
}

// H is a copy of http.HandlerFunc.
// It makes easy to cast handler functions to http.Handler.
type H func(http.ResponseWriter, *http.Request)

// ServeHTTP is required to implement http.Handler interface.
func (f H) ServeHTTP(w http.ResponseWriter, r *http.Request) { f(w, r) }
