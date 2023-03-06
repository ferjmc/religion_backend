package post

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/post/dto"
)

type UseCase interface {
	Create(ctx context.Context, uid string, input dto.CreatePostRequest) (*entities.Post, error)
	GetPosts(ctx context.Context, uid string) ([]entities.Post, error)
}
