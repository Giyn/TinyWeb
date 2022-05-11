/*
-------------------------------------
# @Time    : 2022/5/10 23:22:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : tiny-gin.go
# @Software: GoLand
-------------------------------------
*/

package tiny_gin

import (
	"log"
	"net/http"
)

// HandlerFunc 定义框架的请求处理程序
type HandlerFunc func(*Context)

// Engine 实现了ServeHTTP接口
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
