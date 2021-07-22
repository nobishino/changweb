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

//go:embed client.html
var static embed.FS

func formatDate(t time.Time) string {
	format := "20060102"
	return t.Format(format)
}

func process(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFS(templates, "templates/layout.html", "templates/content.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	static := http.FileServer(http.FS(static))
	http.Handle("/", static)
	http.HandleFunc("/process", process)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
