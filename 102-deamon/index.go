package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	x := os.Getenv("HTTPS")
	//x := r.Header.Get("X-TRACE-ID")
	//x := r.Header
	if len(x) == 0 {
		fmt.Println("nil")
	} else {
		fmt.Println("ARUYAN")
		fmt.Printf("%#v\n", x)
	}
	fmt.Fprintf(w, "TraceID  %s", x)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	http.HandleFunc("/", handler)
	fcgi.Serve(l, nil)
}
