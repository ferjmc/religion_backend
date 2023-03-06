package post

import (
	"context"
	"religion/internal/domain/entities"
)

type PostRepository interface {
	Create(ctx context.Context, post *entities.Post) (*entities.Post, error)
	GetPosts(ctx context.Context, user *entities.User) ([]entities.Post, error)
	GetGroupPosts(ctx context.Context, group []entities.Group) ([]entities.Post, error)
}
