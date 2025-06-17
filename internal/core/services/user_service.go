package services

import (
	"Gin/internal/core/domain"
	"Gin/internal/core/ports"
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository ports.UserDrivenPort
}

func NewUserService(userRepository ports.UserDrivenPort) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(email, name string) (*domain.User, error) {
	user, err := domain.NewUser(email, name)

	if err != nil {
		return nil, err
	}

	user.ID = uuid.NewString()

	if _, err := s.userRepository.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user: " + err.Error())
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found: " + err.Error())
	}

	return user, nil
}
