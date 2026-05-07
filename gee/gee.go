package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc // support middlewares
	parent *RouterGroup  // support nesting
	engine *Engine // all groups share a Engine instance
}

// 将engine作为最顶层的RouterGroup
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// Engine 创建工厂方法
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
