package domain

import (
	"context"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type OrderRepository interface {
	Create(ctx context.Context, userID uint64, items []models.OrderItem) (models.Order, error)
	SetStatusFailed(ctx context.Context, ID uint64) (models.Order, error)
	SetStatusAwaitingPayment(ctx context.Context, ID uint64) (models.Order, error)
	GetByID(ctx context.Context, ID uint64) (models.Order, error)
}
