/*
-------------------------------------
# @Time    : 2022/5/16 2:12:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : context.go
# @Software: GoLand
-------------------------------------
*/

package tinyweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 封装 Request 和 Response
type Context struct {
	// 原生对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求字段
	Path   string
	Method string
	Params map[string]string // 提供对路由参数的访问
	// 响应字段
	StatusCode int
	// 中间件
	handlers []HandlerFunc
	index    int // 记录当前执行到第几个中间件
	// engine 指针
	engine *Engine // 能够通过 Context 访问 Engine 中的 HTML 模板
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	// 保证中间件只执行一次
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 访问PostForm参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 访问Query参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String 快速构造String响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 快速构造JSON响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 快速构造Data响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML html 模板渲染
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
