package gee

/*
框架入口
*/

import (
	"log"
	"net/http"
)

// 定义请求处理函数 handler 参数为 Context
type HandlerFunc func(*Context)

// 定义 Engine
type Engine struct {
	router *router
}

// 构造 Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 调用 router.addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

// 调用 engine.addRoute 添加 GET 路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 调用 engine.addRoute 添加 POST 路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 运行 http server 监听 addr 所有请求由 engine 实例处理
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 所有请求由 ServeHTTP 处理，将请求和响应封装为 Context 类型，调用 router.handle 处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
