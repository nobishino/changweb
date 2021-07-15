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

func process2(w http.ResponseWriter, req *http.Request) {
	t, _ := template.New("template.html").ParseFS(templates, "templates/template.html")
	t.Execute(w, req.FormValue("comment"))
}

func form(w http.ResponseWriter, req *http.Request) {
	t, _ := template.New("form.html").ParseFS(templates, "templates/form.html")
	t.Execute(w, nil)
}

func main() {
	static := http.FileServer(http.FS(static))
	http.Handle("/", static)
	http.HandleFunc("/process", process2)
	http.HandleFunc("/form", form)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
