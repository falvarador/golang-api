package domain

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(email, name string) (*User, error) {

	if email == "" {
		return nil, errors.New("email is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	return &User{
		ID:        "",
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
