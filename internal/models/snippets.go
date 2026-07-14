package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"snippetbox.innolabs.ai/internal/database"
)

// ###################### Models ##########################
type Snippet struct {
	ID      int32
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	Queries *database.Queries
}

// ###################### DB ##########################
func (m *SnippetModel) Insert(title, content string, expires int) (int32, error) {
	return m.Queries.CreateSnippet(context.Background(), database.CreateSnippetParams{
		Title:   title,
		Content: content,
		Column3: expires,
	})

}

func (m *SnippetModel) Get(id int32) (Snippet, error) {
	dbSnippet, err := m.Queries.GetSnippet(context.Background(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return Snippet(dbSnippet), nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	dbSnippets, err := m.Queries.GetLatestSnippets(context.Background())
	if err != nil {
		return nil, err
	}

	return dbSnippetsToSnippets(dbSnippets), nil
}

// #################### Helpers #######################
func dbSnippetsToSnippets(dbSnippets []database.Snippet) []Snippet {
	snippets := make([]Snippet, len(dbSnippets))

	for i, dbSnippet := range dbSnippets {
		snippets[i] = Snippet(dbSnippet)
	}

	return snippets
}
