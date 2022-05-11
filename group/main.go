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
	tinygin "TinyGin/group/tiny-gin"
	"net/http"
)

func main() {
	r := tinygin.New()
	r.GET("/index", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *tinygin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello TinyGin</h1>")
		})
		v1.GET("/hello", func(c *tinygin.Context) {
			// expect /hello?name=Giyn
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tinygin.Context) {
			// expect /hello/Giyn
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *tinygin.Context) {
			c.JSON(http.StatusOK, tinygin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":9999")
}
