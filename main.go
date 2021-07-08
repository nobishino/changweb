package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index(w, r)
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World, %s!", request.URL.Path[1:])
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: new(MyHandler),
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
