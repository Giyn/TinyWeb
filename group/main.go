/*
-------------------------------------
# @Time    : 2022/5/11 19:18:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : main.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"TinyGin/group/tinyweb"
	"net/http"
)

func main() {
	r := tinyweb.New()
	r.GET("/index", func(c *tinyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *tinyweb.Context) {
			c.HTML(http.StatusOK, "<h1>Hello TinyWeb</h1>")
		})
		v1.GET("/hello", func(c *tinyweb.Context) {
			// expect /hello?name=Giyn
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tinyweb.Context) {
			// expect /hello/Giyn
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *tinyweb.Context) {
			c.JSON(http.StatusOK, tinyweb.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":9999")
}
