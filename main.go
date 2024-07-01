package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT    = "8080"
	DEFAULT_MESSAGE = "test message"
)

func main() {
	port := getPort()

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return DEFAULT_PORT
}

func getMessage() string {
	if message := os.Getenv("MESSAGE"); message != "" {
		return message
	}

	return DEFAULT_MESSAGE
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello "+getMessage())
}
