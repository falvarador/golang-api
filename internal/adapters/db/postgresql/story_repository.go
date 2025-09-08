package postgresql

import (
	"Gin/internal/core/domain"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Implements the ports.StoryDrivenPort interface for PostgreSQL.
type StoryRepository struct {
	db *sql.DB
}

// Creates a new instance of StoryRepository.
func NewStoryRepository(db *sql.DB) *StoryRepository {
	return &StoryRepository{db: db}
}

// Implements the logic to save a story in PostgreSQL.
func (r *StoryRepository) SaveStory(story *domain.Story) error {
	// Generate a new UUID if no ID is provided (for new stories)
	if story.ID == "" {
		story.ID = uuid.New().String()
	}

	story.CreatedAt = time.Now()
	story.UpdatedAt = time.Now()

	query := `INSERT INTO stories (id, title, author, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, story.ID, story.Title, story.Author, story.Content, story.CreatedAt, story.UpdatedAt)

	if err != nil {
		return fmt.Errorf("postgresql: failed to insert story: %w", err)
	}

	return nil
}

// Implements the logic to find a story by ID in PostgreSQL.
func (r *StoryRepository) FindStoryByID(id string) (*domain.Story, error) {
	query := `SELECT id, title, author, content, created_at, updated_at FROM stories WHERE id = $1`
	row := r.db.QueryRow(query, id)

	story := &domain.Story{}
	err := row.Scan(&story.ID, &story.Title, &story.Author, &story.Content, &story.CreatedAt, &story.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Story not found
		}

		return nil, fmt.Errorf("postgresql: failed to find story by ID (scan error): %w", err)
	}

	return story, nil
}

// Implements the logic to find all stories in PostgreSQL.
func (r *StoryRepository) FindAllStories() ([]domain.Story, error) {
	query := `SELECT id, title, author, content, created_at, updated_at FROM stories ORDER BY created_at DESC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("postgresql: failed to query all stories: %w", err)
	}

	defer rows.Close()

	stories := make([]domain.Story, 0)

	for rows.Next() {
		story := &domain.Story{}
		err := rows.Scan(&story.ID, &story.Title, &story.Author, &story.Content, &story.CreatedAt, &story.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("postgresql: failed to scan story row: %w", err)
		}

		stories = append(stories, *story)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("postgresql: rows iteration error: %w", err)
	}

	return stories, nil
}

// Implements the logic to update a story in PostgreSQL.
func (r *StoryRepository) UpdateStory(story *domain.Story) error {
	story.UpdatedAt = time.Now() // Update the updated_at column

	query := `UPDATE stories SET title = $1, author = $2, content = $3, updated_at = $4 WHERE id = $5`
	result, err := r.db.Exec(query, story.Title, story.Author, story.Content, story.UpdatedAt, story.ID)

	if err != nil {
		return fmt.Errorf("postgresql: failed to update story: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("story not found or no changes made")
	}

	return nil
}

// Implements the logic to delete a story in PostgreSQL.
func (r *StoryRepository) DeleteStory(id string) error {
	query := `DELETE FROM stories WHERE id = $1`
	result, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("postgresql: failed to delete story: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("story not found for deletion")
	}

	return nil
}
