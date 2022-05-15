/*
-------------------------------------
# @Time    : 2022/5/16 2:11:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : main.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"TinyGin/panic-recover/tinyweb"
	"net/http"
)

func main() {
	r := tinyweb.Default()
	r.GET("/", func(c *tinyweb.Context) {
		c.String(http.StatusOK, "Hello Giyn\n")
	})
	r.GET("/panic", func(c *tinyweb.Context) {
		names := []string{"Giyn"}
		c.String(http.StatusOK, names[7])
	})
	r.Run(":9999")
}
