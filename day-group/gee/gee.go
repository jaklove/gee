package gee

import (
	"fmt"
	"log"
	"net/http"
)

type (
	RouterGroup struct {
		prefix string
		middlewares []HandlerFunc
		parent *RouterGroup
		engine *Engine
	}

	// Engine implement the interface of ServeHTTP
    Engine struct {
    	*RouterGroup
        router *router
    	groups []*RouterGroup //store all groups
    }
)


// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// New is the constructor of gee.Engine
func New() *Engine {
	//初始化一个engine
	engine := &Engine{router:newRouter()}
	//初始化一个routerGroup,并且把engine带入
	engine.RouterGroup = &RouterGroup{engine:engine}
	//然后把当前engine的groups交给
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
	//return &Engine{router: newRouter()}
}


// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup)Group(prefix string)*RouterGroup  {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix:group.prefix + prefix,
		parent:group,
		engine:engine,
	}
	fmt.Println("prefix:",newGroup.prefix)
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp  string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup)  POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}