package repository

import (
	"context"

	"my_gram/internal/infrastructure"
	"my_gram/internal/model"

	"gorm.io/gorm"
)

type UserQuery interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	FindByEmail(ctx context.Context, user model.User) (model.User, error)
	GetUsersByID(ctx context.Context, id uint64) (model.User, error)

	// DeleteUsersByID(ctx context.Context, id uint64) error
}

type userQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewUserQuery(db infrastructure.GormPostgres) UserQuery {
	return &userQueryImpl{
		db: db,
	}
}

func (u *userQueryImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	err := u.db.GetConnection().Create(&user).Error
	return user, err
}

func (u *userQueryImpl) FindByEmail(ctx context.Context, user model.User) (model.User, error) {
	err := u.db.GetConnection().Where("email = ?", user.Email).Take(&user).Error
	return user, err
}

func (u *userQueryImpl) GetUsersByID(ctx context.Context, id uint64) (model.User, error) {
	users := model.User{}
	err := u.db.GetConnection().Where("id = ?", id).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, nil
	}
	return model.User{}, err
}
