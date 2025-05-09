package ports

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

type StockRepository interface {
	Reserve(ctx context.Context, SKU string, quantity int) error
	ReserveCancel(ctx context.Context, SKU string, quantity int) error
	ReserveRemove(ctx context.Context, SKU string, quantity int) error
	GetBySKU(ctx context.Context, SKU string) ([]domain.StockItem, error)
}
