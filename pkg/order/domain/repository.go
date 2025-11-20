package domain

import (
	"context"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, order entity.Order) (entity.Order, error)
	SetStatus(ctx context.Context, ID int, status string) (entity.Order, error)
	GetByID(ctx context.Context, ID int) (entity.Order, error)
}
