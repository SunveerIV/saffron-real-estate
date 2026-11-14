package repository

import (
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"saffronreal.estate/backend/internal/model"
	"saffronreal.estate/backend/tests"
)

func TestNewHouseBuyRequest(t *testing.T) {
	pool := tests.NewTestDB(t)
	ur := NewUserRepository(pool)
	hr := NewHouseBuyRequestRepository(pool)

	user, _ := ur.NewUser(tests.Ctx, "email@example.com")

	params := CreateHouseBuyRequestParams{
		UserId:            user.Id,
		FirstName:         "A",
		LastName:          "B",
		Address:           "123 Road",
		PriceOfferedCents: 10000000,
		Comments:          "comment",
	}

	request, err := hr.NewHouseBuyRequest(tests.Ctx, params)
	assert.NoError(t, err)
	assert.True(t, request.Id.Valid)
	assert.Equal(t, "A", request.FirstName)
	assert.Equal(t, "B", request.LastName)
	assert.Equal(t, "123 Road", request.Address)
	assert.Equal(t, int64(10000000), request.PriceOfferedCents)
	assert.Equal(t, "comment", request.Comments)
	assert.False(t, request.TimeReceived.IsZero())
}

func TestGetHouseBuyRequests(t *testing.T) {
	for _, count := range []int{0, 1, 2} {
		pool := tests.NewTestDB(t)
		users := NewUserRepository(pool)
		requests := NewHouseBuyRequestRepository(pool)

		user, err := users.NewUser(tests.Ctx, "buyer@example.com")
		require.NoError(t, err)
		otherUser, err := users.NewUser(tests.Ctx, "other@example.com")
		require.NoError(t, err)

		_, err = requests.NewHouseBuyRequest(tests.Ctx, CreateHouseBuyRequestParams{
			UserId: otherUser.Id, FirstName: "Other", LastName: "Buyer",
			Address: "999 Other Road", PriceOfferedCents: 20000000, Comments: "Other request",
		})
		require.NoError(t, err)

		want := make([]model.HouseBuyRequest, 0, count)
		for i := 0; i < count; i++ {
			request, err := requests.NewHouseBuyRequest(tests.Ctx, CreateHouseBuyRequestParams{
				UserId: user.Id, FirstName: "First", LastName: "Last",
				Address:           fmt.Sprintf("%d Main Road", i+1),
				PriceOfferedCents: int64(10000000 + i), Comments: fmt.Sprintf("Request %d", i+1),
			})
			require.NoError(t, err)
			want = append(want, request)
		}

		got, err := requests.GetHouseBuyRequests(tests.Ctx, user.Id)
		require.NoError(t, err)
		assert.ElementsMatch(t, want, got)
	}
}

func TestGetHouseBuyRequestsUnknownUser(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewHouseBuyRequestRepository(pool)

	requests, err := r.GetHouseBuyRequests(tests.Ctx, pgtype.UUID{
		Bytes: [16]byte{0x12, 0x34, 0x56, 0x78},
		Valid: true,
	})
	require.NoError(t, err)
	assert.Empty(t, requests)
}

func TestNoUserReference(t *testing.T) {
	pool := tests.NewTestDB(t)
	r := NewHouseBuyRequestRepository(pool)

	params := CreateHouseBuyRequestParams{
		UserId: pgtype.UUID{
			Bytes: [16]byte{0x12, 0x34, 0x56, 0x78},
			Valid: true,
		},
		FirstName:         "A",
		LastName:          "B",
		Address:           "123 Road",
		PriceOfferedCents: 10000000,
		Comments:          "comment",
	}

	_, err := r.NewHouseBuyRequest(tests.Ctx, params)
	assert.ErrorIs(t, err, ErrUserNotFound)
}
