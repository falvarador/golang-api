package ports

import "Gin/internal/core/domain"

// This is the interface that the repository will use to interact with the database.
type StoryDrivenPort interface {
	SaveStory(story *domain.Story) error
	FindStoryByID(id string) (*domain.Story, error)
	FindAllStories() ([]domain.Story, error)
	UpdateStory(story *domain.Story) error
	DeleteStory(id string) error
}

// This is the interface that the handler will use to interact with the service.
type StoryDrivingPort interface {
	CreateStory(input *domain.NewStoryInput) (*domain.Story, error)
	GetStoryByID(id string) (*domain.Story, error)
	GetAllStories() ([]domain.Story, error)
	UpdateStory(id string, input *domain.UpdateStoryInput) (*domain.Story, error)
	DeleteStory(id string) error
}
