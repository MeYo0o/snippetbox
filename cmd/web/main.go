package main

import (
	"crypto/tls"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5/stdlib"
	"snippetbox.innolabs.ai/internal/database"
	"snippetbox.innolabs.ai/internal/models"
)

type application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	user           *models.UserModel
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	//> Commandline Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	certFile := flag.String("cert-file", "tls/localhost.crt", "HTTPS certificate file")
	keyFile := flag.String("key-file", "tls/localhost.key", "HTTPS private key file")
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

	//> Form Decoder: for Parse Forms and assigning them to the pre-configured fields.
	formDecoder := form.NewDecoder()

	//> Session Manager
	sessionDB := stdlib.OpenDBFromPool(pool)
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(sessionDB)
	sessionManager.Lifetime = time.Hour * 12

	//> define application struct that contains dependency injected features
	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{Queries: queries},
		user:           &models.UserModel{Queries: queries},
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	//> Server Configuration
	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{
				tls.X25519, tls.CurveP256,
			},
		},
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", *addr)
	logger.Info("starting server", "certFile", *certFile)
	logger.Info("starting server", "keyFile", *keyFile)

	err = srv.ListenAndServeTLS(*certFile, *keyFile)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
