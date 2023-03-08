package repository

import (
	"context"
	"religion/internal/domain/entities"
	"religion/internal/domain/user"

	"gorm.io/gorm"
)

type userGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserGormRepository {
	return &userGormRepository{db: db}
}

func (u *userGormRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userGormRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	res := u.db.Model(&user).Updates(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (u *userGormRepository) GetSingleUser(ctx context.Context, uid string) (*entities.User, error) {
	var user *entities.User
	res := u.db.Model(entities.User{}).Joins("Groups").Where(&entities.User{Uid: uid}).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (u *userGormRepository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	res := u.db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (g *userGormRepository) CreateGroup(ctx context.Context, group *entities.Group) (*entities.Group, error) {
	if err := g.db.Create(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
