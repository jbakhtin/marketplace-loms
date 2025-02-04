package primary

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/stock/models"
)

type UseCase interface {
	StockInfo(ctx context.Context, SKU models.SKU) (models.Stock, error)
	CancelReservation(ctx context.Context, SKU models.SKU) error
}
