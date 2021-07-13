package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/teml.html
var temp string

func process(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/t1.html", "./templates/t2.html"))
	// if err != nil {
	// 	fmt.Fprintln(w, "Fail", err)
	// 	return
	// }
	// b := rand.Intn(10) > 5
	t.Execute(w, "test-include-action")
}

func main() {
	http.HandleFunc("/", process)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
