package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/MeYo0o/snippetbox/internal/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		// Level:     slog.LevelDebug,
	}))

	app := &config.Application{
		Logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, routes(app))
	logger.Error(err.Error())
	os.Exit(1)
}
