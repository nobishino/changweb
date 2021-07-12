package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(
		writer,
		request.Header["Accept-Encoding"],
		request.Header.Get("Accept-Encoding"),
	)
}

func body(writer http.ResponseWriter, request *http.Request) {
	len := request.ContentLength
	fmt.Fprintf(writer, "BODY. Len=%d\n", len)
	buf := make([]byte, len)
	for {
		n, err := request.Body.Read(buf)
		fmt.Fprintf(writer, string(buf[:n]))
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintln(writer, err)
			return
		}
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("uploaded")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	msg := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello, Go Web!</h1></body>
</html>`
	if _, err := w.Write([]byte(msg)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "no service here")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(http.StatusFound)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	post := Post{
		User: "user",
		Threads: []string{
			"red",
			"blue",
		},
	}
	buf, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "failed to marshal post json")
		return
	}

	if _, err := w.Write(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "failed to write response")
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "name",
		Value:    "nobishii",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "2nd_cookie",
		Value:    "nobishino",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	w.Header().Set("x-test", "test")
}

func getCookie(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, req.Header["Cookie"])
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", writeHeaderExample)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/google", headerExample)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
