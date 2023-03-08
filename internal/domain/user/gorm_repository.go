package user

import (
	"context"
	"religion/internal/domain/entities"
)

type UserGormRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetSingleUser(ctx context.Context, uid string) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	CreateGroup(ctx context.Context, group *entities.Group) (*entities.Group, error)
}
