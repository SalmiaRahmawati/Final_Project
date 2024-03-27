package service

import (
	"context"
	"errors"

	"my_gram/internal/model"
	"my_gram/internal/repository"
	"my_gram/pkg/helper"
)

type UserService interface {
	UserRegister(ctx context.Context, userSignUp model.UserSignUp) (user model.User, err error)
	UserLogin(ctx context.Context, userSignIn model.UserSignIn) (token string, err error)
	GetUsersById(ctx context.Context, id uint64) (model.User, error)
	// DeleteUsersById(ctx context.Context, id uint64) (model.User, error)

	// GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error)
}

type userServiceImpl struct {
	repo repository.UserQuery
}

func NewUserService(repo repository.UserQuery) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

func (u *userServiceImpl) UserRegister(ctx context.Context, userSignUp model.UserSignUp) (user model.User, err error) {
	user = model.User{
		Username: userSignUp.Username,
		Password: userSignUp.Password,
		Email:    userSignUp.Email,
		DoB:      userSignUp.DoB,
	}

	user, err = u.repo.CreateUser(ctx, user)
	return
}

func (u *userServiceImpl) UserLogin(ctx context.Context, userSignIn model.UserSignIn) (token string, err error) {
	user := model.User{
		Email:    userSignIn.Email,
		Password: userSignIn.Password,
	}

	user, err = u.repo.FindByEmail(ctx, user)
	if err != nil {
		err = errors.New("invalid email or password")
		return
	}

	if pass := helper.ComparePass([]byte(user.Password), []byte(userSignIn.Password)); !pass {
		err = errors.New("invalid email or password")
		return
	}

	token = helper.GenerateToken(user.ID, user.Email)
	return
}

func (u *userServiceImpl) GetUsersById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUsersByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
