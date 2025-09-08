package postgresql

import (
	"Gin/internal/core/domain"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Implements the ports.UserDrivenPort interface for PostgreSQL.
type UserRepository struct {
	db *sql.DB
}

// Creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Implements the logic to save a user to PostgreSQL.
func (r *UserRepository) SaveUser(user *domain.User) error {
	// PostgreSQL uses $1, $2, etc., for placeholders instead of ?.
	// Also, TIMESTAMPTZ (with timezone) is a common type.
	query := `INSERT INTO users (id, email, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	// PostgreSQL's `pq` driver and `database/sql` can often handle `time.Time` directly
	// without needing to convert to string first, assuming your DB column is `TIMESTAMP WITH TIME ZONE`.
	// However, if using `TEXT` columns for timestamps, you'd still need util.FormatTimeToString.
	// For standard TIMESTAMP WITH TIME ZONE in Postgres, direct time.Time is preferred.
	_, err := r.db.Exec(query, user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("postgresql: failed to insert user: %w", err)
	}
	return nil
}

// Implements the logic to find a user by ID in PostgreSQL.
func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = $1` // Placeholder $1
	row := r.db.QueryRow(query, id)

	user := &domain.User{}
	// Direct scan into time.Time for TIMESTAMP WITH TIME ZONE columns
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("postgresql: failed to find user by ID (scan error): %w", err)
	}
	return user, nil
}

// Implements the logic to find all users in PostgreSQL.
func (r *UserRepository) FindAllUsers() ([]domain.User, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("postgresql: failed to query all users: %w", err)
	}
	defer rows.Close()

	users := make([]domain.User, 0)

	for rows.Next() {
		user := &domain.User{}
		// Direct scan into time.Time
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("postgresql: failed to scan user row: %w", err)
		}
		users = append(users, *user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("postgresql: rows iteration error: %w", err)
	}
	return users, nil
}

// Implements the logic to update an existing user in PostgreSQL.
func (r *UserRepository) UpdateUser(user *domain.User) error {
	query := `UPDATE users SET email = $1, name = $2, updated_at = $3 WHERE id = $4` // Placeholders $1, $2, $3, $4
	result, err := r.db.Exec(query, user.Email, user.Name, user.UpdatedAt, user.ID)  // Direct time.Time
	if err != nil {
		return fmt.Errorf("postgresql: failed to update user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found or no changes made")
	}
	return nil
}

// Implements the logic to delete a user from PostgreSQL.
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1` // Placeholder $1
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("postgresql: failed to delete user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found for deletion")
	}
	return nil
}
