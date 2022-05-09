/*
-------------------------------------
# @Time    : 2022/5/9 14:55:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : tiny-gin.go
# @Software: GoLand
-------------------------------------
*/

package tiny_gin

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义框架的请求处理程序
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现 ServeHTTP 的接口
type Engine struct {
	router map[string]HandlerFunc // 路由映射表
}

// New 构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 用户注册静态路由
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 添加 GET 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加 POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 封装启动 http 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // engine 实现了 ServeHTTP 方法
}

// Handler 是一个接口，需要实现方法 ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
