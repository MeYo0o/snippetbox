package main

import (
	"html/template"
	"log/slog"

	"github.com/MeYo0o/snippetbox/internal/models"
)

type Application struct {
	Logger        *slog.Logger
	Snippets      *models.SnippetModel
	TemplateCache map[string]*template.Template
}
