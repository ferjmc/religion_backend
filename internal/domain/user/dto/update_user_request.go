package dto

import (
	"religion/internal/domain/entities"

	"gorm.io/gorm"
)

type UpdateUserRequest struct {
	Id         uint `validate:"required,gt:0"`
	Uid        string
	Username   string
	Name       string
	Bio        string
	Website    string
	Email      string
	ProfileUrl string
}

func (r *UpdateUserRequest) UpdateUser() *entities.User {
	return &entities.User{
		Model:      gorm.Model{ID: r.Id},
		Uid:        r.Uid,
		Username:   r.Username,
		Name:       r.Name,
		Bio:        r.Bio,
		Website:    r.Website,
		Email:      r.Email,
		ProfileUrl: r.ProfileUrl,
	}
}
