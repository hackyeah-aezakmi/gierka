package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB    *sql.DB
	Close func() error
}

func InitDB(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTablesQuery := `
		CREATE TABLE IF NOT EXISTS games (
			id TEXT PRIMARY KEY,
			data TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			data TEXT NOT NULL
		);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB:    db,
		Close: db.Close,
	}, nil
}
