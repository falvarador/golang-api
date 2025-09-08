package http

import (
	"Gin/internal/core/domain"
	"Gin/internal/core/ports"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Implements the ports.StoryDrivingPort interface for StoryHandler.
type StoryHandler struct {
	storyService ports.StoryDrivingPort // The handler uses the service interface
	validate     *validator.Validate    // Instance of the validator
}

// Creates a new instance of StoryHandler.
func NewStoryHandler(storyService ports.StoryDrivingPort) *StoryHandler {
	return &StoryHandler{
		storyService: storyService,
		validate:     validator.New(), // Inicializa el validador
	}
}

// CreateStory godoc
// @Summary Create a new story
// @Description Creates a new story with the provided title, author, and content.
// @Tags stories
// @Accept json
// @Produce json
// @Param story body domain.NewStoryInput true "Story creation object"
// @Success 201 {object} domain.Story
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /stories [post]
func (h *StoryHandler) CreateStory(c *gin.Context) {
	var input domain.NewStoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Validate the input using go-playground/validator
	if err := h.validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors.Error()})
		return
	}

	story, err := h.storyService.CreateStory(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create story", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, story)
}

// GetStory godoc
// @Summary Get a story by ID
// @Description Retrieves a single story by its unique ID.
// @Tags stories
// @Produce json
// @Param id path string true "Story ID"
// @Success 200 {object} domain.Story
// @Failure 404 {object} map[string]string "Story not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /stories/{id} [get]
func (h *StoryHandler) GetStory(c *gin.Context) {
	id := c.Param("id")
	story, err := h.storyService.GetStoryByID(id)

	if err != nil {
		// Check if the error is because the story was not found.
		if errors.Is(err, errors.New("story not found")) { // <-- Asegúrate de que el mensaje de error coincida con el del servicio.
			c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve story", "details": err.Error()})
		return
	}
	if story == nil { // This also covers the case of "not found" if the service returns nil, nil
		c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
		return
	}

	c.JSON(http.StatusOK, story)
}

// GetAllStories godoc
// @Summary Get all stories
// @Description Retrieves a list of all stories.
// @Tags stories
// @Produce json
// @Success 200 {array} domain.Story
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /stories [get]
func (h *StoryHandler) GetAllStories(c *gin.Context) {
	stories, err := h.storyService.GetAllStories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stories", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stories)
}

// UpdateStory godoc
// @Summary Update an existing story
// @Description Updates an existing story identified by ID with the provided fields.
// @Tags stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Param story body domain.UpdateStoryInput true "Story update object"
// @Success 200 {object} domain.Story
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Story not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /stories/{id} [put]
func (h *StoryHandler) UpdateStory(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdateStoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Validation for UpdateStoryInput (omitempty on the tags validates only the ones coming)
	if err := h.validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors.Error()})
		return
	}

	story, err := h.storyService.UpdateStory(id, &input)

	if err != nil {
		if errors.Is(err, errors.New("story not found for update")) { // <-- Asegúrate de que el mensaje de error coincida con el del servicio.
			c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update story", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, story)
}

// DeleteStory godoc
// @Summary Delete a story by ID
// @Description Deletes a story by its unique ID.
// @Tags stories
// @Produce json
// @Param id path string true "Story ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string "Story not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /stories/{id} [delete]
func (h *StoryHandler) DeleteStory(c *gin.Context) {
	id := c.Param("id")
	err := h.storyService.DeleteStory(id)

	if err != nil {
		if errors.Is(err, errors.New("story not found for deletion")) { // <-- Asegúrate de que el mensaje de error coincida con el del servicio.
			c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete story", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content for successful deletion
}
