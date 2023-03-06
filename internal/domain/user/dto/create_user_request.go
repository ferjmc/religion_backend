package dto

import "religion/internal/domain/entities"

type CreateUserRequest struct {
	Uid      string `validate:"required"`
	Email    string `validate:"required,email"`
	Username string
	Bio      string
	Name     string
}

func (r *CreateUserRequest) CreateUser() *entities.User {
	return &entities.User{
		Uid:      r.Uid,
		Email:    r.Email,
		Username: r.Username,
		Bio:      r.Bio,
		Name:     r.Name,
	}
}
