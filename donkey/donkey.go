package donkey

import (
	"net/http"
)

// first version handler func
// type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// second version handler func
// 1. 对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但是这两个对象提供的接口粒度太细，
// 比如我们要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body)，而 Header 包含了状态码(StatusCode)，消息类型(ContentType)等几乎每次请求都需要设置的信息。
// 因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。针对常用场景，能够高效地构造出 HTTP 响应是一个好的框架必须考虑的点。
// 2. 针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。
// 对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？
// 再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。
// 因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。
type HandlerFunc func(context *Context)

// Engine implement the interface of ServeHTTP
// todo: need to consider multi program' problem
type Engine struct {
	routers *router
}

func New() *Engine {
	return &Engine{
		routers: newRouter(),
	}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.routers.addRouter(method, pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// key := req.Method + "-" + req.URL.Path
	// if handler, ok := e.routers[key]; ok {
	// 	handler(w, req)
	// } else {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	// }
	context := newContext(w, req)
	e.routers.handle(context)
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
