package config

import (
	"log/slog"

	"github.com/MeYo0o/snippetbox/internal/models"
)

type Application struct {
	Logger   *slog.Logger
	Snippets *models.SnippetModel
}
