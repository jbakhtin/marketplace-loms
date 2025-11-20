package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderItem struct {
	ID        int              `json:"id,omitempty" db:"id"`
	SKU       int32            `json:"sku,omitempty" db:"sku"`
	Count     uint16           `json:"count,omitempty" db:"count"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}
