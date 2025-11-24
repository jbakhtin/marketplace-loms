package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type OrderItem struct {
	ID        int              `json:"id,omitempty" db:"id"`
	OrderID   int64            `json:"order_id,omitempty" db:"order_id"`
	SKU       int32            `json:"sku,omitempty" db:"sku"`
	Quantity  int              `json:"quantity,omitempty" db:"quantity"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

func (oi *OrderItem) ToModel() models.OrderItem {
	return models.OrderItem{
		SKU:   oi.SKU,
		Count: uint16(oi.Quantity),
	}
}
