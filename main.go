package main

import (
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

// main is the server handler
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Server on port%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
