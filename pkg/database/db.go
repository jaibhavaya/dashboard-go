package database

import (
	"database/sql"
	"fmt"
	"log"
)

func setupDatabase(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Check if we shoud create sample data
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	// Add sample data if the table is empty
	if count == 0 {
		_, err = db.Exec("INSERT INTO users (id, name) VALUES ('1', 'Alice'), ('2', 'Bob')")
		if err != nil {
			return err
		}
	}

	return nil
}

func InitializeDatabase() (*sql.DB, error) {
	connStr := "postgres://admin:password@localhost:5432/dashboard?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	log.Println("Connected to PostgreSQL database")

	if err := setupDatabase(db); err != nil {
		return nil, fmt.Errorf("error setting up database: %w", err)
	}

	return db, nil
}
