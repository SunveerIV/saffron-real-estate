package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id          pgtype.UUID `json:"id" db:"id"`
	Email       string      `json:"email" db:"email"`
	TimeCreated time.Time   `json:"time_created" db:"time_created"`
}
