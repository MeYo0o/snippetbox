package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// serving static files => css, img, scripts
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// stripping it from the "static" prefix before it reaches the server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// as short of
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snipperCreatePost)

	//* applying middleware
	return app.LogRequest(commonHeaders(mux))
}
