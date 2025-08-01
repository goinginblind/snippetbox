package models

import (
	"database/sql"
	"time"
)

// Snippet holds the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps around the database connection
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the db
func (sm *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// Get a snippet by its id
func (sm *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Returns some amount of the latest snippets
func (sm *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
