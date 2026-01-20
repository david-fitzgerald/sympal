package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dataDir := filepath.Join(homeDir, ".sympal")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	dbPath := filepath.Join(dataDir, "data.db")
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	return createTables()
}

func createTables() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			due_date TEXT,
			priority TEXT CHECK(priority IN ('high', 'medium', 'low')),
			status TEXT DEFAULT 'not_started' CHECK(status IN ('not_started', 'done')),
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}
