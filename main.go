package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/teml.html
var temp string

func process(w http.ResponseWriter, req *http.Request) {
	t := template.New("tmpl.html")
	t, err := t.Parse(temp)
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
