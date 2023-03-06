package user

import (
	"context"
	"religion/internal/domain/entities"
)

type UserGormRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	GetSingle(ctx context.Context, uid string) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
}
