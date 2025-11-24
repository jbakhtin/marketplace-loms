package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type OrderRepository interface {
	Create(ctx context.Context, items []models.OrderItem) (models.Order, error)
	SetStatus(ctx context.Context, ID int, status string) (models.Order, error)
	GetByID(ctx context.Context, ID int) (models.Order, error)
}
