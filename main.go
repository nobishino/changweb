package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello")
}

type WorldHandler struct{}

func (*WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World")
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello")
}

func world(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// 何も書かないとDefaultMuxが使われる
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	// if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
	// 	log.Fatal(err)
	// }
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
