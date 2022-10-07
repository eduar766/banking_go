package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomer)

	// starting Server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}