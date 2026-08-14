package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type HouseBuyRequest struct {
	Id                pgtype.UUID `json:"id" db:"id"`
	UserId            pgtype.UUID `json:"user_id" db:"user_id"`
	FirstName         string      `json:"first_name" db:"first_name"`
	LastName          string      `json:"last_name" db:"last_name"`
	Address           string      `json:"address" db:"address"`
	TimeReceived      time.Time   `json:"time_received" db:"time_received"`
	PriceOfferedCents int64       `json:"price_offered_cents" db:"price_offered_cents"`
}
