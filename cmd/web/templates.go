package main

import "snippetbox.innolabs.ai/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
