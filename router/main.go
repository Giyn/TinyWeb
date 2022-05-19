/*
-------------------------------------
# @Time    : 2022/5/11 1:07:38
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
	r.GET("/hello/:name", func(c *Context) {
		// expect /hello/Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *Context) {
		c.JSON(http.StatusOK, H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
