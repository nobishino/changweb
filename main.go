package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	c, err := req.Cookie("2nd_cookie")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, c)
	fmt.Fprintln(w, req.Cookies())
}

func setMessage(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "flash",
		Value:    base64.URLEncoding.EncodeToString([]byte("Hello Message")),
		HttpOnly: true,
	})
	fmt.Fprintln(w, "Set Message")
}

func showMessage(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("flash")
	if err != nil {
		fmt.Fprintln(w, "No Message Is Set")
	} else {
		reset := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Time{},
		}
		http.SetCookie(w, &reset)
		msg, err := base64.URLEncoding.DecodeString(c.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "failed to decode message")
		}
		fmt.Fprintln(w, string(msg))
	}
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
	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/show", showMessage)

	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
