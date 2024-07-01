package main

import (
	"net/http"
	"os"

	api "github.com/mauricioacevedo/go_vercel/api"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	port := getPort()

	http.HandleFunc("/hello", api.HelloHandler)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return DEFAULT_PORT
}
