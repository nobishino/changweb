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

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)

	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
