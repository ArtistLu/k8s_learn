package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, strings.Join(v, " "))
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	w.Write([]byte("hello"))
}
