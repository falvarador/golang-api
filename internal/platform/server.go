package platform

import (
	"Gin/internal/adapters/db/turso"
	"Gin/internal/adapters/http"
	"Gin/internal/core/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// InitGinServer configures and returns an instance of Gin Engine.
func InitGinServer(db *sql.DB) *gin.Engine {

	// Initialize the architecture hexagonal
	userRepository := turso.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	// Initialize the Gin server with logger and recovery by default
	r := gin.Default()

	// Configure routes
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUserByID)
		}
	}

	// Routes to serve the frontend of React/Astro (later)
	// r.Static("/", "./frontend/dist")

	return r
}
