package secondary

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/stock/models"
)

type StockService interface {
	ReserveCancel(ctx context.Context, SKU models.SKU) error
}
