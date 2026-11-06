package service

import (
	"saffronreal.estate/backend/internal/repository"
)

type UserService struct {
	users *repository.UserRepository
}

func NewDatabaseService(
	users *repository.UserRepository,
) *UserService {
	return &UserService{
		users: users,
	}
}
