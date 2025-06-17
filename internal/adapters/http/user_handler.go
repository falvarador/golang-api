package http

import (
	"Gin/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler is a primary adapter that handles HTTP requests related to users.
type UserHandler struct {
	// Dependency on the Driver Port (Application Service)
	userService ports.UserDriverPort
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService ports.UserDriverPort) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUserRequest represents the body of the request to create a user.
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

// CreateUser godoc.
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data to create"
// @Success 201 {object} domain.User
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc.
// @Summary Get user by ID
// @Description Get a specific user by its ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} gin.H "User not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		// The service may return nil if it doesn't find the user
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
