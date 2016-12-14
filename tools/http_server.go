package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(404)
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)

}
