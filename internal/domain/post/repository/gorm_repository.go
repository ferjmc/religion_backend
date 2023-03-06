package repository

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/post"

	"gorm.io/gorm"
)

type postGormRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.PostRepository {
	return &postGormRepository{db: db}
}

func (p *postGormRepository) Create(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	res := p.db.Create(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}

func (p *postGormRepository) GetPosts(ctx context.Context, user *entities.User) ([]entities.Post, error) {
	var posts []entities.Post
	res := p.db.Where(&entities.Post{UserID: user.ID}).Find(&posts)
	if res.Error != nil {
		return nil, res.Error
	}
	return posts, nil
}

func (p *postGormRepository) GetGroupPosts(ctx context.Context, group []entities.Group) ([]entities.Post, error) {
	var posts []entities.Post
	var userIds []uint
	for _, g := range group {
		for _, u := range g.Users {
			userIds = append(userIds, u.ID)
		}
	}
	if err := p.db.Model(entities.Post{}).Where("UserId IN (?)", userIds).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
