// internal/adapters/http/user_handler.go
package http

import (
	"Gin/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler is a primary adapter that handles HTTP requests related to users.
type UserHandler struct {
	userService ports.UserDriverPort // Dependency on the Driver Port (Application Service)
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService ports.UserDriverPort) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUserRequest represents the request body for creating a user.
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

// UpdateUserRequest represents the request body for updating a user.
// Fields are pointers to allow partial updates (nil if not provided).
type UpdateUserRequest struct {
	Email *string `json:"email" binding:"omitempty,email"` // omitempty allows field to be missing or empty
	Name  *string `json:"name" binding:"omitempty"`
}

// CreateUser godoc
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

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a specific user by their ID
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
		// Differentiate between "not found" and other errors
		if err.Error() == "user not found: user not found" { // Be careful with string comparison of errors, better to use custom error types
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Note: UserService returns (nil, nil) if not found after FindUserByID in repo,
	//       so the error check above should catch that. If user is nil here, it's already an error case.

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all registered users
// @Tags users
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user's email or name by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "User data to update"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 404 {object} gin.H "User not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract values from pointers, providing empty string if nil
	email := ""
	if req.Email != nil {
		email = *req.Email
	}
	name := ""
	if req.Name != nil {
		name = *req.Name
	}

	user, err := h.userService.UpdateUser(id, email, name)
	if err != nil {
		if err.Error() == "user not found" { // Again, better to use custom error types
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} gin.H "User not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	err := h.userService.DeleteUser(id)
	if err != nil {
		if err.Error() == "failed to delete user: user not found for deletion" { // Better to use custom error types
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent) // 204 No Content for successful deletion
}
