package main

import (
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "echo\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":5000", nil)
}
