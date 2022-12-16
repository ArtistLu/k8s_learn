package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/health", Health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Health(w http.ResponseWriter, r *http.Request) {
	defer func() {
		fmt.Printf("httpCode: %s ip:%s\n", http.StatusOK, getClientIP(r))
	}()
	w.WriteHeader(200)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	defer func() {
		fmt.Printf("httpCode: %d ip:%s\n", http.StatusOK, getClientIP(r))
	}()
	for k, v := range r.Header {
		w.Header().Add(k, strings.Join(v, " "))
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	w.Write([]byte("hello"))
}

func getClientIP(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP != "" {
		fips := strings.Split(clientIP, ",")
		if len(fips) > 0 {
			clientIP = strings.TrimSpace(fips[0])
		}
	}
	if clientIP == "" {
		clientIP = r.Header.Get("X-Real-Ip")
		rips := strings.Split(clientIP, ",")
		if len(rips) > 0 {
			clientIP = strings.TrimSpace(rips[0])
		}
	}

	clientIP, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		return clientIP
	}

	return ""
}
