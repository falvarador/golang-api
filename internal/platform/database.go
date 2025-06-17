package platform

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// InitTursoDB sets up the connection to the Turso database.
func InitTursoDB() (*sql.DB, error) {

	// Make sure you have these environment variables configured
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	env := os.Getenv("ENVIRONMENT")

	if dbURL == "" && (authToken == "" && env == "production") {
		return nil, errors.New("TURSO_DATABASE_URL, TURSO_AUTH_TOKEN and must be set in production")
	}

	// The connection URL for Turso with libsql-client-go
	// Example: "libsql://your-db-name-your-org.turso.io?authToken=YOUR_AUTH_TOKEN"
	dsn := fmt.Sprintf("%s?authToken=%s", dbURL, authToken)

	// Use the "libsql" driver
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open Turso DB connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {

		// Close the connection if the connection fails
		db.Close()
		return nil, fmt.Errorf("failed to connect to Turso DB: %w", err)
	}

	fmt.Println("Successfully connected to Turso database!")

	// Optional: Create the 'users' table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}
	fmt.Println("Users table checked/created.")

	return db, nil
}
