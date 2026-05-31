package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	log.Println("starting the server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
