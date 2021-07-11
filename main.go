package main

import (
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
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)

	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
