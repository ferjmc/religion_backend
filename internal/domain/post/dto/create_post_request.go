package dto

import "religion/internal/domain/entities"

type CreatePostRequest struct {
	UserUid      string `validate:"required"`
	Description  string `validate:"required"`
	PostImageUrl string
}

func (r *CreatePostRequest) CreatePost() *entities.Post {
	return &entities.Post{
		Description:  r.Description,
		PostImageUrl: r.PostImageUrl,
	}
}
