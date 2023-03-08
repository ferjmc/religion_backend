package usecase

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/post"
	"religion/internal/domain/post/dto"
	"religion/internal/domain/user"
	"religion/pkg/logger"
)

type postUseCase struct {
	repo post.PostRepository
	user user.UserGormRepository
	log  logger.Logger
}

func NewPostUsecase(repo post.PostRepository, user user.UserGormRepository, log logger.Logger) post.UseCase {
	return &postUseCase{
		repo: repo,
		user: user,
		log:  log,
	}
}

func (u *postUseCase) getUser(ctx context.Context, uid string) (*entities.User, error) {
	return u.user.GetSingleUser(ctx, uid)
}

func (u *postUseCase) Create(ctx context.Context, uid string, input dto.CreatePostRequest) (*entities.Post, error) {
	u.log.Infof("Create Post with uid: %s", uid)
	user, err := u.getUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	post := input.CreatePost()
	post.UserID = user.ID

	post, err = u.repo.Create(ctx, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *postUseCase) GetPosts(ctx context.Context, uid string) ([]entities.Post, error) {
	u.log.Infof("Get Posts with uid: %v", uid)
	user, err := u.getUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return u.repo.GetGroupPosts(ctx, user.Groups)
}
