package routes

import (
	"Gin/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

// Manages the routes for story-related operations.
func StoryRoutes(rg *gin.RouterGroup, storyHandler *http.StoryHandler) {
	stories := rg.Group("/stories")
	{
		stories.POST("", storyHandler.CreateStory)
		stories.GET("/:id", storyHandler.GetStory)
		stories.GET("", storyHandler.GetAllStories)
		stories.PUT("/:id", storyHandler.UpdateStory) // <-- PUT is used for partial updates
		stories.DELETE("/:id", storyHandler.DeleteStory)
	}
}
