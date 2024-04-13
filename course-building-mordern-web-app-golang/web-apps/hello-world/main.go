package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 = %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)

}

func divideValues(x, y float32) (float32, error) {
	var result float32
	if y <= 0 {
		return result, errors.New("cannot divide by zero or less than zero")
	}
	result = x / y
	return result, nil
}

// addValues adds two integers and returns the result
func addValues(x, y int) int {
	return x + y
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
