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
	"TinyGin/router/tinyweb"
	"net/http"
)

func main() {
	r := tinyweb.New()
	r.GET("/", func(c *tinyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGin</h1>")
	})
	r.GET("/hello", func(c *tinyweb.Context) {
		// expect /hello?name=Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *tinyweb.Context) {
		// expect /hello/Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *tinyweb.Context) {
		c.JSON(http.StatusOK, tinyweb.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
