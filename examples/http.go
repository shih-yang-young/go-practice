package examples

import (
	"fmt"
	"net/http"
)

// net/http 對於handler的接口定義
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// net/http 的http.HandlerFunc 自定義
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func GoHttp() {
	var handler http.Handler = HandlerFunc(helloHandler)

	http.ListenAndServe(":8080", handler)
}
