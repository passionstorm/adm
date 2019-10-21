package context

import (
	"bytes"
	"encoding/json"
	"github.com/CloudyKit/jet"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	Request        *http.Request
	Response       *http.Response
	ResponseWriter http.ResponseWriter
	Data           jet.VarMap
	index          int8
	handlers       Handlers
}

// Path is used in the matching of request and response. URL stores the
// raw register url. RegUrl contains the wildcard which on behalf of
// the route params.
type Path struct {
	URL    string
	Method string
}

// Path return the url path.
func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

// Abort abort the context.
func (ctx *Context) Abort() {
	ctx.index = abortIndex
}

// Next should be used only inside middleware.
func (ctx *Context) Next() {
	ctx.index++
	for s := int8(len(ctx.handlers)); ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

// Next should be used only inside middleware.
func (ctx *Context) SetHandlers(handlers Handlers) *Context {
	ctx.handlers = handlers
	return ctx
}

// Method return the request method.
func (ctx *Context) Method() string {
	return ctx.Request.Method
}

// NewContext used in adapter which return a Context with request
// and slice of UserValue and a default Response.
func NewContext(req *http.Request, resWriter http.ResponseWriter) *Context {

	return &Context{
		Request: req,
		Data:    make(jet.VarMap),
		Response: &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
		},
		ResponseWriter: resWriter,
		index:          -1,
	}
}

// Write save the given status code, header and body string into the response.
func (ctx *Context) Write(code int, Header map[string]string, Body string) {
	ctx.Response.StatusCode = code
	for key, head := range Header {
		ctx.AddHeader(key, head)
	}
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

// Json serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (ctx *Context) Json(code int, Body map[string]interface{}) {
	ctx.Response.StatusCode = code
	ctx.AddHeader("Content-Type", "application/json")
	BodyStr, err := json.Marshal(Body)
	if err != nil {
		panic(err)
	}
	ctx.Response.Body = ioutil.NopCloser(bytes.NewReader(BodyStr))
}

// Data writes some data into the body stream and updates the HTTP code.
//func (ctx *Context) Data(code int, contentType string, data []byte) {
//	ctx.Response.StatusCode = code
//	ctx.AddHeader("Content-Type", contentType)
//	ctx.Response.Body = ioutil.NopCloser(bytes.NewBuffer(data))
//}

// Html output html response.
func (ctx *Context) Html(code int, body string) {
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.SetStatusCode(code)
	ctx.WriteString(body)
}

func (t *Context) Render(view string) {
	var root, _ = os.Getwd()
	var View = jet.NewHTMLSet(filepath.Join(root, "views"))
	View.SetDevelopmentMode(true)
	templ, err := View.GetTemplate(view)
	if err != nil {
		log.Println(err)
	}
	err = templ.Execute(t.ResponseWriter, t.Data, nil)
	if err != nil {
		log.Println(err)
	}
}

// Write save the given body string into the response.
func (ctx *Context) WriteString(Body string) {
	ctx.Response.Body = ioutil.NopCloser(strings.NewReader(Body))
}

// SetStatusCode save the given status code into the response.
func (ctx *Context) SetStatusCode(code int) {
	ctx.Response.StatusCode = code
}

// SetStatusCode save the given content type header into the response header.
func (ctx *Context) SetContentType(contentType string) {
	ctx.AddHeader("Content-Type", contentType)
}

// LocalIP return the request client ip.
func (ctx *Context) LocalIP() string {
	return "127.0.0.1"
}

// SetCookie save the given cookie obj into the response Set-Cookie header.
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	if v := cookie.String(); v != "" {
		ctx.AddHeader("Set-Cookie", v)
	}
}

// Query get the query parameter of url.
func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

// QueryDefault get the query parameter of url. If it is empty, return the default.
func (ctx *Context) QueryDefault(key, def string) string {
	value := ctx.Query(key)
	if value == "" {
		return def
	}
	return value
}

// Headers get the value of request headers key.
func (ctx *Context) Headers(key string) string {
	return ctx.Request.Header.Get(key)
}

// FormValue get the value of request form key.
func (ctx *Context) FormValue(key string) string {
	return ctx.Request.FormValue(key)
}

// AddHeader adds the key, value pair to the header.
func (ctx *Context) AddHeader(key, value string) {
	ctx.Response.Header.Add(key, value)
}

// SetHeader set the key, value pair to the header.
func (ctx *Context) SetHeader(key, value string) {
	ctx.Response.Header.Set(key, value)
}

// User return the current login user.
func (ctx *Context) User() interface{} {
	return ctx.Data["user"]
}

// App is the key struct of the package. App as a member of plugin
// entity contains the request and the corresponding handler. Prefix
// is the url prefix and MiddlewareList is for control flow.
type App struct {
	Requests    []Path
	tree        *node
	Middlewares Handlers
	Prefix      string
}

// NewApp return an empty app.
func NewApp() *App {
	return &App{
		Requests:    make([]Path, 0),
		tree:        Tree(),
		Prefix:      "/",
		Middlewares: make([]Handler, 0),
	}
}

type Handler func(ctx *Context)

type Handlers []Handler

