package models

import (
	"context"

	"snippetbox.innolabs.ai/internal/database"
)

type SnippetModel struct {
	Queries *database.Queries
}

func (m *SnippetModel) Insert(title, content string, expires int) (int32, error) {
	return m.Queries.CreateSnippet(context.Background(), database.CreateSnippetParams{
		Title:   title,
		Content: content,
		Column3: expires,
	})
}

func (m *SnippetModel) Get(id int32) (database.Snippet, error) {
	return m.Queries.GetSnippet(context.Background(), id)
}

func (m *SnippetModel) Latest() ([]database.Snippet, error) {
	return nil, nil
}
