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

func (group *RouterGroup) GET(pattern, power string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodGet, power, group.group+pattern, handler)
}

func (group *RouterGroup) DELETE(pattern, power string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodDelete, power, group.group+pattern, handler)
}
