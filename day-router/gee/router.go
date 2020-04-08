package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	fmt.Println("parts:",parts)
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	//子节点进行insert操作
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	fmt.Println("进入路由")
	searchParts := parsePattern(path)
	params := make(map[string]string)
	fmt.Println("r",r)
	fmt.Println("method",method)
	root, ok := r.roots[method]
	fmt.Println("root children:",root.children)

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		fmt.Println("找到了这个节点n:",n)
		parts := parsePattern(n.pattern)
		fmt.Println("search parts:",searchParts)
		fmt.Println("parts",n.pattern)
		for index, part := range parts {
			fmt.Println("part[0]",part[0])
			if part[0] == ':' {
				fmt.Println(searchParts[index])
				fmt.Println(part[1:])
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	fmt.Println("接管路由的第一阶段")
	n, params := r.getRoute(c.Method, c.Path)
	fmt.Println("handle n:",n)
	fmt.Println("handle params:",params)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}