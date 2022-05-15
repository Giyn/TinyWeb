/*
-------------------------------------
# @Time    : 2022/5/16 2:13:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : logger.go
# @Software: GoLand
-------------------------------------
*/

package tinyweb

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// start timer
		t := time.Now()
		// process request
		c.Next() // 控制权交给下一个中间件
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
