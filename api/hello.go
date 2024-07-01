package handler

import (
	"fmt"
	"net/http"
	"os"
)

const (
	DEFAULT_MESSAGE = "test message"
)

func getMessage() string {
	if message := os.Getenv("MESSAGE"); message != "" {
		return message
	}

	return DEFAULT_MESSAGE
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello "+getMessage())
}
