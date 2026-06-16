package main

import (
	"net/http"
)

func (conf *Config) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// serving static files => css, img, scripts
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// stripping it from the "static" prefix before it reaches the server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// as short of
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	mux.HandleFunc("GET /{$}", conf.homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", conf.snippetView)
	mux.HandleFunc("GET /snippet/create", conf.snippetCreate)
	mux.HandleFunc("POST /snippet/create", conf.snipperCreatePost)

	return mux
}
