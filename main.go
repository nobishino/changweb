package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World, %s!", request.URL.Path[1:])
}

func main() {
	http.ListenAndServe("", nil)
}
