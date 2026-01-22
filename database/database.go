package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect creates a pgsql connection and returns *sql.DB + error
func Connect(databaseURL, port, user, pass, dbName string) (*sql.DB, error) {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseURL, port, user, pass, dbName,
	)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