// AppendReqAndResp stores the request info and handle into app.
// support the route parameter. The route parameter will be recognized as
// wildcard store into the RegUrl of Path struct. For example:
//
//         /user/:id      => /user/(.*)
//         /user/:id/info => /user/(.*?)/info
//
// The RegUrl will be used to recognize the incoming path and find
// the handler.
func (app *App) AppendReqAndResp(url, method string, handler []Handler) {

	app.Requests = append(app.Requests, Path{
		URL:    join(app.Prefix, url),
		Method: method,
	})

	app.tree.addPath(stringToArr(join(app.Prefix, slash(url))), method, append(app.Middlewares, handler...))
}

// Find is public helper method for findPath of tree.
func (app *App) Find(url, method string) []Handler {
	return app.tree.findPath(stringToArr(url), method)
}

// POST is a shortcut for app.AppendReqAndResp(url, "post", handler).
func (app *App) POST(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "post", handler)
}

// GET is a shortcut for app.AppendReqAndResp(url, "get", handler).
func (app *App) GET(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "get", handler)
}

// DELETE is a shortcut for app.AppendReqAndResp(url, "delete", handler).
func (app *App) DELETE(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "delete", handler)
}

// PUT is a shortcut for app.AppendReqAndResp(url, "put", handler).
func (app *App) PUT(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "put", handler)
}

// OPTIONS is a shortcut for app.AppendReqAndResp(url, "options", handler).
func (app *App) OPTIONS(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "options", handler)
}

// HEAD is a shortcut for app.AppendReqAndResp(url, "head", handler).
func (app *App) HEAD(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "head", handler)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, HEAD, OPTIONS, DELETE.
func (app *App) ANY(url string, handler ...Handler) {
	app.AppendReqAndResp(url, "post", handler)
	app.AppendReqAndResp(url, "get", handler)
	app.AppendReqAndResp(url, "delete", handler)
	app.AppendReqAndResp(url, "put", handler)
	app.AppendReqAndResp(url, "options", handler)
	app.AppendReqAndResp(url, "head", handler)
}

// TabGroups add middlewares and prefix for App.
func (app *App) Group(prefix string, middleware ...Handler) *RouterGroup {
	return &RouterGroup{
		app:         app,
		Middlewares: append(app.Middlewares, middleware...),
		Prefix:      slash(prefix),
	}
}

// RouterGroup is a group of routes.
type RouterGroup struct {
	app         *App
	Middlewares Handlers
	Prefix      string
}

// AppendReqAndResp stores the request info and handle into app.
// support the route parameter. The route parameter will be recognized as
// wildcard store into the RegUrl of Path struct. For example:
//
//         /user/:id      => /user/(.*)
//         /user/:id/info => /user/(.*?)/info
//
// The RegUrl will be used to recognize the incoming path and find
// the handler.
func (g *RouterGroup) AppendReqAndResp(url, method string, handler []Handler) {

	g.app.Requests = append(g.app.Requests, Path{
		URL:    join(g.Prefix, url),
		Method: method,
	})

	g.app.tree.addPath(stringToArr(join(g.Prefix, slash(url))), method, append(g.Middlewares, handler...))
}

// POST is a shortcut for app.AppendReqAndResp(url, "post", handler).
func (g *RouterGroup) POST(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "post", handler)
}

// GET is a shortcut for app.AppendReqAndResp(url, "get", handler).
func (g *RouterGroup) GET(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "get", handler)
}

// DELETE is a shortcut for app.AppendReqAndResp(url, "delete", handler).
func (g *RouterGroup) DELETE(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "delete", handler)
}

// PUT is a shortcut for app.AppendReqAndResp(url, "put", handler).
func (g *RouterGroup) PUT(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "put", handler)
}

// OPTIONS is a shortcut for app.AppendReqAndResp(url, "options", handler).
func (g *RouterGroup) OPTIONS(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "options", handler)
}

// HEAD is a shortcut for app.AppendReqAndResp(url, "head", handler).
func (g *RouterGroup) HEAD(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "head", handler)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, HEAD, OPTIONS, DELETE.
func (g *RouterGroup) ANY(url string, handler ...Handler) {
	g.AppendReqAndResp(url, "post", handler)
	g.AppendReqAndResp(url, "get", handler)
	g.AppendReqAndResp(url, "delete", handler)
	g.AppendReqAndResp(url, "put", handler)
	g.AppendReqAndResp(url, "options", handler)
	g.AppendReqAndResp(url, "head", handler)
}

// TabGroups add middlewares and prefix for App.
func (g *RouterGroup) Group(prefix string, middleware ...Handler) *RouterGroup {
	return &RouterGroup{
		app:         g.app,
		Middlewares: append(g.Middlewares, middleware...),
		Prefix:      join(slash(g.Prefix), slash(prefix)),
	}
}

// slash fix the path which has wrong format problem.
//
// 	 ""      => "/"
// 	 "abc/"  => "/abc"
// 	 "/abc/" => "/abc"
// 	 "/abc"  => "/abc"
// 	 "/"     => "/"
//
func slash(prefix string) string {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" || prefix == "/" {
		return "/"
	}
	if prefix[0] != '/' {
		if prefix[len(prefix)-1] == '/' {
			return "/" + prefix[:len(prefix)-1]
		}
		return "/" + prefix
	}
	if prefix[len(prefix)-1] == '/' {
		return prefix[:len(prefix)-1]
	}
	return prefix
}

// join join the path.
func join(prefix, suffix string) string {
	if prefix == "/" {
		return suffix
	}
	if suffix == "/" {
		return prefix
	}
	return prefix + suffix
}
