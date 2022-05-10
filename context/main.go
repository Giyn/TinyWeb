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
	tinygin "TinyGin/context/tiny-gin"
	"net/http"
)

func main() {
	r := tinygin.New()
	r.GET("/", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Giyn</h1>")
	})
	r.GET("/hello", func(c *tinygin.Context) {
		// expect /hello?name=Giyn
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *tinygin.Context) {
		c.JSON(http.StatusOK, tinygin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
