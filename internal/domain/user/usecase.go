package user

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/user/dto"
)

type UseCase interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*entities.User, error)
	UpdateUser(ctx context.Context, req *dto.UpdateUserRequest) (*entities.User, error)
	GetSingleUser(ctx context.Context, uid string) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
}
