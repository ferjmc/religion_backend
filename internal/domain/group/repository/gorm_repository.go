package repository

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/group"

	"gorm.io/gorm"
)

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) group.Repository {
	return &groupRepository{db: db}
}

func (g *groupRepository) Create(ctx context.Context, group *entities.Group) (*entities.Group, error) {
	if err := g.db.Create(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
