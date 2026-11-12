package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"saffronreal.estate/backend/internal/repository"
	"saffronreal.estate/backend/tests"
)

func TestCreateValidEmail(t *testing.T) {
	pool := tests.NewTestDB(t)
	s := NewUserService(repository.NewUserRepository(pool))
	email := "user@example.com"

	user, err := s.Create(tests.Ctx, email)
	assert.Equal(t, email, user.Email)
	assert.NoError(t, err)
}

func TestCreateInvalidEmail(t *testing.T) {
	pool := tests.NewTestDB(t)
	s := NewUserService(repository.NewUserRepository(pool))
	email := "invalid_email"

	_, err := s.Create(tests.Ctx, email)
	assert.ErrorIs(t, err, ErrInvalidEmail)

	_, err = s.Get(tests.Ctx, email)
	assert.ErrorIs(t, err, ErrInvalidEmail)
}

func TestToLowercase(t *testing.T) {
	pool := tests.NewTestDB(t)
	s := NewUserService(repository.NewUserRepository(pool))
	email := "ValidCapitalEmail@Example.com"

	_, err := s.Create(tests.Ctx, email)

	user, err := s.Get(tests.Ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, "validcapitalemail@example.com", user.Email)

	user, err = s.Get(tests.Ctx, strings.ToLower(email))
	assert.NoError(t, err)
	assert.Equal(t, "validcapitalemail@example.com", user.Email)
}
