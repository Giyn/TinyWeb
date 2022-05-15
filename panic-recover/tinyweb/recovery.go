/*
-------------------------------------
# @Time    : 2022/5/16 2:12:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : recovery.go
# @Software: GoLand
-------------------------------------
*/

package tinyweb

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// trace 获取触发 panic 的堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	// Callers 用来返回调用栈的程序计数器
	n := runtime.Callers(3, pcs[:]) // 跳过前3个调用者(第0个是本身,第1个是上层trace,第2个是再上层的defer func)
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)   // 获取对应的函数
		file, line := fn.FileLine(pc) // 获取到调用该函数的文件名和行号
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// Recovery 错误恢复中间件
func Recovery() HandlerFunc {
	return func(c *Context) {
		// 使用 defer 挂载上错误恢复的函数
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
