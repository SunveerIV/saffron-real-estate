package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/model"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(
	pool *pgxpool.Pool,
) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) NewUser(ctx context.Context, email string) (model.User, error) {
	emailExists, err := r.EmailExists(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	if emailExists {
		return model.User{}, fmt.Errorf("user already exists")
	}

	const createUserQuery = `
		INSERT INTO users(email) VALUES ($1)
	`

	_, err = r.pool.Exec(ctx, createUserQuery, email)
	if err != nil {
		return model.User{}, err
	}

	return r.userFromEmail(ctx, email)
}

func (r *UserRepository) GetUser(ctx context.Context, email string) (model.User, error) {
	emailExists, err := r.EmailExists(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	if !emailExists {
		return model.User{}, fmt.Errorf("user does not exist")
	}

	return r.userFromEmail(ctx, email)
}

func (r *UserRepository) userFromEmail(ctx context.Context, email string) (model.User, error) {
	const userFromEmailQuery = `
		SELECT * FROM users WHERE email = $1 LIMIT 1
	`
	var user model.User

	err := r.pool.QueryRow(ctx, userFromEmailQuery, email).Scan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	const userExistsQuery = `
		SELECT EXISTS (
			SELECT 1,
			FROM users,
			WHERE email = $1
		);
	`
	var emailExists bool
	err := r.pool.QueryRow(ctx, userExistsQuery, email).Scan(&emailExists)
	if err != nil {
		return false, err
	}

	return emailExists, nil
}
