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
	group.engine.HandleFunc(http.MethodGet, group.group+pattern, power, handler)
}

func (group *RouterGroup) DELETE(pattern, power string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodDelete, group.group+pattern, power, handler)
}

func (group *RouterGroup) POST(pattern, power string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodPost, group.group+pattern, power, handler)
}

func (group *RouterGroup) PUT(pattern, power string, handler func(*Context)) {
	group.engine.HandleFunc(http.MethodPut, group.group+pattern, power, handler)
}
