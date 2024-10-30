package models

import (
	"database/sql"
	"log/slog"
	"time"
)

// Snippet is a type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel is a type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES( $1, $2, NOW(), NOW() + $3 * INTERVAL '1 day') RETURNING id`

	m.Logger.Info(stmt)

	var id int

	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		m.Logger.Error(err.Error())
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
