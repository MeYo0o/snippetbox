package main

import (
	"html/template"
	"log/slog"

	"github.com/MeYo0o/snippetbox/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type application struct {
	Logger         *slog.Logger
	TemplateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
	// models
	snippets *models.SnippetModel
	users    *models.UserModel
}

type TLS struct {
	cert string
	key  string
}
