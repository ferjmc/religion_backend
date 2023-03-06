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
	return u.repo.Create(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, req *dto.UpdateUserRequest) (*entities.User, error) {
	u.log.Infof("updating user with: %v", req)
	user := req.UpdateUser()
	return u.repo.Update(ctx, user)
}

func (u *userUseCase) GetSingleUser(ctx context.Context, uid string) (*entities.User, error) {
	u.log.Infof("get single user with uid: %s", uid)
	return u.repo.GetSingle(ctx, uid)
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	u.log.Infof("get all users")
	return u.repo.GetAll(ctx)
}
