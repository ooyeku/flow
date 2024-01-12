package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		log.Fatalf("error writing response: %s", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
