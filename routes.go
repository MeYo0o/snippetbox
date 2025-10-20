package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", home)                          //Display the home page
	mux.HandleFunc("GET /snippet/create", snippetCreate)      //Display a form for creating a new snippet
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)     //Display a specific snippet
	mux.HandleFunc("POST /snippet/create", snippetCreatePost) //Create a new snippet
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from SnippetBox"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save an new snippet..."))
}
