package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	SetStatus(ctx context.Context, ID int, status string) (domain.Order, error)
	GetByID(ctx context.Context, ID int) (domain.Order, error)
}
