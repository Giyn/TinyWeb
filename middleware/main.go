/*
-------------------------------------
# @Time    : 2022/5/11 20:37:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : main.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	tinygin "TinyGin/middleware/tiny-gin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() tinygin.HandlerFunc {
	return func(c *tinygin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := tinygin.New()
	r.Use(tinygin.Logger()) // global middleware
	r.GET("/", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGin</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *tinygin.Context) {
			// expect /hello/Giyn
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	r.Run(":9999")
}
