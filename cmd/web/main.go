package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	//>	Commandline Flags
	// define the flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// initiate parsing them
	flag.Parse()

	//> Structured Logger
	// define the handler
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(loggerHandler)

	// serving static files => css, img, scripts
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// stripping it from the "static" prefix before it reaches the server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// as short of
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snipperCreatePost)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
