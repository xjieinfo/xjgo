package xjhttp

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Engine struct {
	routes []Route //map[string]http.HandlerFunc
}

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func Default() *Engine {
	return &Engine{
		routes: make([]Route, 0),
	}
}

func (engine *Engine) HandleFunc(method, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	route := Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}
	engine.routes = append(engine.routes, route)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range engine.routes {
		if r.Method == route.Method && r.URL.Path == route.Pattern {
			route.Handler(w, r)
			return
		}
	}
	mapHandle := make(map[int]http.HandlerFunc)
	for _, route := range engine.routes {
		if r.Method == route.Method {
			matchi := engine.PathMatch(r.URL.Path, route.Pattern)
			if matchi > 0 {
				mapHandle[matchi] = route.Handler
			}
		}
	}
	if len(mapHandle) == 0 {
		fmt.Fprint(w, "404 Page Not Found!")
		return
	}
	handler := engine.GetHandle(mapHandle)
	handler(w, r)
}

func (engine *Engine) GetHandle(mapHandle map[int]http.HandlerFunc) http.HandlerFunc {
	max := 0
	var handler http.HandlerFunc
	for k, v := range mapHandle {
		if k > max {
			max = k
			handler = v
		}
	}
	return handler
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

func (engine *Engine) Run(address string) (err error) {
	fmt.Printf("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}
