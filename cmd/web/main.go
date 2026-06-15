package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	//>	Commandline Flags
	// define the flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// initialize parsing them
	flag.Parse()

	//> Structured Logger
	// define the handler
	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	// initialize the logger
	logger := slog.New(loggerHandler)

	//> define application struct that contains dependency injected features
	app := &application{
		logger: logger,
	}

	//> app routes
	mux := app.routes()

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
