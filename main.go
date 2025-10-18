package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// new serve mux router instance
	mux := http.NewServeMux()

	// handling all routes
	mux.HandleFunc("/{$}", Home)
	mux.HandleFunc("/snippet/create", SnippetCreate)
	mux.HandleFunc("/snippet/view", SnippetView)

	fmt.Println("SnippetBox server started on :4000")

	// start the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from SnippetBox"))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
