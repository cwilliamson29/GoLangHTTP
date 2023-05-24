package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addvalues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about apge and value: %d", sum))
}

// addvalues is adding 2 integers and returns sum
func addvalues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the server handler
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Server on port%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
