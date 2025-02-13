package primary

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/stock/models"
)

type UseCase interface {
	StockInfo(ctx context.Context, SKU models.StockItemSKU) (models.StockItem, error)
	CancelReservation(ctx context.Context, SKU models.StockItemSKU) error
}
