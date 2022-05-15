/*
-------------------------------------
# @Time    : 2022/5/16 2:13:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : tinyweb.go
# @Software: GoLand
-------------------------------------
*/

package tinyweb

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc 定义框架的请求处理程序
type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc // 支持中间件
		engine      *Engine       // 通过 Engine 间接地访问各种接口
	}
	// Engine 框架的所有资源由 Engine 统一协调
	Engine struct {
		*RouterGroup  // 继承底层模块所拥有的能力
		router        *router
		groups        []*RouterGroup     // 存储所有路由组
		htmlTemplates *template.Template // 将所有的模板加载进内存
		funcMap       template.FuncMap   // 所有的自定义模板渲染函数
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Default 使用 Logger 和 Recovery 中间件
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// Group 用于创建一个新路由组，所有路由组共享同一个 Engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Use 用于向分组中添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler) // 实现了路由的映射
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// createStaticHandler 创建静态资源
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// 判断文件是否存在或者是否有权限
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// Static 提供静态文件
func (group *RouterGroup) Static(relativePath, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

// SetFuncMap 设置自定义渲染函数 funcMap
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// LoadHTMLGlob 加载模板
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 每接收到一个请求都需逐一匹配前缀来添加中间件
	// gin 是在前缀树的节点中添加中间件的切片，这样在匹配动态路由并解析参数时，就可以同时获得各分组的中间件。
	for _, group := range engine.groups {
		// 判断请求适用于哪些中间件
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c) // 使用底层模块提供的能力
}
