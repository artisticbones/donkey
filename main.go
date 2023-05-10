package main

import (
	"net/http"

	"github.com/artisticbones/donkey/donkey"
)

// Engine is the uni handler for all requests,just for first version
// type Engine struct{}

// func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	switch req.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
// 	case "/hello":
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		}
// 	default:
// 		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
// 	}
// }

// second version
func main() {
	r := donkey.New()
	// r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	// })

	// r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	// 	for k, v := range req.Header {
	// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// 	}
	// })
	r.GET("/", func(c *donkey.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Donkey</h1>")
	})

	r.GET("/hello", func(c *donkey.Context) {
		// expect /hello?name=xxx
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *donkey.Context) {
		c.JSON(http.StatusOK, donkey.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
