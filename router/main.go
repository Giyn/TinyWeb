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
	tinygin "TinyGin/router/tiny-gin"
	"net/http"
)

func main() {
	r := tinygin.New()
	r.GET("/", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGin</h1>")
	})
	r.GET("/hello", func(c *tinygin.Context) {
		// expect /hello?name=Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *tinygin.Context) {
		// expect /hello/Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *tinygin.Context) {
		c.JSON(http.StatusOK, tinygin.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
