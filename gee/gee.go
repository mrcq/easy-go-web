package gee

import (
	"fmt"
	"net/http"
)

// 定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 定义路由映射表
type Engine struct {
	router map[string]HandlerFunc
}

// 返回初始化对象
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// key 由请求方法和静态路由地址构成，value 是用户映射的处理方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// 定义静态路由地址 pattern 的 GET 方法（key）和处理方法（val）到路由映射表 router（map）中
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 定义静态路由地址 pattern 的 POST 方法（key）和处理方法（val）到路由映射表 router（map）中
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 运行 http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ListenAndServe第二个参数非nil，所以engine需要实现 ServerHTTP 接口，所有请求由ServeHTTP处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
