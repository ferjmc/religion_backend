package usecase

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/user"
	"religion/internal/domain/user/dto"
	"religion/pkg/logger"
)

type userUseCase struct {
	repo user.UserGormRepository
	log  logger.Logger
}

func New(repo user.UserGormRepository, log logger.Logger) user.UseCase {
	return &userUseCase{repo: repo, log: log}
}

func (u *userUseCase) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*entities.User, error) {
	u.log.Infof("creating user with: %v", req)
	user := req.CreateUser()
	return u.repo.CreateUser(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, req *dto.UpdateUserRequest) (*entities.User, error) {
	u.log.Infof("updating user with: %v", req)
	user := req.UpdateUser()
	return u.repo.UpdateUser(ctx, user)
}

func (u *userUseCase) GetSingleUser(ctx context.Context, uid string) (*entities.User, error) {
	u.log.Infof("get single user with uid: %s", uid)
	return u.repo.GetSingleUser(ctx, uid)
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	u.log.Infof("get all users")
	return u.repo.GetAllUsers(ctx)
}

func (u *userUseCase) CreateGroup(ctx context.Context, req *dto.GroupRequest) (*entities.Group, error) {
	u.log.Infof("create group with name: %s", req.Name)
	user, err := u.repo.GetSingleUser(ctx, req.UserUid)
	if err != nil {
		return nil, err
	}
	greq := req.CreateGroup(user.ID)
	return u.repo.CreateGroup(ctx, greq)
}
