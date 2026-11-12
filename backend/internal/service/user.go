package service

import (
	"context"
	"errors"
	"strings"

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
	if !util.ValidateEmail(email) {
		return model.User{}, ErrInvalidEmail
	}

	email = strings.ToLower(email)

	return s.users.NewUser(ctx, email)
}

func (s *UserService) Get(ctx context.Context, email string) (model.User, error) {
	if !util.ValidateEmail(email) {
		return model.User{}, ErrInvalidEmail
	}

	email = strings.ToLower(email)

	return s.users.GetUser(ctx, email)
}
