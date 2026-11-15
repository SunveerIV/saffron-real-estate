package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saffronreal.estate/backend/internal/repository"
	"saffronreal.estate/backend/tests"
)

func TestParams(t *testing.T) {
	pool := tests.NewTestDB(t)
	us := NewUserService(repository.NewUserRepository(pool))
	hbrs := NewHouseBuyRequestService(repository.NewHouseBuyRequestRepository(pool), us)

	const email = "user@example.com"

	_, err := us.Create(tests.Ctx, email)
	assert.NoError(t, err)

	params := CreateHouseBuyRequestParams{
		Email:             email,
		FirstName:         "A",
		LastName:          "B",
		Address:           "123 way",
		PriceOfferedCents: int64(10000000),
		Comments:          "comment",
	}

	params.Email = ""
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidEmail)
	params.Email = email
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	params.FirstName = ""
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidFirstName)
	params.FirstName = "A"
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	params.LastName = ""
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidLastName)
	params.LastName = "B"
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	params.Address = ""
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidAddress)
	params.Address = "123 way"
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	params.Comments = "........................................................................................................................................................................................................."
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidComments)
	params.Comments = ""
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	params.PriceOfferedCents = int64(-1)
	_, err = hbrs.Create(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrInvalidPriceOfferedCents)
	params.PriceOfferedCents = int64(1000)
	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)
}

func TestNewUserCreation(t *testing.T) {
	pool := tests.NewTestDB(t)
	us := NewUserService(repository.NewUserRepository(pool))
	hbrs := NewHouseBuyRequestService(repository.NewHouseBuyRequestRepository(pool), us)

	const email = "user@example.com"

	params := CreateHouseBuyRequestParams{
		Email:             email,
		FirstName:         "A",
		LastName:          "B",
		Address:           "123 way",
		PriceOfferedCents: int64(10000000),
		Comments:          "comment",
	}

	_, err := hbrs.GetHouseBuyRequests(tests.Ctx, email)
	assert.ErrorIs(t, err, repository.ErrUserNotFound)

	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)

	user, err := us.Get(tests.Ctx, email)
	assert.NoError(t, err)
	assert.True(t, user.Id.Valid)
	assert.Equal(t, email, user.Email)

	requests, err := hbrs.GetHouseBuyRequests(tests.Ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(requests))
}

func TestMultipleBuyRequestsOnSameUser(t *testing.T) {
	pool := tests.NewTestDB(t)
	us := NewUserService(repository.NewUserRepository(pool))
	hbrs := NewHouseBuyRequestService(repository.NewHouseBuyRequestRepository(pool), us)

	const email = "user@example.com"

	user, err := us.Create(tests.Ctx, email)
	assert.NoError(t, err)

	requests, err := hbrs.GetHouseBuyRequests(tests.Ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(requests))

	params := CreateHouseBuyRequestParams{
		Email:             user.Email,
		FirstName:         "A",
		LastName:          "B",
		Address:           "123 way",
		PriceOfferedCents: int64(10000000),
		Comments:          "comment",
	}

	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)
	requests, err = hbrs.GetHouseBuyRequests(tests.Ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(requests))

	_, err = hbrs.Create(tests.Ctx, params)
	assert.NoError(t, err)
	requests, err = hbrs.GetHouseBuyRequests(tests.Ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(requests))
}
