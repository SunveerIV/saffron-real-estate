package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"saffronreal.estate/backend/tests"
)

func TestNewUserNoCollision(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewUserRepository(pool)
	email := "user@example.com"

	user, err := r.NewUser(tests.Ctx, email)
	require.NoError(t, err)
	assert.True(t, user.Id.Valid)
	assert.Equal(t, email, user.Email)
	assert.False(t, user.TimeCreated.IsZero())
}

func TestGetUser(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewUserRepository(pool)
	email := "user@example.com"

	created, err := r.NewUser(tests.Ctx, email)
	require.NoError(t, err)

	user, err := r.GetUser(tests.Ctx, email)
	require.NoError(t, err)
	assert.Equal(t, created, user)
}

func TestNewUserWithCollision(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewUserRepository(pool)
	email := "user@example.com"

	_, err := r.NewUser(tests.Ctx, email)
	require.NoError(t, err)

	_, err = r.NewUser(tests.Ctx, email)
	assert.ErrorIs(t, err, ErrEmailConflict)
}

func TestUserNotFound(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewUserRepository(pool)
	email := "user@example.com"

	_, err := r.GetUser(tests.Ctx, email)
	assert.ErrorIs(t, err, ErrUserNotFound)
}
