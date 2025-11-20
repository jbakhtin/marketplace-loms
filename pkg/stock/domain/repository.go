package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
)

type StockRepository interface {
	Reserve(ctx context.Context, SKU int32, qty uint16) error
	ReserveCancel(ctx context.Context, SKU int32, qty uint16) error
	ReserveRemove(ctx context.Context, SKU int32, qty uint16) error
	GetBySKU(ctx context.Context, SKU int32) (models.StockItem, error)
}
