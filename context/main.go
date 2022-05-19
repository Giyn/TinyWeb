/*
-------------------------------------
# @Time    : 2022/5/9 23:52:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : main.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "TinyWeb"
	"net/http"
)

func main() {
	r := New()
	r.GET("/", func(c *Context) {
		c.String(http.StatusOK, "<h1>Hello TinyWeb</h1>")
	})
	r.GET("/hello", func(c *Context) {
		// expect /hello?name=Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
