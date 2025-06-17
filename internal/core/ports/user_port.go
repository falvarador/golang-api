package ports

import (
	"Gin/internal/core/domain"
)

type UserDriverPort interface {
	CreateUser(email, name string) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
}

type UserDrivenPort interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
}
