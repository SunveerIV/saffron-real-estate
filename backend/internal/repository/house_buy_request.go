package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/model"
)

type HouseBuyRequestRepository struct {
	pool *pgxpool.Pool
}

type CreateHouseBuyRequestParams struct {
	UserId            pgtype.UUID
	FirstName         string
	LastName          string
	Address           string
	PriceOfferedCents int64
	Comments          string
}

func NewHouseBuyRequestRepository(pool *pgxpool.Pool) *HouseBuyRequestRepository {
	return &HouseBuyRequestRepository{
		pool: pool,
	}
}

func (r *HouseBuyRequestRepository) NewHouseBuyRequest(
	ctx context.Context,
	params CreateHouseBuyRequestParams,
) (model.HouseBuyRequest, error) {

	const query = `
		INSERT INTO house_buy_requests(
		    first_name,
		    last_name,
		    address,
			price_offered_cents,
		    user_id,
			comments
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		params.FirstName,
		params.LastName,
		params.Address,
		params.PriceOfferedCents,
		params.UserId,
		params.Comments,
	)
	if err != nil {
		return model.HouseBuyRequest{}, err
	}

	request, err := pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[model.HouseBuyRequest],
	)

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok &&
		pgErr.Code == "23503" {
		return model.HouseBuyRequest{}, ErrUserNotFound
	}

	return request, err
}

func (r *HouseBuyRequestRepository) GetHouseBuyRequests(
	ctx context.Context,
	userId pgtype.UUID,
) ([]model.HouseBuyRequest, error) {

	const query = `
		SELECT * FROM house_buy_requests
		WHERE user_id = $1
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		userId,
	)
	if err != nil {
		return []model.HouseBuyRequest{}, err
	}

	requests, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[model.HouseBuyRequest],
	)

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok &&
		pgErr.Code == "23503" {
		return []model.HouseBuyRequest{}, ErrUserNotFound
	}

	return requests, err
}
