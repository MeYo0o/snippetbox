package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.innolabs.ai/internal/database"
	"snippetbox.innolabs.ai/internal/models"
)

type application struct {
	Logger   *slog.Logger
	Snippets *models.SnippetModel
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

	//> Database Setup
	pool, err := setupDB(logger)
	if err != nil {
		os.Exit(1)
	}
	defer pool.Close()
	queries := database.New(pool)

	//> define application struct that contains dependency injected features
	app := &application{
		Logger:   logger,
		Snippets: &models.SnippetModel{Queries: queries},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
