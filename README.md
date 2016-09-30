

# air
`import "github.com/minoritea/air"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package air provides a http router which wraps `github.com/julienschmidt/httprouter`
to support `http.Handler` with URL params via `Context` package.




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Compose(h http.Handler, mws ...Middleware) http.Handler](#Compose)
* [func Param(r *http.Request, key string) string](#Param)
* [type H](#H)
  * [func (f H) ServeHTTP(w http.ResponseWriter, r *http.Request)](#H.ServeHTTP)
* [type Middleware](#Middleware)
  * [func Composer(mws ...Middleware) Middleware](#Composer)
* [type Router](#Router)
  * [func New() *Router](#New)
  * [func (r *Router) DELETE(path string, h http.Handler)](#Router.DELETE)
  * [func (r *Router) GET(path string, h http.Handler)](#Router.GET)
  * [func (r *Router) HEAD(path string, h http.Handler)](#Router.HEAD)
  * [func (r *Router) Handle(method, path string, h http.Handler)](#Router.Handle)
  * [func (r *Router) OPTIONS(path string, h http.Handler)](#Router.OPTIONS)
  * [func (r *Router) PATCH(path string, h http.Handler)](#Router.PATCH)
  * [func (r *Router) POST(path string, h http.Handler)](#Router.POST)
  * [func (r *Router) PUT(path string, h http.Handler)](#Router.PUT)


#### <a name="pkg-files">Package files</a>
[air.go](/src/github.com/minoritea/air/air.go) 



## <a name="pkg-variables">Variables</a>
``` go
var ParamsKey = &contextKey{"key of params"}
```
ParamsKey is the key to get URL params from a Context.



## <a name="Compose">func</a> [Compose](/src/target/air.go?s=2714:2774#L70)
``` go
func Compose(h http.Handler, mws ...Middleware) http.Handler
```
Compose assembles middlewares into a http.Handler.
Note that middlewares are applied from right to left, then handler is called.



## <a name="Param">func</a> [Param](/src/target/air.go?s=2019:2065#L50)
``` go
func Param(r *http.Request, key string) string
```
Param returns the URL parameter from a http.Request object.




## <a name="H">type</a> [H](/src/target/air.go?s=3187:3234#L87)
``` go
type H func(http.ResponseWriter, *http.Request)
```
H is a copy of http.HandlerFunc.
It makes easy to cast handler functions to http.Handler.










### <a name="H.ServeHTTP">func</a> (H) [ServeHTTP](/src/target/air.go?s=3298:3358#L90)
``` go
func (f H) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP is required to implement http.Handler interface.




## <a name="Middleware">type</a> [Middleware](/src/target/air.go?s=2530:2577#L66)
``` go
type Middleware func(http.Handler) http.Handler
```
Middleware is an alias of `func(http.Handler) http.Handler`.
These functions wrap a http.Handler with some additonal features.


	func middleware(next http.Handler) http.Handler {
	    return http.HandlerFunc(w http.ResponseWriter, r *http.Request) {
	        // some additional features...
	        next(w, r)
	    }
	}







### <a name="Composer">func</a> [Composer](/src/target/air.go?s=2967:3010#L79)
``` go
func Composer(mws ...Middleware) Middleware
```
Composer returns a new middleware which is composed by passed middlewares.
It is just a partially applied function of `Compose` .





## <a name="Router">type</a> [Router](/src/target/air.go?s=450:490#L8)
``` go
type Router struct{ *httprouter.Router }
```
Router is a simple URL router which acts as a http.Handler.







### <a name="New">func</a> [New](/src/target/air.go?s=521:539#L11)
``` go
func New() *Router
```
New returns a new Router.





### <a name="Router.DELETE">func</a> (\*Router) [DELETE](/src/target/air.go?s=1052:1104#L29)
``` go
func (r *Router) DELETE(path string, h http.Handler)
```
DELETE is a shortcut for Handle with "DELETE" method.




### <a name="Router.GET">func</a> (\*Router) [GET](/src/target/air.go?s=1189:1238#L32)
``` go
func (r *Router) GET(path string, h http.Handler)
```
GET is a shortcut for Handle with "GET" method.




### <a name="Router.HEAD">func</a> (\*Router) [HEAD](/src/target/air.go?s=1322:1372#L35)
``` go
func (r *Router) HEAD(path string, h http.Handler)
```
HEAD is a shortcut for Handle with "HEAD" method.




### <a name="Router.Handle">func</a> (\*Router) [Handle](/src/target/air.go?s=866:926#L23)
``` go
func (r *Router) Handle(method, path string, h http.Handler)
```
Handle registers a new http.Handler to the Router.




### <a name="Router.OPTIONS">func</a> (\*Router) [OPTIONS](/src/target/air.go?s=1463:1516#L38)
``` go
func (r *Router) OPTIONS(path string, h http.Handler)
```
OPTIONS is a shortcut for Handle with "OPTIONS" method.




### <a name="Router.PATCH">func</a> (\*Router) [PATCH](/src/target/air.go?s=1606:1657#L41)
``` go
func (r *Router) PATCH(path string, h http.Handler)
```
PATCH is a shortcut for Handle with "PATCH" method.




### <a name="Router.POST">func</a> (\*Router) [POST](/src/target/air.go?s=1743:1793#L44)
``` go
func (r *Router) POST(path string, h http.Handler)
```
POST is a shortcut for Handle with "POST" method.




### <a name="Router.PUT">func</a> (\*Router) [PUT](/src/target/air.go?s=1876:1925#L47)
``` go
func (r *Router) PUT(path string, h http.Handler)
```
PUT is a shortcut for Handle with "PUT" method.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
