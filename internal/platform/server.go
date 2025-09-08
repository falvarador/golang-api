package platform

import (
	"Gin/internal/platform/middlewares"
	"Gin/internal/platform/routes"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// InitGinServer configures and returns a Gin Engine instance.
func InitGinServer(db *sql.DB) *gin.Engine {
	// Initialize the hexagonal architecture components
	container := SetupContainer(db)

	app := gin.Default() // Gin with default logger and recovery middleware

	// Apply global middlewares
	app.Use(middlewares.CORSMiddleware()) // Use your centralized CORS middleware here

	// Setup API routes group
	api := app.Group("/api")
	{
		// Register user routes using the new routes package
		routes.UserRoutes(api, container.UserHandler)
		routes.StoryRoutes(api, container.StoryHandler)
	}

	// Routes to serve React/Astro frontend (later)
	// r.Static("/", "./frontend/dist") // Example for React SPA

	return app
}
