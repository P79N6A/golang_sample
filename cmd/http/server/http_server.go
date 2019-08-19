package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// HelloHandler http handler
type HelloHandler struct {
	welcome string
}

// ServeHttp is http.Handler implementation
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name, ok := values["name"]
	if ok {
		fmt.Fprintf(w, "%v, %v", h.welcome, name)
	} else {
		fmt.Fprintf(w, h.welcome)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.Handle("/hello/handler", &HelloHandler{"Hello World"})
	http.HandleFunc("/hello/func", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
