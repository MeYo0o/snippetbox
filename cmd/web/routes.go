package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// serving static files => css, img, scripts
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// stripping it from the "static" prefix before it reaches the server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// as short of
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.Then(http.HandlerFunc(app.homeHandler)))
	mux.Handle("GET /snippet/view/{id}", dynamic.Then(http.HandlerFunc(app.snippetView)))
	mux.Handle("GET /snippet/create", dynamic.Then(http.HandlerFunc(app.snippetCreate)))
	mux.Handle("POST /snippet/create", dynamic.Then(http.HandlerFunc(app.snippetCreatePost)))

	standard := alice.New(app.recoverPanic, app.LogRequest, commonHeaders)

	//* applying middleware
	return standard.Then(mux)
}
