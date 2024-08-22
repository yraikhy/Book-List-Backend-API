package application

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, *sql.DB, error) {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		return nil, db, fmt.Errorf("failed to open database: %w", err)
	}

	createBooksTableQuery := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		reading_status TEXT );`

	_, err = db.Exec(createBooksTableQuery)
	if err != nil {
		return nil, db, fmt.Errorf("failed to create books table: %w", err)
	}

	db2, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, db, fmt.Errorf("failed to open database: %w", err)
	}

	createUsersTableQuery := `CREATE TABLE IF NOT EXISTS users (
		userid INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT);`

	_, err = db2.Exec(createUsersTableQuery)
	if err != nil {
		return nil, db, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, db2, nil
}
