package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the server handler
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Server on port%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
