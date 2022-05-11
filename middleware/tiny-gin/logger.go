/*
-------------------------------------
# @Time    : 2022/5/11 20:38:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : logger.go
# @Software: GoLand
-------------------------------------
*/

package tiny_gin

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
