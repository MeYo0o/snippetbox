package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	resp := fmt.Sprintf("Display a specific snippet with id: (%d)", id)
	w.Write([]byte(resp))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", homeHandler)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("starting the server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
