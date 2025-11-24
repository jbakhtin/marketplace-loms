package domain

import (
	"context"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type OrderRepository interface {
	Create(ctx context.Context, userID uint64, items []models.OrderItem) (models.Order, error)
	SetStatus(ctx context.Context, ID int64, status string) (models.Order, error)
	GetByID(ctx context.Context, ID int64) (models.Order, error)
}
