package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/teml.html")
	if err != nil {
		fmt.Fprintln(w, "Fail")
		return
	}
	t.Execute(w, "Hello, Template!")
}

func main() {
	http.HandleFunc("/", process)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
