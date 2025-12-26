package main

import (
	"html/template"
	"log/slog"

	"github.com/MeYo0o/snippetbox/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type Application struct {
	Logger         *slog.Logger
	Snippets       *models.SnippetModel
	TemplateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}
