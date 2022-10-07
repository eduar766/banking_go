package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)


	// starting Server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
