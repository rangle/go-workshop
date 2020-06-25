package main

import (
	"fmt"
	"log"
	"net/http"
)

// Function to handle a single request
func handler(w http.ResponseWriter, r *http.Request) {
	// `fmt.Fprintf` formats according to a format specifier and writes to `w`
	fmt.Fprintf(w, "Hi there, you are at %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server Starting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
