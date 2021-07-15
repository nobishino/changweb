package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//go:embed templates/*
var templates embed.FS

func formatDate(t time.Time) string {
	format := "20060102"
	return t.Format(format)
}

func process(w http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("time.html").Funcs(funcMap)
	t, err := t.ParseFS(templates, "templates/time.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "parse error:", err)
		return
	}
	t.Execute(w, time.Now())
}

func main() {
	http.HandleFunc("/", process)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
