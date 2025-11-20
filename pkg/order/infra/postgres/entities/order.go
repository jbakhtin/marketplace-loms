package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Order struct {
	UserId    int64            `json:"user_id,omitempty" db:"user_id"`
	Status    string           `json:"status,omitempty" db:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}
