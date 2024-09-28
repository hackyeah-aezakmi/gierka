package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTablesQuery := `
		CREATE TABLE IF NOT EXISTS games (
			id TEXT PRIMARY KEY,
			data TEXT NOT NULL
		);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}
