package xjhttp

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Engine struct {
	routes      []Route //map[string]http.HandlerFunc
	static      map[string]string
	redirect    map[string]string
	useHandlers []HandlerFunc //拦截的处理函数
}

type Route struct {
	Method   string
	Pattern  string
	Handlers []HandlerFunc
	Power    string
}

func Default() *Engine {
	return &Engine{
		routes:      make([]Route, 0),
		static:      make(map[string]string),
		redirect:    make(map[string]string),
		useHandlers: make([]HandlerFunc, 0),
	}
}

type HandlerFunc func(*Context)

func (engine *Engine) HandleFunc(method, pattern, power string, handler func(*Context)) {
	//_handler := func(w http.ResponseWriter, r *http.Request) {
	//	handler(&Context{
	//		Request: r,
	//		Writer:  w,
	//	})
	//}
	_handlers := make([]HandlerFunc, 0)
	_handlers = append(_handlers, engine.useHandlers...)
	_handlers = append(_handlers, handler)
	route := Route{
		Method:   method,
		Pattern:  pattern,
		Handlers: _handlers,
		Power:    power,
	}
	engine.routes = append(engine.routes, route)
}

func (engine *Engine) POST(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodPost, pattern, power, handler)
}

func (engine *Engine) GET(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodGet, pattern, power, handler)
}

func (engine *Engine) DELETE(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodDelete, pattern, power, handler)
}

func (engine *Engine) PATCH(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodPatch, pattern, power, handler)
}

func (engine *Engine) PUT(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodPut, pattern, power, handler)
}

func (engine *Engine) OPTIONS(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodOptions, pattern, power, handler)
}

func (engine *Engine) HEAD(pattern, power string, handler func(*Context)) {
	engine.HandleFunc(http.MethodHead, pattern, power, handler)
}

// Use attaches a global middleware to the router.
func (engine *Engine) Use(handler func(*Context)) {
	engine.useHandlers = append(engine.useHandlers, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//跳转
	for k, v := range engine.redirect {
		if r.URL.Path == k {
			log.Printf("%s redirect to: %s \n", k, v)
			//r.URL.Path = v
			//break
			w.Header().Set("Cache-Control", "must-revalidate, no-store")
			w.Header().Set("Content-Type", " text/html;charset=UTF-8")
			w.Header().Set("Location", v) //跳转地址设置
			w.WriteHeader(307)            //关键在这里！
			return
		}
	}
	//静态目录匹配
	for k, v := range engine.static {
		if strings.Index(r.URL.Path, k) == 0 {
			log.Printf("%s of static: %s \n", r.URL.Path, k)
			fileServer := http.StripPrefix(k, http.FileServer(http.Dir(v)))
			fileServer.ServeHTTP(w, r)
			return
		}
	}
	//完全匹配
	for _, route := range engine.routes {
		if r.Method == route.Method && r.URL.Path == route.Pattern {
			log.Printf("method: %s, path: %s \n", route.Method, r.RequestURI)
			r.Header.Set("xjgo-traceid", uuid.New().String())
			r.Header.Set("xjgo-power", route.Power)
			route.Handlers[0](&Context{Request: r, Writer: w, handlers: route.Handlers})
			return
		}
	}
	//路径参数匹配
	mapRoute := make(map[int]Route)
	for _, route := range engine.routes {
		if r.Method == route.Method {
			matchi := engine.PathMatch(r.URL.Path, route.Pattern)
			if matchi > 0 {
				mapRoute[matchi] = route
			}
		}
	}
	if len(mapRoute) == 0 {
		w.WriteHeader(404)
		//fmt.Fprint(w, "404, Page Not Found!")
		log.Printf("404,method: %s, path: %s \n", r.Method, r.RequestURI)
		return
	}
	route := engine.GetRoute(mapRoute)
	r.Header.Set("xjgo-path-pattern", route.Pattern)
	r.Header.Set("xjgo-traceid", uuid.New().String())
	log.Printf("method: %s, path: %s, pattern: %s \n", route.Method, r.RequestURI, route.Pattern)
	r.Header.Set("xjgo-power", route.Power)
	route.Handlers[0](&Context{Request: r, Writer: w, handlers: route.Handlers})
}

func (engine *Engine) GetRoute(mapRoute map[int]Route) Route {
	max := 0
	var route Route
	for k, v := range mapRoute {
		if k > max {
			max = k
			route = v
		}
	}
	return route
}

func (engine *Engine) PathMatch(path, pattern string) int {
	paths := strings.Split(path, "/")
	patterns := strings.Split(pattern, "/")
	if len(paths) != len(patterns) {
		return 0
	}
	match := ""
	for i, p1 := range paths {
		if p1 == patterns[i] {
			match += "2"
		} else if len(patterns[i]) > 0 && patterns[i][0:1] == ":" {
			match += "1"
		} else {
			return 0
		}
	}
	matchi, _ := strconv.Atoi(match)
	return matchi
}

func (engine *Engine) Static(relativePath, root string) {
	engine.static[relativePath] = root
}

func (engine *Engine) Redirect(relativePath, root string) {
	engine.redirect[relativePath] = root
}

func (engine *Engine) Run(address string) (err error) {
	fmt.Printf("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}
