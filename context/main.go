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
	"TinyGin/context/tinyweb"
	"net/http"
)

func main() {
	r := tinyweb.New()
	r.GET("/", func(c *tinyweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyWeb</h1>")
	})
	r.GET("/hello", func(c *tinyweb.Context) {
		// expect /hello?name=Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *tinyweb.Context) {
		c.JSON(http.StatusOK, tinyweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
