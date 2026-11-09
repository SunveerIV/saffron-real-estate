package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/model"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrEmailConflict = errors.New("email is already in use")
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
		return model.User{}, ErrEmailConflict
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
		return model.User{}, ErrUserNotFound
	}

	return r.userFromEmail(ctx, email)
}

func (r *UserRepository) userFromEmail(ctx context.Context, email string) (model.User, error) {
	const userFromEmailQuery = `
		SELECT * FROM users WHERE email = $1 LIMIT 1
	`
	var user model.User

	err := r.pool.QueryRow(ctx, userFromEmailQuery, email).Scan(
		&user.Id,
		&user.Email,
		&user.TimeCreated,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	const userExistsQuery = `
		SELECT EXISTS (
			SELECT * FROM users WHERE email = $1
		);
	`
	var emailExists bool
	err := r.pool.QueryRow(ctx, userExistsQuery, email).Scan(&emailExists)
	if err != nil {
		return false, err
	}

	return emailExists, nil
}
