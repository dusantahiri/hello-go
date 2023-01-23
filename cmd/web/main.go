package main

import (
	"fmt"
	"net/http"

	"github.com/dusantahiri/hello-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Staring application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
