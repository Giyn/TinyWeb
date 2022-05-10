/*
-------------------------------------
# @Time    : 2022/5/10 9:42:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : router.go
# @Software: GoLand
-------------------------------------
*/

package tiny_gin

import (
	"log"
	"net/http"
)

// router 把路由独立出来
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
