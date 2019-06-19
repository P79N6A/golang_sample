package main

import (
	"fmt"
	"net/http"
	"time"
)

// HelloHandler http handler
type HelloHandler struct {
	welcome string
}

// ServeHttp is http.Handler implementation
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.welcome)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.Handle("/hello/handler", &HelloHandler{"Hello World"})
	http.HandleFunc("/hello/func", welcome)
	http.ListenAndServe(":8080", nil)
}
