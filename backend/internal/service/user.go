package service

import (
	"context"
	"errors"
	"fmt"

	"saffronreal.estate/backend/internal/model"
	"saffronreal.estate/backend/internal/repository"
	"saffronreal.estate/backend/internal/util"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type UserService struct {
	users *repository.UserRepository
}

func NewUserService(
	users *repository.UserRepository,
) *UserService {
	return &UserService{
		users: users,
	}
}

func (s *UserService) Create(ctx context.Context, email string) (model.User, error) {
	err := util.ValidateEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("error creating user: %v", err)
	}

	return s.users.NewUser(ctx, email)
}

func (s *UserService) Get(ctx context.Context, email string) (model.User, error) {
	err := util.ValidateEmail(email)
	if err != nil {
		return model.User{}, ErrInvalidEmail
	}

	return s.users.GetUser(ctx, email)
}
