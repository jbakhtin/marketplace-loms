package primary

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/models"
)

type UseCase interface {
	CreateOrder(ctx context.Context, orderItems []models.OrderItem) (models.Order, error)
	OrderInfo(ctx context.Context, orderID models.OrderID) (models.Order, error)
	OrderPay(ctx context.Context, ID models.OrderID) error
	CancelOrder(ctx context.Context, ID models.OrderID) error
}
