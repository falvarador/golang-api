// internal/core/services/user_service.go
package services

import (
	"Gin/internal/core/domain"
	"Gin/internal/core/ports"
	"Gin/pkg/util"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// UserService implements the UserDriverPort interface.
type UserService struct {
	userRepo ports.UserDrivenPort // Dependency on the Driven Port (Repository)
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepo ports.UserDrivenPort) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser implements the use case for creating a new user.
func (s *UserService) CreateUser(email, name string) (*domain.User, error) {
	user, err := domain.NewUser(email, name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, &util.InternalError{Message: "failed to check existing user by email", Err: err}
		}
	}

	user.ID = uuid.New().String() // Generate a unique ID
	// CreatedAt and UpdatedAt are set in domain.NewUser

	// Save the user using the repository (driven port)
	if err := s.userRepo.SaveUser(user); err != nil {
		return nil, &util.InternalError{Message: "failed to save user", Err: err}
	}

	return user, nil
}

// GetUserByID implements the use case for getting a user by ID.
func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	user, err := s.userRepo.FindUserByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &util.NotFoundError{Message: fmt.Sprintf("user with ID %s not found", id)}
		}

		return nil, &util.InternalError{Message: "failed to retrieve user from repository", Err: err}
	}

	if user == nil {
		return nil, &util.NotFoundError{Message: fmt.Sprintf("user with ID %s not found", id)}
	}

	return user, nil
}

// GetAllUsers implements the use case for getting all users.
func (s *UserService) GetAllUsers() ([]domain.User, error) {
	users, err := s.userRepo.FindAllUsers()

	if err != nil {
		return nil, &util.InternalError{Message: "failed to retrieve all users", Err: err}
	}

	return users, nil
}

// UpdateUser implements the use case for updating an existing user.
func (s *UserService) UpdateUser(id, email, name string) (*domain.User, error) {
	user, err := s.userRepo.FindUserByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &util.NotFoundError{Message: fmt.Sprintf("user with ID %s not found for update", id)}
		}

		return nil, &util.InternalError{Message: "failed to retrieve user for update from repository", Err: err}
	}

	if user == nil {
		return nil, &util.NotFoundError{Message: fmt.Sprintf("user with ID %s not found for update", id)}
	}

	// Update fields if provided
	if email != "" {
		user.Email = email
	}

	if name != "" {
		user.Name = name
	}

	user.UpdatedAt = time.Now() // Update timestamp

	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, &util.InternalError{Message: "failed to update user in repository", Err: err}
	}

	return user, nil
}

// DeleteUser implements the use case for deleting a user.
func (s *UserService) DeleteUser(id string) error {
	if err := s.userRepo.DeleteUser(id); err != nil {
		if errors.Is(err, errors.New("user not found for deletion")) { // Este error debe provenir del repo
			return &util.NotFoundError{Message: fmt.Sprintf("user with ID %s not found for deletion", id)}
		}

		return &util.InternalError{Message: "failed to delete user from repository", Err: err}
	}

	return nil
}
