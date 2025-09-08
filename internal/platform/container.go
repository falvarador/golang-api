package platform

import (
	"Gin/internal/adapters/db/postgresql"
	"Gin/internal/adapters/http"
	"Gin/internal/core/services"

	"database/sql"
)

// Represents the container for the application.
type Container struct {
	UserHandler  *http.UserHandler
	StoryHandler *http.StoryHandler
}

// Creates a new instance of Container.
func SetupContainer(db *sql.DB) *Container {

	// Repositories are used to interact with the database.
	userRepo := postgresql.NewUserRepository(db)
	storyRepo := postgresql.NewStoryRepository(db)

	// Services are used to interact with the domain.
	userService := services.NewUserService(userRepo)
	storyService := services.NewStoryService(storyRepo)

	// Adapters are used to interact with the ports.
	userHandler := http.NewUserHandler(userService)
	storyHandler := http.NewStoryHandler(storyService)

	return &Container{
		UserHandler:  userHandler,
		StoryHandler: storyHandler,
	}
}
