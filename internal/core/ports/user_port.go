package ports

import (
	"Gin/internal/core/domain"
)

// UserDriverPort (or Application Service Port)
// Defines the operations that the Core exposes to external adapters (HTTP, CLI, etc.).
type UserDriverPort interface {
	CreateUser(email, name string) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)                     // New: List all users
	UpdateUser(id, email, name string) (*domain.User, error) // New: Update an existing user
	DeleteUser(id string) error                              // New: Delete a user
}

// UserDrivenPort (or Repository Port)
// Defines the operations that the Core needs from infrastructure adapters (DB, external services).
type UserDrivenPort interface {
	SaveUser(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
	FindAllUsers() ([]domain.User, error) // New: Find all users
	UpdateUser(user *domain.User) error   // New: Update user in DB
	DeleteUser(id string) error           // New: Delete user from DB
}
