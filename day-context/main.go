package main

import (
	"gee/day-context/gee"
	"net/http"
)

func main()  {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK,"<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK,"hello %s,you're at %s\n",c.Query("name"),c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"ws": c.PostForm("ws"),
		})
	})

	r.Run(":9999")

}

// curl "http://localhost:9999/hello?name=geektutu"
// curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
// curl "http://localhost:9999/xxx"