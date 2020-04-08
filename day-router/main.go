package main

import (
	"gee/day-router/gee"
	"net/http"
)

func main() {
	r := gee.New()
	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//
	//r.GET("/hello", func(c *gee.Context) {
	//	// expect /hello?name=geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})

	//r.GET("/hello/:name", func(c *gee.Context) {
	//	// expect /hello/geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

//1.gee包中包含添加 addRoute注册对路由 =》 处理的funcHandler ,以及开始监听服务,并且实现ServeHTTP方法，接管请求的操作,其中engine包含router结构体
//2.其中router中包含 roots map和handlers map,其中key都是以 method+'-'+pattern，其中roots中存放着关于node结构体的节点
//3.在trie.go中定义node的结构体，具体有insert、search、matchChild、matchChildren


//2.




