package dto

import "religion/internal/domain/entities"

type GroupRequest struct {
	UserUid string `validate:"required"`
	Name    string
}

func (r *GroupRequest) CreateGroup(admin uint) *entities.Group {
	return &entities.Group{
		Admin: admin,
		Name:  r.Name,
	}
}
