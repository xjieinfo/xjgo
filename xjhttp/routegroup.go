package xjhttp

import "net/http"

type RouterGroup struct {
	engine *Engine
	group  string
}

func (engine *Engine) Group(group string) *RouterGroup {
	return &RouterGroup{
		engine: engine,
		group:  group,
	}
}

func (group *RouterGroup) GET(pattern string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodGet, group.group+pattern, handler)
}

func (group *RouterGroup) DELETE(pattern string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodDelete, group.group+pattern, handler)
}
