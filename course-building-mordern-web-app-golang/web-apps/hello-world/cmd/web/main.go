package main

import (
	"github.com/joaodiasdev/hellowordwebapp/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
