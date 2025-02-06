package secondary

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order models.Order) (models.Order, error)
	SetStatus(ctx context.Context, ID int, status string) (models.Order, error)
	GetByID(ctx context.Context, ID int) (models.Order, error)
}
