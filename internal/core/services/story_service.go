package services

import (
	"Gin/internal/core/domain"
	"Gin/internal/core/ports"
	"Gin/pkg/util"
	"database/sql"

	"errors"
	"fmt"
)

// Implrsments the ports.StoryDrivingPort interface for StoryService.
type StoryService struct {
	repo ports.StoryDrivenPort
}

// Creates a new instance of StoryService.
func NewStoryService(repo ports.StoryDrivenPort) *StoryService {
	return &StoryService{repo: repo}
}

// Handles the creation of a new story.
func (s *StoryService) CreateStory(input *domain.NewStoryInput) (*domain.Story, error) {
	// Here you can perform additional validations or checks before saving the story.
	// For example, you can check if the title or content are empty or if the author is not empty.

	story := &domain.Story{
		Title:   input.Title,
		Author:  input.Author,
		Content: input.Content,
		// ID, CreatedAt, UpdatedAt are automatically set by the repository
	}

	if err := s.repo.SaveStory(story); err != nil {
		return nil, &util.InternalError{Message: "failed to save story", Err: err}
	}

	return story, nil
}

// Handles the retrieval of a story by ID.
func (s *StoryService) GetStoryByID(id string) (*domain.Story, error) {
	story, err := s.repo.FindStoryByID(id)

	if err != nil {
		// If the error is sql.ErrNoRows, it means the story was not found
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found", id)}
		}

		// Any other error is internal
		return nil, &util.InternalError{Message: "failed to retrieve story from repository", Err: err}
	}

	// Verify if the story was found
	if story == nil {
		return nil, &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found", id)}
	}

	return story, nil
}

// Handles the retrieval of all stories.
func (s *StoryService) GetAllStories() ([]domain.Story, error) {
	stories, err := s.repo.FindAllStories()

	if err != nil {
		return nil, &util.InternalError{Message: "failed to retrieve all stories", Err: err}
	}

	return stories, nil
}

// Handles the update of a story.
func (s *StoryService) UpdateStory(id string, input *domain.UpdateStoryInput) (*domain.Story, error) {
	// First, retrieve the story from the repository.
	story, err := s.repo.FindStoryByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found for update", id)}
		}

		return nil, &util.InternalError{Message: "failed to retrieve story for update from repository", Err: err}
	}

	if story == nil {
		return nil, &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found for update", id)}
	}

	// Apply the changes to the story
	if input.Title != nil {
		story.Title = *input.Title
	}

	if input.Author != nil {
		story.Author = *input.Author
	}

	if input.Content != nil {
		story.Content = *input.Content
	}

	// The updated_at column is automatically updated by the repository
	if err := s.repo.UpdateStory(story); err != nil {

		if errors.Is(err, errors.New("story not found or no changes made")) {
			return nil, &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found for update (or no changes)", id)}
		}

		return nil, &util.InternalError{Message: "failed to update story in repository", Err: err}
	}

	return story, nil
}

// Handles the deletion of a story.
func (s *StoryService) DeleteStory(id string) error {
	err := s.repo.DeleteStory(id)

	if err != nil {
		// Assume that DeleteStory from the repository can return a specific error if not found
		if errors.Is(err, errors.New("story not found for deletion")) {
			return &util.NotFoundError{Message: fmt.Sprintf("story with ID %s not found for deletion", id)}
		}

		return &util.InternalError{Message: "failed to delete story from repository", Err: err}
	}

	return nil
}
