package domain

import (
	"context"
)

type SKU int32

type StockService interface {
	Reserve(ctx context.Context, SKU SKU) error
	ReserveRemove(ctx context.Context, SKU SKU) error
	ReserveCancel(ctx context.Context, SKU SKU) error
}
