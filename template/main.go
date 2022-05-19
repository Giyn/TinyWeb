/*
-------------------------------------
# @Time    : 2022/5/14 1:17:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : main.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "TinyWeb"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := New()
	r.Use(Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Giyn", Age: 22}
	stu2 := &student{Name: "Jack", Age: 20}
	r.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *Context) {
		c.HTML(http.StatusOK, "arr.tmpl", H{
			"title":  "TinyWeb",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", H{
			"title": "TinyWeb",
			"now":   time.Date(2022, 5, 14, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":9999")
}
