package gee

/*
路由文件，负责路由添加和处理请求
*/

import (
	"net/http"
)

// 定义路由映射表
type router struct {
	handlers map[string]HandlerFunc
}

// 构造路由映射表
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 添加路由，key 由请求方法和静态路由地址构成，value 是用户映射的处理方法
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 定义请求处理逻辑
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
