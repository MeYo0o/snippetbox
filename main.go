package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", homeHandler)
	mux.HandleFunc("/view", snippetView)
	mux.HandleFunc("/create", snippetCreate)

	log.Println("starting the server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
