package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snipperCreatePost)

	log.Println("starting the server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
