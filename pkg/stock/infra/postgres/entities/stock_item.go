package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
)

type StockItem struct {
	ID        int              `json:"id,omitempty" db:"id"`
	SKU       int              `json:"sku,omitempty" db:"sku"`
	Available int              `json:"available,omitempty" db:"available"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

func (s *StockItem) ToModel() models.StockItem {
	return models.StockItem{
		SKU:       int32(s.SKU),
		Available: uint64(s.Available),
	}
}
