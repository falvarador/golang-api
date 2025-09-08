package routes

import (
	"Gin/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

// UserRoutes sets up the routes for user-related operations.
// It takes a Gin RouterGroup and a UserHandler to bind the handlers to specific paths.
func UserRoutes(rg *gin.RouterGroup, userHandler *http.UserHandler) {
	users := rg.Group("/users") // Creates a /api/users group
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
