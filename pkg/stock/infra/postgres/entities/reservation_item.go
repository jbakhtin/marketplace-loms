package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
)

type ReservationItem struct {
	ID        uint             `json:"id,omitempty" db:"id"`
	OrderID   int64            `json:"order_id,omitempty" db:"order_id"`
	SKU       int              `json:"sku,omitempty" db:"sku"`
	Reserved  int              `json:"reserved,omitempty" db:"reserved"`
	Status    string           `json:"status,omitempty" db:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

func (s *ReservationItem) ToModel() models.ReservationItem {
	return models.ReservationItem{
		OrderID:   uint(s.OrderID),
		SKU:       int32(s.SKU),
		Reserved:  uint64(s.Reserved),
		CreatedAt: s.CreatedAt.Time,
	}
}
