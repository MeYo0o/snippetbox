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
	handleRoutes(mux)

	fmt.Println("SnippetBox server started on :4000")

	// start the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
