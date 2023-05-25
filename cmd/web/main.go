package main

import (
	"GoLangHTTP/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the server handler
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting Server on port%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
