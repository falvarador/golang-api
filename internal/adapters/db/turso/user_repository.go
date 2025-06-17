package turso

import (
	"Gin/internal/core/domain"
	"Gin/pkg/util"

	"database/sql"
	"errors"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (id, email, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	// Format the dates to a string with the exact format you expect to read.
	// time.RFC3339Nano is a good ISO 8601 format with nanoseconds for SQLite TEXT.
	createdAtStr := util.FormatTimeToString(user.CreatedAt)
	updatedAtStr := util.FormatTimeToString(user.UpdatedAt)

	_, err := r.db.Exec(query, user.ID, user.Email, user.Name, createdAtStr, updatedAtStr)

	if err != nil {
		return nil, fmt.Errorf("turso: failed to insert user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	user := &domain.User{}

	// Scan as string (TEXT in SQLite)
	var createdAtStr, updatedAtStr string

	err := row.Scan(&user.ID, &user.Email, &user.Name, &createdAtStr, &updatedAtStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found
			return nil, nil
		}

		// It's important to report the exact error for debugging
		return nil, fmt.Errorf("turso: failed to find user by ID (scan error): %w", err)
	}

	// Now, parse the strings to time.Time.
	// If the format is not exactly RFC3339Nano, this is the point where it will fail.
	parsedCreatedAt, err := util.ParseTimeFromString(createdAtStr)
	if err != nil {
		// The error of ParseTimeFromString is already descriptive
		return nil, fmt.Errorf("turso: %w", err)
	}

	parsedUpdatedAt, err := util.ParseTimeFromString(updatedAtStr)
	if err != nil {
		// The error of ParseTimeFromString is already descriptive
		return nil, fmt.Errorf("turso: %w", err)
	}

	user.CreatedAt = parsedCreatedAt
	user.UpdatedAt = parsedUpdatedAt

	return user, nil
}
