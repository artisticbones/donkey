package donkey

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	routers map[string]HandlerFunc // todo: finding why use map structure
}

func New() *Engine {
	return &Engine{
		routers: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.routers[key] = handler
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.routers[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// GET defines the method to add GET request
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter(http.MethodGet, pattern, handler)
}

// POST defines the method to add POST request
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter(http.MethodPost, pattern, handler)
}
