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
	"TinyWeb"
	"net/http"
)

func main() {
	r := TinyWeb.Default()
	r.GET("/", func(c *TinyWeb.Context) {
		c.String(http.StatusOK, "Hello Giyn\n")
	})
	r.GET("/panic", func(c *TinyWeb.Context) {
		names := []string{"Giyn"}
		c.String(http.StatusOK, names[7])
	})
	r.Run(":9999")
}
