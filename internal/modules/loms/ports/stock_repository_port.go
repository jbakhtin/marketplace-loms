package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/stock/models"
)

type StockStorage interface {
	Reserve(ctx context.Context, SKU string, quantity int) error
	ReserveCancel(ctx context.Context, SKU string, quantity int) error
	ReserveRemove(ctx context.Context, SKU string, quantity int) error
	GetBySKU(ctx context.Context, SKU string) ([]models.StockItem, error)
}
