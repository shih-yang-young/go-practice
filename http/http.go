package http

import (
	"fmt"
	"net/http"
)

// 定義一個簡單的 handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Print("aaaaaaa")
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a custom handler!")
	fmt.Print("bbbbbb")
}

func GoHttp() {
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", &Handler{})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
