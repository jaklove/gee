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
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

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


//addRoute生命周期:
//首先对自定义的路由的解析，形成例如[hello,:gee]这种切片，method+'-'+ pattern作为key,r.roots是map结构，如果当前r.roots[key]是为空，则
//初始化node节点到r.roots[key] = &node{},判断当前是否有节点，如果没有节点，则进行子节点插入，初始高度为0，

//在trie.go中insert周期中，如果len(parts) == height，说明当前到达最后的节点，把当前的绑定的url赋值给子节点的pattern。

//接口请求服务器的时候，由于router中handle接管了当前http请求，根据path对应匹配的节点，根据节点返回对应的pattern，然后组成对应的key = method+'-'+pattern,取出
//对应的handlerFunc，然后去执行对应的操作





